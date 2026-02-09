/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ishell

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "get an interactive tmp shell pod"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "ishell",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// get logger
		logger := logx.GetLogger()
		logger.Info("get an interactive tmp shell pod")
	},
}
