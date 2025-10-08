/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package node

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
	"github.com/spf13/cobra"
)

// Description
var yamlSDesc = "display a single node manifest."
var yamlLDesc = yamlSDesc

// root Command
var YamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: yamlSDesc,
	Long:  yamlLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get list
		output, err := kubectl.ListNode(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// print list
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
		fmt.Println(id, output)

		// get resource property from ID and output
		nodeName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// define cli
		cli, err := kubectl.Resource{Type: "node", Name: nodeName}.Yaml()
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

		fmt.Println(output)
	},
}

func init() {
	YamlCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "uses by default the remote Helm client unless the flag is provided (it will use the local Helm client)")
}
