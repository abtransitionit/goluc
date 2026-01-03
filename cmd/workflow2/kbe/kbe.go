/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
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

	// code to ensure node:ssh status
	registry.Add(worflowName, "CheckSshConf", node.CheckSshConf)
	registry.Add(worflowName, "CheckSshAccess", node.CheckSshAccess)

	// code to ensure node:OS is up to date
	registry.Add(worflowName, "UpgradeOs", onpm.UpgradeOs)
	registry.Add(worflowName, "UpdateOs", onpm.UpdateOs)
	registry.Add(worflowName, "RebootIfNeeded", node.RebootIfNeeded)
	registry.Add(worflowName, "WaitIsOnline", node.WaitIsSshOnline)

	// deploy an agent
	registry.Add(worflowName, "DeployAgent", file.CopyFileWithSudo)

	// mange RC file
	registry.Add(worflowName, "CreateRcFile", file.CreateRcFile)
	registry.Add(worflowName, "RcAddPath", file.RcAddPath)
	// registry.Add(worflowName, "UpdateRcFile", file.AddString)

	// code to manage Linux OS Native Package
	registry.Add(worflowName, "AddRepoNative", onpm.AddRepo)
	registry.Add(worflowName, "AddPkgNative", onpm.AddPkg)

	// code Linux OS specific
	registry.Add(worflowName, "AddKModule", oskernel.LoadModule)
	registry.Add(worflowName, "AddKParam", oskernel.LoadParam)
	registry.Add(worflowName, "EnableService", sys.Enable)
	registry.Add(worflowName, "StartService", sys.Start)
	registry.Add(worflowName, "SelinuxConfigure", selinux.Configure)

	// code K8s specific
	registry.Add(worflowName, "ResetNode", k8s.ResetNode)
	registry.Add(worflowName, "InitCplane", k8s.InitCplane)
	registry.Add(worflowName, "AddWorker", k8s.AddWorker)
	registry.Add(worflowName, "ConfigureKubectl", k8s.ConfigureKubectl)
	registry.Add(worflowName, "AddPkgGo", gopm.AddPkgGo)
	registry.Add(worflowName, "AddRepoHelm", k8s.AddRepoHelm)
	registry.Add(worflowName, "AddPkgHelm", k8s.AddChartHelm)

}
