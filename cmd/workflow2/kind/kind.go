/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/file"
	"github.com/abtransitionit/gotask/mock/node"
	"github.com/abtransitionit/gotask/mock/onpm"
)

// Package variables : confifg2
var (
	cmdPathName = "kind"
	shortDesc   = "create KIND (Kubernetes in Docker) clusters."
)

// registred function
func registerFunctions() {

	// get objects
	registry := phase2.GetFnRegistry()        // an instance of the FnRegistry that is shared by all phases
	worflowName := filepath.Base(cmdPathName) // the workflow name

	// code to ensure node status
	registry.Add(worflowName, "CheckSshConf", node.CheckSshConf)
	registry.Add(worflowName, "CheckSshAccess", node.CheckSshAccess)

	// code to ensure OS is up to date
	registry.Add(worflowName, "UpgradeOs", onpm.UpgradeOs)
	registry.Add(worflowName, "UpdateOs", onpm.UpdateOs)
	registry.Add(worflowName, "RebootIfNeeded", node.RebootIfNeeded)
	registry.Add(worflowName, "WaitIsOnline", node.WaitIsSshOnline)

	// deploy an agent
	registry.Add(worflowName, "DeployAgent", file.CopyFileWithSudo)
}
