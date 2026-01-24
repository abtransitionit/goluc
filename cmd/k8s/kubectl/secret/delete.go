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
var DeleteSDesc = "delete a secret."
var DeleteLDesc = DeleteSDesc

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: DeleteSDesc,
	Long:  DeleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		// list items
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

		// define resource property from user choice
		resNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get res Ns from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s:%s ", resNs, resName)
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResSecret, Ns: resNs, Name: resName}
		_, err = i.Delete("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

		// log
		logger.Infof("resource still in the cluster for: %s", resName)
		// list items
		// - get instance and operate
		output, err = kubectl.List(kubectl.ResSecret, "local", shared.HelmHost, logger)
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
