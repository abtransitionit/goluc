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

// Description
var readmeSDesc = "view the Readme file of a chart"
var readmeLDesc = readmeSDesc + `
manage following use case:
	- chart is on the local FS where the helm client lives
	- chart is part of a chart repository configured in the helm client configuration files
`

// root Command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: readmeSDesc,
	Long:  readmeLDesc,
	Example: fmt.Sprintf(`
	# view the Readme file of a chart
  %[1]s helm chart readme
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(readmeSDesc)
		// ctx := context.Background()

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// display the list of installed repos
		output, err := helm2.GetRepo("", "").List("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
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
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}
		repoUrl, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get property repo:url from ID: %s > %w", id, err)
			return
		}

		// get instance and operate
		repo := helm2.GetRepo(repoName, repoUrl)
		output, err = repo.ListChart("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// print
		rowCount := list.CountNbLine(output)
		if rowCount == 1 {
			return
		}
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err = ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}
		// define resource property from ID and output
		chartQName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// define/create instance from property
		helmChart := helm2.GetChart("", chartQName, "")

		// get the list and nb of resources kind
		output, err = helmChart.ViewReadme("local", helmHost, logger)
		if err != nil {
			logger.Errorf("failed to view chart kind > %v", err)
			return
		}

	},
}
