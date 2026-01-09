/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	helm2 "github.com/abtransitionit/golinux/mock/k8scli/helm"
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

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// get instance and operate to get list repo
		output, err := helm2.GetRepo("", "").List("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
			return
		}
		// no action is needed based on the number of row
		rowCount := list.CountNbLine(output)
		if rowCount == 1 {
			return
		}
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose item (enter ID): ")
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

		// get instance and operate
		i := helm2.GetRepo(repoName, "")
		output, err = i.ListChart("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
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
