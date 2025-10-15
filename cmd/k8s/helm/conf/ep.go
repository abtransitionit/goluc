/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package conf

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var forceFlag bool
var localFlag bool

// Description
var epSDesc = "displaying helm configuration."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "conf",
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
	EpCmd.AddCommand(envCmd)
	// EpCmd.AddCommand(fileCmd)
	// EpCmd.AddCommand(versionCmd)
}
