/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
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

		// list installed release
		// - get instance and operate
		i := helm.Resource{Type: helm.ResRelease}
		output, err := i.List("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
		}
		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		resName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}
		resNs, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}
		// log
		logger.Infof("selected item: %s / %s", resName, resNs)
		// list manifest
		// - get instance and operate
		i = helm.Resource{Type: helm.ResRelease, Name: resName, Namespace: resNs}
		output, err = i.Describe("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			fmt.Println(output)
		}

	},
}

func init() {
	describeCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	describeCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
