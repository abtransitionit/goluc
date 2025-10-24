/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kernel

import (
	"github.com/abtransitionit/goluc/cmd/osx/kernel/module"
	"github.com/abtransitionit/goluc/cmd/osx/kernel/param"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "manage os kernel resources."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "kernel",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(module.EpCmd)
	EpCmd.AddCommand(param.EpCmd)
}
