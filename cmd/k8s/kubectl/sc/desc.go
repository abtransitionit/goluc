/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package sc

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var describeSDesc = "display single node details."
var describeLDesc = describeSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: describeSDesc,
	Long:  describeLDesc,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) != 1 {
	// 		return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the node, got %d", len(args))
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list nodes
		// - get instance and operate
		output, err := kubectl.List(kubectl.ResSC, "local", shared.HelmHost, logger)
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
		id, err := ui.AskUserInt("\nchoose node (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		resName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get res name from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s ", resName)
		// describe instance
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResSC, Name: resName}
		output, err = i.Describe("local", shared.HelmHost, logger)
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
	},
}
