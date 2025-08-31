/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	coregocli "github.com/abtransitionit/gocore/gocli"
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/abtransitionit/gotask/dnfapt"
	taskgocli "github.com/abtransitionit/gotask/gocli"
	"github.com/abtransitionit/gotask/luc"
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
	vmList                = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	listRequiredDaPackage = []string{"gnupg"} // gnupg/{gpg}
	listGoCli             = []coregocli.GoCli{
		{Name: "kind", Version: "latest"},
	}

	// listGoCli             = []string{"gocli"}
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
		corephase.NewPhase("dapack2", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"installGoCli"}),
		corephase.NewPhase("darepo", "provision dnfapt repositories.", internal.GenerateReport, []string{"installGoCli"}),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}
