/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"github.com/spf13/cobra"
)

var localFlag bool

// Description
var epSDesc = "managing linux os package repositories."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "repo",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local config files if the flag is set; otherwise, use the remote config files")
	EpCmd.AddCommand(addCmd)
	EpCmd.AddCommand(listCmd)
	EpCmd.AddCommand(deleteCmd)
}
