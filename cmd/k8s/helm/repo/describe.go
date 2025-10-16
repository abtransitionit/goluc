/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var describeSDesc = "list [helm] charts in a chart repository."

var describeLDesc = describeSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: describeSDesc,
	Long:  describeLDesc,
	Example: fmt.Sprintf(`
  # list chart in repo cilium 
  %[1]s chart list cilium
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(describeSDesc)
		// ctx := context.Background()

		// get list of installed repos
		output, err := helm.ListRepo(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose repo (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		repoName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// create a Helm repo object
		helmRepo := helm.HelmRepo{Name: repoName}

		// List helm charts
		output, err = helm.ListChart(localFlag, "o1u", helmRepo, logger)
		if err != nil {
			logger.Errorf("failed to list helm charts: %v", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

	},
}

func init() {
	DescribeCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	DescribeCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
