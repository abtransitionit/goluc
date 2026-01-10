/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

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
var describeSDesc = "display all Kubernetes resources that were uploaded to the server for a release."

var describeLDesc = describeSDesc

// root Command
var describeCmd = &cobra.Command{
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

		// get list of installed releases
		output, err := helm.HelmRelease{}.List(localFlag, "o1u", logger)
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

		// print
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose repo (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		releaseName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}
		releaseK8sNs, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// create the object
		helmRelease := helm.HelmRelease{Name: releaseName, Namespace: releaseK8sNs}

		// operate on this object
		output, err = helmRelease.Describe(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to list helm charts: %v", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

	},
}

func init() {
	describeCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	describeCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
