/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pkg

import (
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var addSDesc = "add a native os package either from a specific or standard repository."
var addLDesc = addSDesc

// root Command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: addSDesc,
	Long:  addLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s repo add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(addSDesc)
		// ctx := context.Background()

		// list repos in whitelist
		output, _ := helm.ListRepoReferenced(false, "", logger)

		// print it
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose repo (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		repoName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}
		repoUrl, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// define object from the resource property
		helmRepo := helm.HelmRepo{Name: repoName, Url: repoUrl}

		// Add repo
		output, err = helm.AddRepo(localFlag, "o1u", helmRepo, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		list.PrettyPrintTable(output)

	},
}
