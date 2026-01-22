/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage [helm] repos on a helm client."
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
	EpCmd.AddCommand(AddCmd)
	EpCmd.AddCommand(DeleteCmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(ChartCmd)
	// EpCmd.AddCommand(DescribeCmd)
	// EpCmd.AddCommand(YamlCmd)
}
