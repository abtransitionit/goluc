/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package secret

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var DescribeSDesc = "display details for a single secret."
var DescribeLDesc = DescribeSDesc

// root Command
var describeCmd = &cobra.Command{
	Use:   "desc",
	Short: DescribeSDesc,
	Long:  DescribeLDesc,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) != 1 {
	// 		return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the node, got %d", len(args))
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list authorized manifest
		// - get instance and operate
		output, err := kubectl.List(kubectl.ResSecret, "local", shared.HelmHost, logger)
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
		id, err := ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		resName, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get res name from ID: %s: %v", id, err)
			return
		}
		resNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get res name from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s ", resName)
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResSecret, Name: resName, Ns: resNs}
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
