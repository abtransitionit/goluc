package file

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/file"
)

var (
	cmdPathName = "file"
	shortDesc   = "sudo manage file"
)

func registerFunctions() {

	// get an instance of the FnRegistry that is shared by all phases of a
	registry := phase2.GetFnRegistry()
	// get the workflow name
	worflowName := filepath.Base(cmdPathName)
	// register function used
	registry.Add(worflowName, "DeployAgent", file.CopyFileWithSudo)

}
