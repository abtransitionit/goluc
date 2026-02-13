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
var buildSDesc = "build a chart's artifact from it source folder"
var buildLDesc = buildSDesc + `
manage following use case:
	- the chart repository is on the local FS where the helm client is installed
`

// root Command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: buildSDesc,
	Long:  buildLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s build add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(buildSDesc)
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

		// list possible chart to build
		// get instance and operate
		i = helm.Resource{Type: helm.ResChart, SType: helm.STypeChartBuild, Name: resName}
		err = i.Build("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
