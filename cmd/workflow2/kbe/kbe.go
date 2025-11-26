/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/file"
	"github.com/abtransitionit/gotask/mock/node"
	"github.com/abtransitionit/gotask/mock/onpm"
)

// Package variables
var (
	cmdPathName = "kbe"
	shortDesc   = "create KBE (Kubernetes Easy) clusters."
)

// registred function
func registerFunctions() {

	// get an instance of the FnRegistry that is shared by all phases of a
	registry := phase2.GetFnRegistry()
	// get the workflow name
	worflowName := filepath.Base(cmdPathName)
	// register function used
	registry.Add(worflowName, "AddNativeRepo", onpm.AddRepo)
	registry.Add(worflowName, "AddNativePkg", onpm.AddPkg)
	registry.Add(worflowName, "CheckSshConf", node.CheckSshConf)
	registry.Add(worflowName, "CheckSshAccess", node.CheckSshAccess)
	registry.Add(worflowName, "DeployAgent", file.CopyFileWithSudo)
	registry.Add(worflowName, "RebootIfNeeded", node.RebootIfNeeded)
	registry.Add(worflowName, "RebootIfNeeded", node.RebootIfNeeded)
	registry.Add(worflowName, "UpgradeOs", onpm.UpgradeOs)
	registry.Add(worflowName, "WaitIsOnline", node.WaitIsSshOnline)
}
