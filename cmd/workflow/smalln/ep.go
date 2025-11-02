/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package smalln

import (
	"github.com/abtransitionit/golinux/onpm"
	common "github.com/abtransitionit/goluc/cmd/workflow/_common"
)

// root Command
var EpCmd = common.GetEpCmd(
	cmdName,
	shortDesc,
)

func init() {
	// sub cde
	EpCmd.AddCommand(common.GetPrintCmd(cmdName))
	EpCmd.AddCommand(common.GetRunCmd(cmdName))

	// function mapping
	// Registry.Add("vm.CheckVmSshAccess", vm.CheckVmSshAccess)
	common.FunctionRegistry.Add("onpm.UpgradeOs", onpm.UpgradeOs)
	common.FunctionRegistry.Add("onpm.UpgradePkg", onpm.UpgradePkg)
}
