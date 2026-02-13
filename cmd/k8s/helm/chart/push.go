/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var pushSDesc = "push a chart's artifact to an OCI registry"
var pushLDesc = pushSDesc + `
manage following use case:
	- the chart artifact (targz) is on the local FS where the helm client is installed
`

// root Command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: pushSDesc,
	Long:  pushLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s build add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(pushSDesc)
		// ctx := context.Background()

		// - get instance and operate
		i := helm.Resource{Type: helm.ResChart, SType: helm.STypeChartBuild}
		output, err := i.List("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
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

		logger.Infof("pushing chart: %s", resName)

		// // 3- get instance and operate
		// i := helm.Resource{Type: helm.ResChart, Name: resName, Param: param}
		// _, err := i.Build("local", "local", logger)
		// if err != nil {
		// 	logger.Errorf("%v", err)
		// 	return
		// }

	},
}
