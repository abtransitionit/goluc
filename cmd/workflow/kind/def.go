/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"context"

	coregocli "github.com/abtransitionit/gocore/gocli"
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/abtransitionit/gotask/dnfapt"
	taskgocli "github.com/abtransitionit/gotask/gocli"
	"github.com/abtransitionit/gotask/luc"
	"github.com/abtransitionit/gotask/vm"
	"github.com/abtransitionit/gotask/workflow"
)

// Package variables
var (
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
	targets []corephase.Target
)

// Package variables : confifg1
var (
	cmdName = "kind"
	SDesc   = "This is the KIND workflow."
)

// Package variables : confifg2
var (
	vmList                = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
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
)

// install on debian: packageName = "gnupg" / cliName = "gpg" (to check existence)
// install packageName = "uidmap" on all except debian
// if pkgName == "uidmap" && data.OsFamily != "debian" {
// 	logx.L.Debugf("[%s] [%s] Skipping package : OS family is %s, not 'debian'.", vm, pkgName, data.OsFamily)
// 	continue

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
		corephase.NewPhase("service", "configure OS services on Kind VMs.", internal.GenerateReport, []string{"installGoCli"}),
		corephase.NewPhase("linger", "Allow non root user to run OS services.", internal.GenerateReport, []string{"installGoCli"}),
		corephase.NewPhase("path", "configure OS PATH envvar.", internal.GenerateReport, []string{"installGoCli"}),
		corephase.NewPhase("rc", "Add a line to non-root user RC file.", internal.GenerateReport, []string{"installGoCli"}),
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
