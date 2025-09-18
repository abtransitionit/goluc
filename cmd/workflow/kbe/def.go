/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	coregocli "github.com/abtransitionit/gocore/gocli"
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	linuxdnfapt "github.com/abtransitionit/golinux/dnfapt"
	linuxk8s "github.com/abtransitionit/golinux/k8s"
	liuxoservice "github.com/abtransitionit/golinux/oservice"
	linuxkernel "github.com/abtransitionit/golinux/oskernel"
	"github.com/abtransitionit/gotask/dnfapt"
	"github.com/abtransitionit/gotask/gocli"
	taskk8s "github.com/abtransitionit/gotask/k8s"
	"github.com/abtransitionit/gotask/luc"
	"github.com/abtransitionit/gotask/oservice"
	taskoskernel "github.com/abtransitionit/gotask/oskernel"
	"github.com/abtransitionit/gotask/selinux"
	"github.com/abtransitionit/gotask/util"
	"github.com/abtransitionit/gotask/vm"
)

// Package variables
var (
	logger           = logx.GetLogger()
	wkf              *corephase.Workflow
	targets          []corephase.Target
	targetsCP        []corephase.Target
	targetsWorker    []corephase.Target
	customRcFileName = ".profile.luc "  // name of the user's custom rc file
	binFolderPath    = "/usr/local/bin" // location of binaries
)

// Package variables : confifg1s
var (
	cmdName    = "kbe"
	SDesc      = "This is the KuBernetes Easy workflow."
	K8sVersion = "1.32.0"
)

// Package variables : confifg2
var (
	// vmListNode = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	vmListNode = []string{"o1u"}
	// vmListNode            = []string{"o1u", "o2a"}
	// vmListControlPlaneNode = []string{"o1u", "o3r"}
	vmListControlPlaneNode = []string{"o1u"}
	vmListWorkerNode       = []string{"o2a"}
	// vmListWorkerNode       = []string{"o2a", "o5d"}
	listRequiredDaPackage = []string{"gnupg"} // gnupg/{gpg}
	listGoCli             = coregocli.SliceGoCli{
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
	sliceDaPackCplane = linuxdnfapt.SliceDaPack{
		{Name: "kubectl"},
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
	k8sConf = linuxk8s.K8sConf{
		K8sVersion:     K8sVersion,
		K8sPodCidr:     "192.168.0.0/16",
		K8sServiceCidr: "172.16.0.0/16",
		CrSocketName:   "unix:///var/run/crio/crio.sock",
	}
)

func init() {
	// create the targets slice from vmListNode
	for _, vmName := range vmListNode {
		targets = append(targets, &corephase.Vm{NameStr: vmName})
	}

	// create the targets slice from vmListControlPlaneNode
	for _, vmName := range vmListControlPlaneNode {
		targetsCP = append(targetsCP, &corephase.Vm{NameStr: vmName})
	}

	// create the targets slice from vmListWorkerNode
	for _, vmName := range vmListWorkerNode {
		targetsWorker = append(targetsWorker, &corephase.Vm{NameStr: vmName})
	}

	// create the workflow
	var err error
	wkf, err = corephase.NewWorkflowFromPhases(
		corephase.NewPhase("checkVmAccess", "Check if VMs are SSH reachable", vm.CheckVmSshAccess, nil),
		corephase.NewPhase("copyAgent", "copy LUC CLI agent to all VMs", luc.DeployLuc, []string{"checkVmAccess"}),
		corephase.NewPhase("upgradeOs", "provision OS nodes with latest dnfapt packages and repositories.", dnfapt.UpgradeVmOs, []string{"copyAgent"}),
		corephase.NewPhase("updateApp", "provision required/missing standard dnfapt packages.", dnfapt.UpdateVmOsApp(listRequiredDaPackage), []string{"upgradeOs"}),
		corephase.NewPhase("installDaRepository", "provision Dnfapt package repositor(y)(ies).", dnfapt.InstallDaRepository(sliceDaRepo), []string{"updateApp"}),
		corephase.NewPhase("installDaPackage", "provision Dnfapt package(s) on all nodes.", dnfapt.InstallDaPackage(sliceDaPackNode), []string{"installDaRepository"}),
		corephase.NewPhase("installDaPackageCplane", "provision Dnfapt package(s) on CPlane only.", dnfapt.InstallDaPackage(sliceDaPackCplane, targetsCP), []string{"installDaPackage"}),
		corephase.NewPhase("loadOsKernelModule", "load OS kernel module(s).", taskoskernel.LoadOsKModule(sliceOsKModule, kFilename), []string{"installDaPackageCplane"}),
		corephase.NewPhase("loadOsKernelParam", "set OS kernel paramleter(s).", taskoskernel.LoadOsKParam(sliceOsKParam, kFilename), []string{"loadOsKernelModule"}),
		corephase.NewPhase("confSelinux", "Configure Selinux.", selinux.ConfigureSelinux(), []string{"loadOsKernelParam"}),
		corephase.NewPhase("enableOsService", "enable OS services to start after a reboot", oservice.EnableOsService(sliceOsServiceEnable), []string{"confSelinux"}),
		corephase.NewPhase("startOsService", "start OS services for current session", oservice.StartOsService(sliceOsServiceStart), []string{"confSelinux"}),
		corephase.NewPhase("resetCPlane", "reset the control plane(s).", taskk8s.ResetNode(targetsCP), []string{"startOsService"}),
		corephase.NewPhase("initCPlane", "initialize the control plane(s) (aka. boostrap the cluster).", taskk8s.InitCPlane(targetsCP, k8sConf), []string{"resetCPlane"}),
		corephase.NewPhase("resetWorker", "reset the workers(s).", taskk8s.ResetNode(targetsWorker), []string{"initCPlane"}),
		corephase.NewPhase("addWorker", "Add the K8s worker(s) to the K8s cluster.", taskk8s.AddWorker(targetsCP[0], targetsWorker), []string{"resetWorker"}),
		corephase.NewPhase("confKubectlOnCPlane", "Configure kubectl on the control plane(s).", taskk8s.ConfigureKubectlOnCplane(targetsCP[0]), []string{"resetWorker"}),
		corephase.NewPhase("installGoCli", "provision Go CLI(s).", gocli.InstallOnVm(listGoCli), []string{"confKubectlOnCPlane"}),
		corephase.NewPhase("createRcFile", "create a custom RC file in user's home.", util.CreateCustomRcFile(customRcFileName), []string{"installGoCli"}),
		corephase.NewPhase("setPathEnvar", "configure PATH envvar into current user's custom RC file.", util.SetPath(binFolderPath, customRcFileName), []string{"createRcFile"}), // corephase.NewPhase("installGoCli", "provision Go CLI(s).", taskgocli.InstallOnVm(listGoCli), []string{"updateApp"}),
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
