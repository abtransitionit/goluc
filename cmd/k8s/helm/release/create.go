/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

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
var filePathFlag string

// Description
var createSDesc = "create a [helm] release from a [helm] chart."
var createLDesc = createSDesc

// root Command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: createSDesc,
	Long:  createLDesc,
	Example: fmt.Sprintf(`
	# create the release kbe-cilium-prod from chart named cilium in repo named cilium 
  %[1]s chart list cilium
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(createSDesc)

		// 1 - get the list of helm repositories
		output, err := helm.ListRepo(localFlag, "o1u", logger)
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

		// 3 - get the list of k8s namespaces
		output, err = kubectl.ListNs(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to list helm repo: %v", err)
			return
		}

		// print the list
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err = ui.AskUserInt("\nchoose namespace (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		k8sNsName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// Ask user which ID (to choose) from the printed list
		helmReleaseName := ui.AskUser("\ndefine the release name: ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// 4 - define object from the resource property
		helmRelease := helm.HelmRelease{
			Name:      helmReleaseName,
			Repo:      helmRepo,
			Chart:     helmChart,
			Namespace: k8sNsName,
			ValueFile: filePathFlag,
		}

		// 5 - create release in k8s
		output, err = helm.CreateRelease(localFlag, "o1u", helmRelease, logger)
		if err != nil {
			logger.Errorf("failed to create helm release: %v", err)
			return
		}

		fmt.Println("🔹 play > ", strings.TrimSpace(output))

	},
}

func init() {
	createCmd.Flags().StringVarP(&filePathFlag, "file", "f", "", "The path to the values file")
}
