/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package deploy

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
var deleteSDesc = "delete Deployment"
var deleteLDesc = deleteSDesc

// root Command
var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list cm
		// - get instance and operate
		output, err := kubectl.List(kubectl.ResPv, "local", shared.HelmHost, logger)
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
		// // define resource property from user choice
		// resNs, err := list.GetFieldByID(output, id, 0)
		// if err != nil {
		// 	logger.Errorf("failed to get res ns from ID: %s: %v", id, err)
		// 	return
		// }

		// log
		logger.Infof("selected item: %s ", resName)
		// describe cm
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResPv, Name: resName}
		output, err = i.Delete("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		fmt.Println(output)
	},
}
