/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var srcListSDesc = "list buildable local source chart folder"
var srcListLDesc = srcListSDesc

// root Command
var listSCmd = &cobra.Command{
	Use:   "lists",
	Short: srcListSDesc,
	Long:  srcListLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s build add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(srcListSDesc)
		// ctx := context.Background()

		// get instance and operate
		i := helm.Resource{Type: helm.ResChart, SType: helm.STypeChartBuild}
		output, err := i.List("local", "local", logger)
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

	},
}
