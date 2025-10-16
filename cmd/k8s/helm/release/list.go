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
var listSDesc = "list all helm releases in the cluster."
var listLDesc = listSDesc

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listSDesc)
		// ctx := context.Background()

		// get list
		output, err := helm.ListRelease(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// no action is needed based on the number of row
		rowCount := list.CountNbLine(output)
		if rowCount == 1 {
			logger.Warn("no item to list")
			return
		}

		list.PrettyPrintTable(output)
	},
}

func init() {
	listCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos (default)")
	listCmd.Flags().BoolP("whitelist", "w", false, "show installable Helm repos (organization whitelist helm repos)")
}
