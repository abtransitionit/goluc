/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package mnf

import (
	// "github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage manifest [yaml] files."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "mnf",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(applyCmd)
	EpCmd.AddCommand(kindCmd)
	EpCmd.AddCommand(deleteCmd)
	EpCmd.AddCommand(descCmd)
}
