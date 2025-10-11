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

		// list repos in whitelist
		output, _ := helm.ListRepoReferenced(false, "", logger)

		// print it
		list.PrettyPrintTable(output)

		// Ask user which ID to describe
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

		// get resource property from ID and output
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

		fmt.Println("\nAdding repo: ", repoName)

		// define cli
		cli, err := helm.HelmRepo{Name: repoName, Url: repoUrl}.Add()
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
