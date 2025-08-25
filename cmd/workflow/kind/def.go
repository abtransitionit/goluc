/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
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
	vmList = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
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
		corephase.NewPhase("upgrade", "provision OS nodes with latest dnfapt packages and repositories.", internal.GenerateReport, []string{"cpluc"}),
		corephase.NewPhase("showConfig", "display the desired KIND Cluster's configuration", vm.CheckVmSshAccess, nil),
		corephase.NewPhase("show2", "display the desired KIND Cluster's configuration", internal.CheckSystemStatus, nil),
		corephase.NewPhase("checklist", "check VMs are SSH reachable.", internal.FetchLatestData, nil),
		corephase.NewPhase("cpluc", "provision LUC CLI", internal.ProcessData, nil),
		corephase.NewPhase("dapack1", "provision standard/required/missing OS CLI (via dnfapt  packages).", internal.CheckSystemStatus, []string{"upgrade"}),
		corephase.NewPhase("dapack2", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"upgrade"}),
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
