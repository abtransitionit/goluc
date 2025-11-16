/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package luc

import (
	common "github.com/abtransitionit/goluc/cmd/workflow2/_common"
)

// root Command
var EpCmd = common.GetEpCmd(
	cmdPathName,
	shortDesc,
)

func init() {
	// sub cde
	common.SetInitSubCmd(EpCmd, cmdPathName)

	// register functions
	registerFunctions()

}
