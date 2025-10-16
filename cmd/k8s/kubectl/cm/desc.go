/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cm

import (
	"fmt"

	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/spf13/cobra"
)

// Description
var describeSDesc = "display details for a single ConfigMap."
var describeLDesc = describeSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: describeSDesc,
	Long:  describeLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get list
		output, err := kubectl.ListCm(localFlag, "o1u", logger)
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

		// define object properties from ID and output
		cmName, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}
		cmNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// define an object from its properties
		cm := kubectl.Resource{Type: "cm", Name: cmName, Ns: cmNs}

		// operate on this object
		output, err = kubectl.DescribeCm(localFlag, "o1u", cm, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		fmt.Println(output)
	},
}
