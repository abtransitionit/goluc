/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/abtransitionit/gotask/dnfapt"
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
	vmList              = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	listRequiredPackage = []string{"uidmap"}
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
		// corephase.NewPhase("showConfig", "display the desired KIND Cluster's configuration", vm.CheckVmSshAccess, nil),
		// corephase.NewPhase("show2", "display the desired KIND Cluster's configuration", internal.CheckSystemStatus, nil),
		corephase.NewPhase("updateApp", "provision required/missing standard dnfapt packages.", dnfapt.UpdateVmOsApp(listRequiredPackage), []string{"upgradeOs"}),
		corephase.NewPhase("dapack1", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"upgradeOs"}),
		corephase.NewPhase("gocli", "provision Go toolchain", internal.GenerateReport, []string{"dapack1"}),
		corephase.NewPhase("service", "configure OS services on Kind VMs.", internal.GenerateReport, []string{"dapack1"}),
		corephase.NewPhase("linger", "Allow non root user to run OS services.", internal.GenerateReport, []string{"dapack1"}),
		corephase.NewPhase("path", "configure OS PATH envvar.", internal.GenerateReport, []string{"dapack1"}),
		corephase.NewPhase("rc", "Add a line to non-root user RC file.", internal.GenerateReport, []string{"dapack1"}),
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
