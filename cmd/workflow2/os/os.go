package os

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/onpm"
)

// Package variables : confifg2
var (
	cmdPathName = "os"
	shortDesc   = "test os upgrade and update."
)

// registred function
func registerFunctions() {

	// get an instance of the FnRegistry that is shared by all phases of a
	registry := phase2.GetFnRegistry()
	// get the workflow name
	worflowName := filepath.Base(cmdPathName)
	// register function used
	registry.Add(worflowName, "UpgradeOs", onpm.UpgradeOs)

}
