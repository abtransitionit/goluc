/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "managing helm release."
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

// func init() {
// 	EpCmd.AddCommand(createCmd)
// 	EpCmd.AddCommand(listCmd)
// 	EpCmd.AddCommand(deleteCmd)
// }
