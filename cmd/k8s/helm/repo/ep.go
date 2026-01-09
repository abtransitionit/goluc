/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"github.com/spf13/cobra"
)

var localFlag bool

// Description
var epSDesc = "managing helm repositories."
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
	// EpCmd.PersistentFlags().BoolVarP(&remoteFlag, "remote", "r", false, "uses by default the local Helm client unless the flag is provided (it will use the remote Helm client)")
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
	EpCmd.AddCommand(addCmd)
	EpCmd.AddCommand(listCmd)
	// EpCmd.AddCommand(list2Cmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(deleteCmd)
}
