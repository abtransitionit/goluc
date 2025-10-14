/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete a helm repo."
var deleteLDesc = deleteSDesc

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(createSDesc)

		// list release
		output, err := helm.ListRelease(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to list helm repo: %v", err)
			return
		}

		// print the list
		list.PrettyPrintTable(output)
	},
}
