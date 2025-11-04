/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package dep

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
	EpCmd.AddCommand(common.GetPrintCmd(cmdPathName))
	// EpCmd.AddCommand(common.GetRunCmd(cmdName))

}
