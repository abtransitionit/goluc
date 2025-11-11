/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ssh

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	common "github.com/abtransitionit/goluc/cmd/workflow2/_common"
	"github.com/abtransitionit/gotask/mock/node"
)

// root Command
var EpCmd = common.GetEpCmd(
	cmdPathName,
	shortDesc,
)

func init() {

	// sub cde
	common.SetInitSubCmd(EpCmd, cmdPathName)

	// get an instance of a FnRegistry
	registry := phase2.GetFnRegistry()
	// register function used
	worflowName := filepath.Base(cmdPathName)
	registry.Add(worflowName, "CheckSshConf", node.CheckSshConf)
	registry.Add(worflowName, "CheckSshAccess", node.CheckSshAccess)

}
