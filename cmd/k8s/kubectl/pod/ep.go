/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pod

import (
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// var
var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage pods."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "pod",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(EventCmd)
	EpCmd.AddCommand(YamlCmd)
}
