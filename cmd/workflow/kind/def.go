/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/abtransitionit/gotask/workflow"
)

// Package variables

var (
	SDesc   = "This is the KIND workflow."
	cmdName = "kind"
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
)

func init() {
	var err error
	wkf, err = corephase.NewWorkflowFromPhases(
		corephase.NewPhase("showPhase", "display the worflow execution plan", workflow.ShowWorkflow, nil),
		corephase.NewPhase("show2", "display the desired KIND Cluster's configuration", internal.CheckSystemStatus, nil),
		corephase.NewPhase("checklist", "check VMs are SSH reachable.", internal.FetchLatestData, nil),
		corephase.NewPhase("cpluc", "provision LUC CLI", internal.ProcessData, nil),
		corephase.NewPhase("upgrade", "provision OS nodes with latest dnfapt packages and repositories.", internal.GenerateReport, []string{"cpluc"}),
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
