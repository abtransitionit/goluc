/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package mnf

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var ApplySDesc = "apply a yaml or manifest to a cluster."
var ApplyLDesc = ApplySDesc

// root Command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: ApplySDesc,
	Long:  ApplyLDesc,
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
		i := kubectl.Resource{Type: kubectl.ResManifest}
		output, err := i.ListAuth("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// print
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
		// - get instance and operate
		i = kubectl.Resource{Type: kubectl.ResManifest, Name: resName}
		_, err = i.Apply("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
	},
}
