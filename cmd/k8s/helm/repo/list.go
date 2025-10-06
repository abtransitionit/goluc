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
var listSDesc = "list helm repositories that exist in the Helm client configuration file."
var listLDesc = listSDesc

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Example: fmt.Sprintf(`
# list installed repos
  %[1]s repo list -I

  # list installable repos (whitelist)
  %[1]s repo list -w

  # list all repos (installed and installable)
  %[1]s repo list -a
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listSDesc)
		ctx := context.Background()

		// Read flags
		showInstalled, _ := cmd.Flags().GetBool("installed")
		showWhitelist, _ := cmd.Flags().GetBool("whitelist")

		// default behavior: show installed
		if !showInstalled && !showWhitelist {
			showInstalled = true
		}

		// list whitelist repos
		if showWhitelist {
			logger.Info("Installable repositories (organization whitelist):")
			for _, r := range helm.MapHelmRepoReference {
				fmt.Printf("- %-10s %s\n", r.Name, r.Url)
			}
			return
		}

		// Here we want to list installed repos - define var needed by cli
		repo := helm.HelmRepo{}

		// define cli
		cli, err := repo.List(ctx, logger)
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

		list.PrettyPrintTable(output)
	},
}

func init() {
	listCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	listCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
