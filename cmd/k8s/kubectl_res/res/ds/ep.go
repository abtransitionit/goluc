/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ds

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage Deployment"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "ds",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(DeleteCmd)
	EpCmd.AddCommand(YamlCmd)
}
