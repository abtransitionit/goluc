/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	coregocli "github.com/abtransitionit/gocore/gocli"
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	linuxdnfapt "github.com/abtransitionit/golinux/dnfapt"
	"github.com/abtransitionit/gotask/dnfapt"
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
	// vmList = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	vmList                = []string{"o1u", "o2a", "o4f"}
	listRequiredDaPackage = []string{"gnupg"} // gnupg/{gpg}
	listGoCli             = coregocli.SliceGoCli{
		{Name: "kind", Version: "latest"},
		{Name: "kubeadm", Version: "1.32.0"},
		{Name: "kubectl", Version: "1.32.0"},
		{Name: "helm", Version: "3.17.3"},
	}
	sliceDaRepo = linuxdnfapt.SliceDaRepo{
		{Name: "crio", FileName: "kbe-crio", Version: "1.32"},
		{Name: "k8s", FileName: "kbe-k8s", Version: "1.32"},
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
		corephase.NewPhase("installDaRepository", "provision Dnfapt package repositor(y)(ies).", dnfapt.InstallDaRepository(sliceDaRepo), []string{"updateApp"}),
		// corephase.NewPhase("installGoCli", "provision Go CLI(s).", taskgocli.InstallOnVm(listGoCli), []string{"updateApp"}),
		// corephase.NewPhase("installOsService", "provision Os service(s).", oservice.InstallOsService(listOsService), []string{"installGoCli"}),
		// corephase.NewPhase("dapack2", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"installGoCli"}),
		// corephase.NewPhase("darepo", "provision dnfapt repositories.", internal.GenerateReport, []string{"installGoCli"}),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}

// var KbeGoCliConfigMap = config.CustomCLIConfigMap{
// 	"kubeadm": {
// 		Name:      "kubeadm",
// 		Version:   KbeVersion,
// 		DstFolder: "/usr/local/bin",
// 	},
// 	"kubectl": {
// 		Name:      "kubectl",
// 		Version:   KbeVersion,
// 		DstFolder: "/usr/local/bin",
// 	},
// 	"helm": {
// 		Name:      "helm",
// 		Version:   "3.17.3",
// 		DstFolder: "/usr/local/bin",
// 	},
// }
