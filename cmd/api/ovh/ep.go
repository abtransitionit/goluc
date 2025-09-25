/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "Test some code."
var epLDesc = epSDesc + `
This command is used to run different tests client API(s).
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "ovh",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(playCmd)
	EpCmd.AddCommand(tokenCmd)
}
