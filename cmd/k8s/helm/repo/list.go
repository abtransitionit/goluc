/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var listSDesc = "list installed helm repositories (ie. that exist in the Helm client configuration file)."
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
		// ctx := context.Background()

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

		// Here we want to list installed repos - get list
		output, err := helm.ListRepo(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		list.PrettyPrintTable(output)
	},
}

func init() {
	listCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	listCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
