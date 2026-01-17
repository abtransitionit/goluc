/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/dns"
	"github.com/abtransitionit/gotask/mock/file"
	"github.com/abtransitionit/gotask/mock/gopm"
	"github.com/abtransitionit/gotask/mock/k8s"
	"github.com/abtransitionit/gotask/mock/node"
	"github.com/abtransitionit/gotask/mock/onpm"
	"github.com/abtransitionit/gotask/mock/oskernel"
	"github.com/abtransitionit/gotask/mock/osservice/sys"
	"github.com/abtransitionit/gotask/mock/selinux"
)

// register function used in the workflow
var (
	cmdPathName = "kbe"
	shortDesc   = "create KBE (Kubernetes Easy) clusters."
)

// registred function
func registerFunctions() {

	// get objects
	registry := phase2.GetFnRegistry()        // an instance of the FnRegistry that is shared by all phases
	worflowName := filepath.Base(cmdPathName) // the workflow name

	// ensure node:ssh status
	registry.Add(worflowName, "CheckSshConf", node.CheckSshConf)
	registry.Add(worflowName, "CheckSshAccess", node.CheckSshAccess)

	// ensure node:OS is up to date
	registry.Add(worflowName, "UpgradeOs", onpm.UpgradeOs)
	registry.Add(worflowName, "UpdateOs", onpm.UpdateOs)
	registry.Add(worflowName, "RebootIfNeeded", node.RebootIfNeeded)
	registry.Add(worflowName, "WaitIsOnline", node.WaitIsSshOnline)

	// deploy an agent
	registry.Add(worflowName, "DeployAgent", file.CopyFileWithSudo)

	// mange Linux RC file
	registry.Add(worflowName, "CreateRcFile", file.CreateRcFile)
	registry.Add(worflowName, "RcAddPath", file.RcAddPath)

	// manage Linux OS Native Package
	registry.Add(worflowName, "AddRepoNative", onpm.AddRepo)
	registry.Add(worflowName, "AddPkgNative", onpm.AddPkg)

	// manage file resolv.conf
	registry.Add(worflowName, "FixDns", dns.FixDns)

	// manage Go Package
	registry.Add(worflowName, "AddPkgGo", gopm.AddPkgGo)

	// manage Linux OS kernel
	registry.Add(worflowName, "AddKModule", oskernel.LoadModule)
	registry.Add(worflowName, "AddKParam", oskernel.LoadParam)

	// manage Linux service
	registry.Add(worflowName, "EnableService", sys.Enable)
	registry.Add(worflowName, "StartService", sys.Start)
	registry.Add(worflowName, "StopService", sys.Stop)

	// manage Linux Selinux
	registry.Add(worflowName, "SelinuxConfigure", selinux.Configure)

	// manage kubernetes cluster
	registry.Add(worflowName, "ResetNode", k8s.ResetNode)
	registry.Add(worflowName, "InitCplane", k8s.InitCplane)
	registry.Add(worflowName, "AddWorker", k8s.AddWorker)
	// - kubectl
	registry.Add(worflowName, "ConfigureKubectl", k8s.ConfigureKubectl)
	// - Helm
	registry.Add(worflowName, "AddRepoHelm", k8s.AddRepoHelm)
	registry.Add(worflowName, "InstallReleaseHelm", k8s.InstallReleaseHelm)
	// - ingress
	// registry.Add(worflowName, "InstallIngressCilium", k8s.InstallIngressCilium)

}
