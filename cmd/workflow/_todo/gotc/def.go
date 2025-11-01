/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gotc

import (
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
	cmdName = "gotc"
	SDesc   = "This is the GO toochain workflow."
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
		corephase.NewPhase("showPhase", "display the worflow execution plan", workflow.ShowWorkflow, nil),
		corephase.NewPhase("checkVmAccess", "Check if VMs are SSH reachable", vm.CheckVmSshAccess, nil),
		corephase.NewPhase("gocli", "provision Go toolchain", internal.GenerateReport, nil),
		// corephase.NewPhase("showConfig", "display the desired KIND Cluster's configuration", vm.CheckVmSshAccess, nil),
		// corephase.NewPhase("show2", "display the desired KIND Cluster's configuration", internal.CheckSystemStatus, nil),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}
