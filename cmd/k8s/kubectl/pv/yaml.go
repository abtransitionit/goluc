/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pv

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
var yamlSDesc = "get the yaml manifest for a CRD."
var yamlLDesc = yamlSDesc

// root Command
var YamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: yamlSDesc,
	Long:  yamlLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get list
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
			logger.Errorf("failed to get item name from ID: %s: %v", id, err)
			return
		}
		// define resource property from user choice
		resNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get item ns from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s ", resName)
		// yaml cm
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResPv, Name: resName, Ns: resNs}
		output, err = i.GetYaml("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		fmt.Println(output)

	},
}
