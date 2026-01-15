/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

var localFlag bool
var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage helm release."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "release",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(epSDesc)
		cmd.Help()
	},
}

func init() {
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
	// EpCmd.AddCommand(createCmd)
	EpCmd.AddCommand(listCmd)
	EpCmd.AddCommand(describeCmd)
	EpCmd.AddCommand(deleteCmd)
	EpCmd.AddCommand(historyCmd)
}

// go run . k8s helm  release create  -d --name kbe-toto -p '~/wkspc/chart/mxtest' -f '~/wkspc/chart/myval.yaml'
// go run . k8s helm  release list
// go run . k8s helm  release delete
