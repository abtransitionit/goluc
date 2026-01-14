/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cm

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
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
		output, err := kubectl.List(kubectl.ResCM, "local", HelmHost, logger)
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
		resName, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// get instance
		logger.Infof("ns name: %s", resName)
		i := kubectl.Resource{Type: kubectl.ResCM, Name: resName}

		// get detail
		output, err = i.Describe("local", HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to describe resource: %v", err)
			return
		}

		// print detail
		fmt.Println(output)
	},
}
