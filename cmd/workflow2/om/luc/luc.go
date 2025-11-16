package luc

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/git"
)

var (
	cmdPathName = "om/luc"
	shortDesc   = "manage LUC build and deploy."
)

// registred function
func registerFunctions() {

	// get an instance of the FnRegistry that is shared by all phases of a
	registry := phase2.GetFnRegistry()

	// get the workflow name
	worflowName := filepath.Base(cmdPathName)

	// register function used
	registry.Add(worflowName, "MergeDevToMain", git.MergeDevToMain)

}
