/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var forceFlag bool
var localFlag bool

// Description
var epSDesc = "managing helm chart."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "chart",
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
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "uses by default the remote Helm client unless the flag is provided (it will use the local Helm client)")
	EpCmd.AddCommand(describeCmd)
	EpCmd.AddCommand(listCmd)
	EpCmd.AddCommand(valueCmd)
}
