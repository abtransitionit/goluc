/*
Copyright © 2025 Amar BELGACEM abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var createSDesc = "Create a local [Helm] chart. A top-level folder with starter files."

// Parent command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: createSDesc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("❌ you must pass exactly 1 arguments, the full path name of the chart, got %d", len(args))
		}
		return nil
	},
	Example: `
	desc --ingnginx ingress-nginx
	desc --cilium    cilium
	desc ingress-nginx --ingngin
	desc cilium        --cilium
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(createSDesc)
		cmd.Help()

		// define resource property from flag
		helmChartPath := args[0]

		// define object from the resource property
		helmChart := helm.HelmChart{FullPath: helmChartPath}

		// create the chart
		output, err := helm.CreateChart(localFlag, "o1u", helmChart, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// output
		list.PrettyPrintTable(output)

	},
}
