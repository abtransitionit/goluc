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
var listChartSDesc = "list [helm] charts in a chart repository."

var listChartLDesc = listChartSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: listChartSDesc,
	Long:  listChartLDesc,
	Example: fmt.Sprintf(`
# list chart in repo cilium 
  %[1]s chart list cilium
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listChartSDesc)
		// ctx := context.Background()

		// get list of installed repos
		output, err := helm.ListRepo(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// print
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

		// define cli
		cli, err := helm.HelmRepo{Name: repoName}.ListChart()
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

func init() {
	DescribeCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	DescribeCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
