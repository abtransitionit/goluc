/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cm

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage a ConfigMap"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "cm",
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
