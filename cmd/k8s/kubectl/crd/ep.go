/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package crd

import (
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// var
var localFlag bool
var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage k8s CRDs."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "crd",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(DescribeCmd)
	EpCmd.AddCommand(YamlCmd)
}
