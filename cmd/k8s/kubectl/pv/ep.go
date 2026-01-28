/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pv

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage CRDs."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "pv",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(YamlCmd)
}
