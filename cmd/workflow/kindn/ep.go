/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kindn

import common "github.com/abtransitionit/goluc/cmd/workflow/_common"

// root Command
var EpCmd = common.GetEpCmd(
	cmdName,
	shortDesc,
)

func init() {
	EpCmd.AddCommand(common.GetPrintCmd(cmdName))
	EpCmd.AddCommand(common.GetRunCmd(cmdName))
}
