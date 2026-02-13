/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package sa

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage a ServiceAccounts"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "sa",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(DescribeCmd)
	// EpCmd.AddCommand(YamlCmd)
}
