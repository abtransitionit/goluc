/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"
	"strings"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// flags
var (
	filePathFlag  string
	chartPathFlag string
)

// Description
var kindSDesc = "list the kind of object a chart will create into the K8s cluster."
var kindLDesc = kindSDesc

// root Command
var kindCmd = &cobra.Command{
	Use:   "listKind",
	Short: kindSDesc,
	Long:  kindLDesc,
	Example: fmt.Sprintf(`
  # list the kind of object that will be created in the cluster for the chart: cilium
  %[1]s chart listKind cilium/cilium
  %[1]s chart listKind ~/wkspc/chart/nlos
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(kindSDesc)

		// check parameters
		if chartPathFlag == "" {
			logger.Error("chart localisation is required (--path)")
			return
		}

		// 1 - define the namespace interactively
		// 11 - define the ns interactively - get the list of k8s namespaces
		output, err := kubectl.ListNs(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to list helm repo: %v", err)
			return
		}

		// 12 - print the list
		list.PrettyPrintTable(output)

		// 13 - Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose namespace (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// 14 - define resource property from ID and output

		// 2 - create the release if flag is set to use a chart on the FS
		if chartPathFlag != "" {

			// 21 - define object from the resource property
			helmChart := helm.HelmChart{
				FullName: chartPathFlag,
			}

			// 23 - create the release
			output, err = helm.ListChartKind(localFlag, "o1u", helmChart, logger)
			if err != nil {
				logger.Errorf("failed to create helm release: %v", err)
				return
			}
			// success
			list.PrettyPrintTable(output)
			// logger.Infof("🔹 play > %s", strings.TrimSpace(output))

			// exit
			return
		}

		// create a release using a chart in a repo configured in helm client
		// 1 - get the list of helm repositories
		output, err = helm.ListRepo(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to list helm repo: %v", err)
			return
		}

		// print the list
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err = ui.AskUserInt("\nchoose repo (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		repoName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// define object from the resource property
		helmRepo := helm.HelmRepo{Name: repoName}

		// 2 - get the list of charts in this repo
		output, err = helm.ListChart(localFlag, "o1u", helmRepo, logger)
		if err != nil {
			logger.Errorf("failed to list helm charts: %v", err)
			return
		}

		// print the list
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err = ui.AskUserInt("\nchoose chart (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		chartName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}
		chartVersion, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// define object from the resource property
		helmChart := helm.HelmChart{FullName: chartName, Version: chartVersion, Repo: helmRepo}

		// create the release
		output, err = helm.ListChartKind(localFlag, "o1u", helmChart, logger)
		if err != nil {
			logger.Errorf("failed to create helm release: %v", err)
			return
		}
		// success
		logger.Infof("🔹 play > %s ", strings.TrimSpace(output))

	},
}

func init() {
	kindCmd.Flags().StringVarP(&filePathFlag, "file", "f", "", "The path to the values file")
	kindCmd.Flags().StringVarP(&chartPathFlag, "path", "p", "", "Path to a local unpacked chart directory")
	kindCmd.MarkFlagRequired("name")
}
