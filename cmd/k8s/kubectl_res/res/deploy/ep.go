/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package deploy

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage a Deployment"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "deploy",
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
