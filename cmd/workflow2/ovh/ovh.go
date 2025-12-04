package ovh

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/node"
	"github.com/abtransitionit/gotask/mock/ovh"
)

// Package variables : confifg2
var (
	cmdPathName = "ovh"
	shortDesc   = "send a request to OVH infra to get infos or make changes."
)

// registred function
func registerFunctions() {
	var registry *phase2.FnRegistry
	// get an instance of the FnRegistry that is shared by all phases of a
	// get the workflow name
	worflowName := filepath.Base(cmdPathName)
	// register function used
	registry = phase2.GetFnRegistry()
	registry.Add(worflowName, "ListInfo", ovh.ListInfo)
	registry.Add(worflowName, "InstallVpsImage", ovh.InstallVpsImage)
	registry.Add(worflowName, "WaitIsOnline", node.WaitIsSshOnline)
	registry.Add(worflowName, "RenewOvhToken", ovh.RenewOvhToken)
	registry.Add(worflowName, "ListVpsOsImage", ovh.ListImageAvailable)

}
