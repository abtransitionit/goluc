/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete a helm repo."
var deleteLDesc = deleteSDesc

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		// logger.Info(createSDesc)

		// list installed release
		// - get instance and operate
		i := helm.Resource{Type: helm.ResRelease}
		output, err := i.List("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
		}

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nitem do you want to delete (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}
		// defrine resource property from ID and output
		resName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}
		resNs, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}
		// log
		logger.Infof("selected item: %s ", resName)
		// log
		logger.Infof("deleting release: %s from namespace: %s", resName, resNs)
		// delete release
		// - get instance and operate
		i = helm.Resource{Type: helm.ResRelease, Name: resName, Namespace: resNs}
		output, err = i.Delete("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// log
		logger.Infof("deleted release: %s from namespace: %s", resName, resNs)

	},
}
