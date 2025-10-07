/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"context"
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var listChartSDesc = "list [helm] charts in a chart repository."

var listChartLDesc = listChartSDesc

// root Command
var listChartCmd = &cobra.Command{
	Use:   "listChart",
	Short: listChartSDesc,
	Long:  listChartLDesc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the repository (in the whitelist) to list the charts from, got %d", len(args))
		}
		return nil
	},
	Example: fmt.Sprintf(`
# list chart in repo cilium 
  %[1]s chart list cilium
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listChartSDesc)
		ctx := context.Background()

		// get repo key from args
		repoKey := args[0]

		// check provided name exists in the whitelist and get url
		repoObj, ok := helm.MapHelmRepoReference[repoKey]
		if !ok {
			logger.Errorf("repository '%s' is not in the allowed helm repository whitelist", repoKey)
			return
		}

		// define var needed by cli
		repo := helm.HelmRepo{
			Name: repoObj.Name,
		}
		var err error

		// define cli
		cli, err := repo.ListChart(ctx, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// run cli on local or remote
		var output string
		if localFlag {
			logger.Debugf("running on local helm client: %s", cli)
			output, err = helm.QueryHelm("", cli, logger)
		} else {
			remoteHelmHost := "o1u"
			logger.Debugf("running on remote helm client: %s : %s", remoteHelmHost, cli)
			output, err = helm.QueryHelm(remoteHelmHost, cli, logger)
		}

		if err != nil {
			logger.Errorf("failed to run helm command: %s: %w", cli, err)
			return
		}
		// fmt.Print(output)
		list.PrettyPrintTable(output)

	},
}

func init() {
	listChartCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	listChartCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
