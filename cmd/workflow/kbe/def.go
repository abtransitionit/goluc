/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/abtransitionit/gotask/vm"
)

// Package variables
var (
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
	targets []corephase.Target
)

// Package variables : confifg1s
var (
	cmdName = "kbe"
	SDesc   = "This is the KuBernetes Easy workflow."
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
		corephase.NewPhase("show2", "display the desired KBE Cluster's configuration.", internal.CheckSystemStatus, nil),
		corephase.NewPhase("cpluc", "provision LUC CLI", internal.ProcessData, nil),
		corephase.NewPhase("upgrade", "provision OS nodes with latest dnfapt packages and repositories.", internal.GenerateReport, []string{"cpluc"}),
		corephase.NewPhase("dapack1", "provision standard/required/missing OS CLI (via dnfapt  packages).", internal.CheckSystemStatus, []string{"upgrade"}),
		corephase.NewPhase("dapack2", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"dapack1"}),
		corephase.NewPhase("darepo", "provision dnfapt repositories.", internal.GenerateReport, []string{"dapack1"}),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}
