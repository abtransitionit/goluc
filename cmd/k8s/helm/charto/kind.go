/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	helm2 "github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// flags
var (
	filePathFlag  string
	chartPathFlag string
)

// Description
var kindSDesc = "list the kind of K8s resources a chart will create into the K8s cluster."
var kindLDesc = kindSDesc + `
manage following use case:
	- chart is on the local FS where the helm client lives
	- chart is part of a chart repository configured in the helm client configuration files
`

// root Command
var kindCmd = &cobra.Command{
	Use:   "kind",
	Short: kindSDesc,
	Long:  kindLDesc,
	Example: fmt.Sprintf(`
  # list the kind of object that will be created in the cluster for the chart: cilium
  %[1]s chart listKind cilium/cilium
  %[1]s chart listKind ~/wkspc/chart/nlos
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		var helmChart helm2.Chart
		var helmRepo helm2.Repo

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(kindSDesc)

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// 1 - check if flag is set to use a chart on the FS
		if chartPathFlag != "" {

			// 11 - define/create instance from the property
			helmChart = helm2.Chart{FullName: chartPathFlag}

		} else {
			// 2 - define/create instance
			// 21 - get the list of helm installed repo
			// get instance and operate
			output, err := helm2.GetRepo("", "").List("local", helmHost, logger)
			if err != nil {
				logger.Errorf("%w", err)
				return
			}

			// print the list
			list.PrettyPrintTable(output)

			// Ask user which ID (to choose) from the printed list
			id, err := ui.AskUserInt("\nchoose repo (enter ID): ")
			if err != nil {
				logger.Errorf("invalid ID: %v", err)
				return
			}

			// 22 - define resource property from ID and output
			repoName, err := list.GetFieldByID(output, id, 0)
			if err != nil {
				logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
				return
			}

			// define object from the resource property
			helmRepo = helm2.Repo{Name: repoName}

			// 23 - get the list of charts in this repo
			output, err = helmRepo.ListChart("local", helmHost, logger)
			if err != nil {
				logger.Errorf("failed to list helm charts: %v", err)
				return
			}

			// print the list
			list.PrettyPrintTable(output)

			// 24 - Ask user which ID (to choose) from the list
			id, err = ui.AskUserInt("\nchoose chart (enter ID): ")
			if err != nil {
				logger.Errorf("invalid ID: %v", err)
				return
			}

			// 25 - define resource property from ID and output
			chartName, err := list.GetFieldByID(output, id, 0)
			if err != nil {
				logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
				return
			}

			// 26 - define/create instance from property
			helmChart = helm2.Chart{FullName: chartName, Repo: &helmRepo}
		}
		// 3 - get the list and nb of resources kind
		output, err := helmChart.ListResKind("local", helmHost, logger)
		if err != nil {
			logger.Errorf("failed to list chart kind: %w", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

		// 4 - get the list of resources kind and name
		output, err = helmChart.ListRes("local", helmHost, logger)
		if err != nil {
			logger.Errorf("failed to list chart kind: %v", err)
			return
		}
		// print
		list.PrettyPrintTable(output)

	},
}

func init() {
	kindCmd.Flags().StringVarP(&filePathFlag, "file", "f", "", "The path to the values file")
	kindCmd.Flags().StringVarP(&chartPathFlag, "path", "p", "", "Path to a local unpacked chart directory")
}
