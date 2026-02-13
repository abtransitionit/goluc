/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package registry

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage helm registry."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "registry",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(loginCmd)
	EpCmd.AddCommand(logoutCmd)
}
