/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gox

import (
	"github.com/abtransitionit/goluc/cmd/osx/gox/luca"
	"github.com/abtransitionit/goluc/cmd/osx/gox/lucm"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "build linux CLI from Go source code."
var epLDesc = epSDesc + `
This command allows to Build a CLI binary for any OS from GO source code.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "go",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(luca.EpCmd)
	EpCmd.AddCommand(lucm.EpCmd)
}
