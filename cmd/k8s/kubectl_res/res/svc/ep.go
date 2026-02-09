/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package svc

import (
	// "github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage StorageClasses (related to PVC)."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "svc",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(PortFwdCmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(DeleteCmd)
	EpCmd.AddCommand(YamlCmd)
}
