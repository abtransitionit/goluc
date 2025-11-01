/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kben

import common "github.com/abtransitionit/goluc/cmd/workflow/_common"

// root Command
var EpCmd = common.GetEpCmd(
	cmdName,
	shortDesc,
)

func init() {
	EpCmd.AddCommand(common.GetPrintcCmd(cmdName))
	EpCmd.AddCommand(common.GetPrintwCmd(cmdName))
	EpCmd.AddCommand(common.GetPrinttCmd(cmdName))
	EpCmd.AddCommand(common.GetRunCmd(cmdName))
}
