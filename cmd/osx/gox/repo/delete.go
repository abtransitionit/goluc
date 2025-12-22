/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete a native os package repository that was previously installed from the whitelist."
var deleteLDesc = deleteSDesc

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

		// get list of installed repos
		output, err := helm.ListRepo(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
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
		fmt.Print("\nWhich item do you want to describe (enter ID): ")

		// convert user input to int
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		id, err := strconv.Atoi(input)
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		repoName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// define object from the resource property
		helmRepo := helm.HelmRepo{Name: repoName}

		// delete the repo
		output, err = helm.DeleteRepo(localFlag, "o1u", helmRepo, logger)
		if err != nil {
			logger.Errorf("failed to list helm charts: %v", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

	},
}
