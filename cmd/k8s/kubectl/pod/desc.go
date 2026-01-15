/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pod

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var describeSDesc = "list all pods."
var describeLDesc = describeSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: describeSDesc,
	Long:  describeLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list pods
		// - get instance and operate
		output, err := kubectl.List(kubectl.ResPod, "local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
		}

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose pod (enter ID): ")
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
		// define resource property from user choice
		resNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s ", resName)
		// describe pod
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResPod, Name: resName, Ns: resNs}
		output, err = i.Describe("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// print detail
		fmt.Println(output)

	},
}
