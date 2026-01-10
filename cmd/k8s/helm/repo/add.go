/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"
	"log"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	helm2 "github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var addSDesc = "add a helm repo from the whitelist (ie. authorized Helm repo)."
var addLDesc = addSDesc + `
- This command add the helm repo by just updating the Helm client configuration file in the user's home directory.
- If the repository name is not already in the Helm configuration file, it adds it.
- If the repository name is already in the Helm configuration file, it updates the URL of the repository.
`

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

		// get the yaml as a printable string
		output, err := helm2.RepoSvc.GetWhitelist("")
		if err != nil {
			log.Fatal(err)
			return
		}
		// display it
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
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

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}
		// get instance from resource property and operate
		i := helm2.GetRepo(repoName, repoUrl)
		if err := i.Add("local", helmHost, logger); err != nil {
			logger.Errorf("%w", err)
			return
		}

		// get instance and operate
		output, err = helm2.RepoSvc.List("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
			return
		}
		// no action is needed based on the number of row
		rowCount := list.CountNbLine(output)
		if rowCount == 1 {
			return
		}
		list.PrettyPrintTable(output)

	},
}
