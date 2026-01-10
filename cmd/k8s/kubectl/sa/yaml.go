/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package sa

import (
	"fmt"

	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/spf13/cobra"
)

// Description
var yamlSDesc = "display manifest for a single ServiceAccount."
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
		output, err := kubectl.ListSa(localFlag, "o1u", logger)
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

		// define resource property from user choice
		saName, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}
		saNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// define object from property
		sa := kubectl.Resource{Type: "sa", Name: saName, Ns: saNs}

		// get detail
		output, err = kubectl.YamlSa(localFlag, "o1u", sa, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		fmt.Println(output)
	},
}
