/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"context"

	coregocli "github.com/abtransitionit/gocore/gocli"
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/golinux/envar"
	liuxoservice "github.com/abtransitionit/golinux/oservice"
	"github.com/abtransitionit/gotask/dnfapt"
	taskgocli "github.com/abtransitionit/gotask/gocli"
	"github.com/abtransitionit/gotask/luc"
	"github.com/abtransitionit/gotask/oservice"
	"github.com/abtransitionit/gotask/util"
	"github.com/abtransitionit/gotask/vm"
	"github.com/abtransitionit/gotask/workflow"
)

// Package variables : confifg1
var (
	logger           = logx.GetLogger()
	wkf              *corephase.Workflow
	targets          []corephase.Target
	customRcFileName = ".profile.luc "  // name of the user's custom rc file
	binFolderPath    = "/usr/local/bin" // location of binaries
)

// Package variables : confifg2
var (
	cmdName = "kind" // the app name - should also be the workflow name
	SDesc   = "This is the KIND workflow."
)

// Package variables : confifg3
var (
	// vmList = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	vmList                = []string{"o1u"}
	listRequiredDaPackage = []string{"uidmap"} // uidmap/{newuidmap, newgidmap}
	listGoCli             = []coregocli.GoCli{
		{Name: "cni", Version: "1.7.1"},
		{Name: "containerd", Version: "2.1.1"},
		{Name: "kind", Version: "latest"},
		{Name: "nerdctl", Version: "2.1.2"},
		{Name: "rootlesskit", Version: "2.3.5"},
		{Name: "runc", Version: "1.3.0"},
		{Name: "slirp4netns", Version: "1.3.3"},
	}
	apparmorContent = `
		# Allow rootlesskit to create user namespaces (userns)
		# Ref: https://ubuntu.com/blog/ubuntu-23-10-restricted-unprivileged-user-namespaces
		abi <abi/4.0>,
		include <tunables/global>

		/usr/local/bin/rootlesskit/rootlesskit flags=(unconfined) {
			userns,

			# Site-specific additions and overrides. See local/README for details.
			include if exists <local/usr.local.bin.rootlesskit.rootlesskit>
		}
	`

	listOsService = []liuxoservice.OsService{
		{Name: "apparmor", Path: "/etc/apparmor.d/usr.local.bin.rootlesskit.rootlesskit", Content: apparmorContent}, // active and enabled by default
	}
	listEnvVar = []envar.EnvVar{
		{Name: "CNI_PATH", Value: "/usr/local/bin/cni"},
	}
)

func init() {
	// create the targets slice from vmList
	for _, vmName := range vmList {
		targets = append(targets, &corephase.Vm{NameStr: vmName})
	}

	// create the workflow
	var err error
	wkf, err = corephase.NewWorkflowFromPhases(
		corephase.NewPhase("checkVmAccess", "Check if VMs are SSH reachable", vm.CheckVmSshAccess, nil),
		corephase.NewPhase("copyAgent", "copy LUC CLI agent to all VMs", luc.DeployLuc, []string{"checkVmAccess"}),
		corephase.NewPhase("upgradeOs", "provision OS nodes with latest dnfapt packages and repositories.", dnfapt.UpgradeVmOs, []string{"copyAgent"}),
		corephase.NewPhase("updateApp", "provision required/missing standard dnfapt packages.", dnfapt.UpdateVmOsApp(listRequiredDaPackage), []string{"upgradeOs"}),
		corephase.NewPhase("installGoCli", "provision Go CLI(s).", taskgocli.InstallOnVm(listGoCli), []string{"updateApp"}),
		corephase.NewPhase("installOsService", "provision Os service(s).", oservice.InstallOsService(listOsService), []string{"installGoCli"}),
		corephase.NewPhase("enablelinger", "Allows user services to be session independant", oservice.EnableLinger, []string{"installGoCli"}),
		corephase.NewPhase("createRcFile", "create a custom RC file in user's home.", util.CreateCustomRcFile(customRcFileName), []string{"enablelinger"}),
		corephase.NewPhase("setPathEnvar", "configure PATH envvar into current user's custom RC file.", util.SetPath(binFolderPath, customRcFileName), []string{"createRcFile"}),
		corephase.NewPhase("setEnvar", "define envvars into current user's custom RC file.", util.SetEnvar(customRcFileName, listEnvVar), []string{"setPathEnvar"}),
		// corephase.NewPhase("setContainerd", "sets up a rootless session independant containerd env for the current user.", util.SetContainerd(), []string{"setPathEnvar"}),
		// corephase.NewPhase("startOsService", "start OS services needed by thge app", oservice.StartOsService(listOsService), []string{"setPathEnvar"}),
		// corephase.NewPhase("service", "configure OS services on Kind VMs.", internal.GenerateReport, []string{"installGoCli"}),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}

// The adapter function to bridge the type mismatch.
// It accepts the required PhaseFunc signature and calls the specific ShowPhase.
func showPhaseAdapter(ctx context.Context, l logx.Logger, cmd ...string) (string, error) {
	// The adapter can access the package-level 'wkf' variable since it's in the same package.
	// it can also acces external method and feed them with anything that nelongs to KIND
	err := workflow.ShowPhase(wkf, l)
	if err != nil {
		return "", err
	}
	// Return the required string and a nil error on success.
	return "Workflow plan displayed.", nil
}

// install on debian: packageName = "gnupg" / cliName = "gpg" (to check existence)
// install packageName = "uidmap" on all except debian
// if pkgName == "uidmap" && data.OsFamily != "debian" {
// 	logx.L.Debugf("[%s] [%s] Skipping package : OS family is %s, not 'debian'.", vm, pkgName, data.OsFamily)
// 	continue
