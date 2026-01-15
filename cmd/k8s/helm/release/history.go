/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var historySDesc = "list the installation history of a release."
var historyLDesc = historySDesc

// root Command
var historyCmd = &cobra.Command{
	Use:   "history",
	Short: historySDesc,
	Long:  historyLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(historySDesc)
		// ctx := context.Background()

		// list installed release
		// - get instance and operate
		i := helm.Resource{Type: helm.ResRelease}
		output, err := i.List("local", shared.HelmHost, logger)
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
		resName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property from ID: %s > %v", id, err)
			return
		}
		resNs, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get property from ID: %s > %v", id, err)
			return
		}
		// log
		logger.Infof("selected item: %s / %s", resName, resNs)

		// list revision history
		// - get instance and operate
		i = helm.Resource{Type: helm.ResRelease, Name: resName, Namespace: resNs}
		output, err = i.ListHistory("local", shared.HelmHost, logger)
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
		id, err = ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		resRevision, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property from ID: %s > %v", id, err)
			return
		}
		// log
		logger.Infof("selected item: %s / %s / %s", resName, resNs, resRevision)
		// list revision detail
		// - get instance and operate
		i = helm.Resource{Type: helm.ResRelease, Name: resName, Namespace: resNs, Revision: resRevision}
		output, err = i.Detail("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		fmt.Println(output)

	},
}
