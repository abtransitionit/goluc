/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package htpwd

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete secret from the cluster."
var deleteLDesc = deleteSDesc

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info("use > ... kubectl secret delete")

	},
}
