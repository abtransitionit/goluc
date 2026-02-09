/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pod

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var iShellSDesc = "get a temporary interactive shell pod."
var iShellLDesc = iShellSDesc

// root Command
var iShellCmd = &cobra.Command{
	Use:   "ishell",
	Short: iShellSDesc,
	Long:  iShellLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list ns
		// - get instance and operate
		output, err := kubectl.List(kubectl.ResNS, "local", HelmHost, logger)
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
		id, err := ui.AskUserInt("\nchoose namespace (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		resNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get res name from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("creating ishell pod in ns: %s ", resNs)
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResPod, Ns: resNs}
		i.CreateIShellTodo("local", shared.HelmHost, logger)

	},
}
