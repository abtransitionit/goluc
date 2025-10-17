/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"fmt"

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
	filePathFlag    string
	chartPathFlag   string
	releaseNameFlag string
	dryRun          bool
	helmRelease     helm.HelmRelease
)

// Description
var createSDesc = "create a [helm] release from a [helm] chart in a [k8s] namespace."
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

		// check parameters
		if releaseNameFlag == "" {
			logger.Error("release name is required (--name)")
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
		k8sNsName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// 2 - create the release if flag is set to use a chart on the FS
		if chartPathFlag != "" {

			// 21 - define object from the resource property
			helmChart := helm.HelmChart{
				FullName: chartPathFlag,
			}
			// 22 - define object from the resource property
			helmRelease = helm.HelmRelease{
				Name:      releaseNameFlag,
				Namespace: k8sNsName,
				Chart:     helmChart,
				ValueFile: filePathFlag,
			}

		}

		// 2 - create the release if flag is not set to use a chart from a repo
		if chartPathFlag == "" {

			// 21 - get the list of helm repositories
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

			// 22 - get the list of charts in this repo
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

			// define object from the resource property
			helmRelease = helm.HelmRelease{
				Name:      releaseNameFlag,
				Repo:      helmRepo,
				Chart:     helmChart,
				Namespace: k8sNsName,
				ValueFile: filePathFlag,
			}

		}
		// 3 - create or dryCreate the release
		if dryRun {
			output, err = helm.DryCreateRelease(localFlag, "o1u", helmRelease, logger)
		} else {
			output, err = helm.CreateRelease(localFlag, "o1u", helmRelease, logger)
		}
		if err != nil {
			logger.Errorf("failed to create helm release: %v", err)
			return
		}
		// print
		list.PrettyPrintTable(output)
		// fmt.Println(output)

	},
}

func init() {
	createCmd.Flags().StringVarP(&filePathFlag, "file", "f", "", "The path to the values file")
	createCmd.Flags().StringVarP(&chartPathFlag, "path", "p", "", "Path to a local unpacked chart directory")
	createCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "dry run the creation of the release")
	createCmd.MarkFlagRequired("name")
	createCmd.Flags().StringVarP(&releaseNameFlag, "name", "n", "", "Name of the Helm release (required)")

}
