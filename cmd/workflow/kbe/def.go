/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	coregocli "github.com/abtransitionit/gocore/gocli"
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	linuxdnfapt "github.com/abtransitionit/golinux/dnfapt"
	liuxoservice "github.com/abtransitionit/golinux/oservice"
	linuxkernel "github.com/abtransitionit/golinux/oskernel"
	"github.com/abtransitionit/gotask/dnfapt"
	"github.com/abtransitionit/gotask/luc"
	"github.com/abtransitionit/gotask/oservice"
	taskoskernel "github.com/abtransitionit/gotask/oskernel"
	"github.com/abtransitionit/gotask/selinux"
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
	vmListNode = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	// vmListNode            = []string{"o1u", "o2a"}
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
	sliceDaPackNode = linuxdnfapt.SliceDaPack{
		{Name: "crio"},
		{Name: "kubeadm"},
		{Name: "kubelet"},
	}
	kFilename      = "99-kbe.conf"
	sliceOsKModule = []string{"overlay", "br_netfilter"}
	sliceOsKParam  = linuxkernel.SliceOsKParam{
		{Kvp: "net.ipv4.ip_forward=1", Description: "Enable IPv4 packet forwarding - core kernel parameter"},
		{Kvp: "net.bridge.bridge-nf-call-iptables=1", Description: "Pass bridged IPv4 traffic to iptables - br_netfilter module parameter"},
		{Kvp: "net.bridge.bridge-nf-call-ip6tables=1", Description: "Pass bridged IPv6 traffic to iptables - br_netfilter module parameter"},
	}
	sliceOsServiceStart = []liuxoservice.OsService{
		{Name: "crio"},
	}
	sliceOsServiceEnable = []liuxoservice.OsService{
		{Name: "crio"},
		{Name: "kubelet"},
	}
)

func init() {
	// create the targets slice from vmListNode
	for _, vmName := range vmListNode {
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
		corephase.NewPhase("installDaPackage", "provision Dnfapt package(s).", dnfapt.InstallDaPackage(sliceDaPackNode), []string{"installDaRepository"}),
		corephase.NewPhase("loadOsKernelModule", "load OS kernel module(s).", taskoskernel.LoadOsKModule(sliceOsKModule, kFilename), []string{"checkVmAccess"}),
		corephase.NewPhase("loadOsKernelParam", "set OS kernel paramleter(s).", taskoskernel.LoadOsKParam(sliceOsKParam, kFilename), []string{"loadOsKernelModule"}),
		corephase.NewPhase("confSelinux", "Configure Selinux.", selinux.ConfigureSelinux(), []string{"loadOsKernelParam"}),
		// att this point kubelet service status should be activating only
		corephase.NewPhase("enableOsService", "enable OS services to start after a reboot", oservice.EnableOsService(sliceOsServiceEnable), []string{"confSelinux"}),
		corephase.NewPhase("startOsService", "start OS services for current session", oservice.StartOsService(sliceOsServiceStart), []string{"confSelinux"}),
		corephase.NewPhase("initCPlane", "initialize thz Control plane", oservice.StartOsService(sliceOsServiceStart), []string{"confSelinux"}),
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
