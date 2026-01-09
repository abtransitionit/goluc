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

// // flags
// var (
// 	filePathFlag  string
// 	chartPathFlag string
// )

// Description
var ReadmeSDesc = "view the Readme file of a chart"
var ReadmeLDesc = ReadmeSDesc + `
manage following use case:
	- chart is on the local FS where the helm client lives
	- chart is part of a chart repository configured in the helm client configuration files
`

// root Command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: ReadmeSDesc,
	Long:  ReadmeLDesc,
	Example: fmt.Sprintf(`
  # lview the Readme file of a chart
  %[1]s chart listKind cilium/cilium
  %[1]s chart listKind ~/wkspc/chart/nlos
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		var helmChart helm2.Chart
		var helmRepo helm2.Repo

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(ReadmeSDesc)

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
			// 21 - get the list of helm repo
			output, err := helmChart.ViewReadme("local", helmHost, logger)
			if err != nil {
				logger.Errorf("failed to list helm repo: %v", err)
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
		output, err := helmChart.ViewReadme("local", helmHost, logger)
		if err != nil {
			logger.Errorf("failed to list chart kind: %v", err)
			return
		}

		// print
		fmt.Printf("%s\n", output[:10])

	},
}

func init() {
	readmeCmd.Flags().StringVarP(&filePathFlag, "file", "f", "", "The path to the values file")
	readmeCmd.Flags().StringVarP(&chartPathFlag, "path", "p", "", "Path to a local unpacked chart directory")
}
