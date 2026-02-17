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
var pullSDesc = "download a chart's artifact from an OCI registry"
var pullLDesc = pullSDesc + `
manage following use case:
	- the chart artifact (targz) is on the local FS where the helm client is installed
`

// root Command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: pullSDesc,
	Long:  pullLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s build add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(pullSDesc)
		// ctx := context.Background()

		// - get instance and operate
		i := helm.Resource{Type: helm.ResRegistry}
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
		registryName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}

		// - get instance and operate
		i = helm.Resource{Type: helm.ResChart, SType: helm.STypeChartBuild}
		output, err = i.List("local", "local", logger)
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
		id, err = ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		chartName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}
		// create param map
		param := make(map[string]string)
		param["registry"] = registryName

		// - get instance and operate
		logger.Infof("pushing chart artifact %q to registry %q", chartName, registryName)
		i = helm.Resource{Type: helm.ResChart, SType: helm.STypeChartBuild, Name: chartName, Param: param}
		err = i.Pull("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
