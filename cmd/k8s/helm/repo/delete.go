/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

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
var deleteSDesc = "delete a repo."
var deleteLDesc = deleteSDesc + `
- This command delete the helm repo by just updating the Helm client configuration file in the user's home directory.
`

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Example: fmt.Sprintf(`
  # add helm repo
  %[1]s repo add myrepo https://charts.bitnami.com/bitnami	
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(deleteSDesc)
		// ctx := context.Background()

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// get instance and operate
		output, err := helm2.GetRepo("", "").List("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// no action is needed based on the number of row
		rowCount := list.CountNbLine(output)
		if rowCount == 1 {
			logger.Warn("no item to delete")
			return
		}

		// print
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
			logger.Errorf("failed to get property repo:name from ID: %s >  %w", id, err)
			return
		}

		// get helm host
		helmHost, err = helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// get instance and operate
		i := helm2.GetRepo(repoName, "")
		output, err = i.Delete("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

	},
}
