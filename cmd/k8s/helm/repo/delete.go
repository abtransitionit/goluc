/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete a repo from the helm client config/cache files."
var deleteLDesc = deleteSDesc

// root Command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list configured repos
		// - get instance and operate
		i := helm.Resource{Type: helm.ResRepo}
		output, err := i.List("local", shared.HelmHost, logger)
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
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s ", resName)
		// delete the repo
		// - get instance and operate
		i = helm.Resource{Type: helm.ResRepo, Name: resName}
		output, err = i.Delete("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
