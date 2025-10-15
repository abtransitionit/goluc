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
	"github.com/abtransitionit/gocore/run"
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

		fmt.Println("\nDeleting repo: ", repoName)

		// define cli
		cli, err := helm.HelmRepo{Name: repoName}.Delete()
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// play cli
		output, err = run.ExecuteCliQuery(cli, logger, localFlag, "o1u", helm.HandleHelmError)
		if err != nil {
			logger.Errorf("failed to run command: %s: %w", cli, err)
			return
		}

		list.PrettyPrintTable(output)

	},
}
