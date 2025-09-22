/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package api

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "Test some code."
var epLDesc = epSDesc + `
This command is used to run different tests client API(s).
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "api",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(playCmd)
}
