/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	helm2 "github.com/abtransitionit/golinux/mock/k8scli/helm"
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

		// get stateless instance and operate
		output, err := helm2.ReleaseSvc.List("local", HelmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
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
