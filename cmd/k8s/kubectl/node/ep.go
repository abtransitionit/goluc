/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package node

import (
	// "github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage nodes."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "node",
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
