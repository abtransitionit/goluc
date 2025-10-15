/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package node

import (
	"fmt"

	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
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

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose node (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		nodeName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// define object from property
		node := kubectl.Resource{Type: "node", Name: nodeName}

		// get detail
		output, err = kubectl.YamlNode(localFlag, "o1u", node, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		fmt.Println(output)
	},
}
