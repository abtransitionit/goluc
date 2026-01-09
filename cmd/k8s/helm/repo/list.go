/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	helm2 "github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var listSDesc = "list [helm] repositories (ie. whitelist or installed[default])."
var listLDesc = listSDesc

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Example: fmt.Sprintf(`
# list installed repos
  %[1]s repo list -i

  # list installable repos (whitelist)
  %[1]s repo list -w
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

		// list whitelisted repos
		if showWhitelist {
			logger.Info("list Installable repositories (organization whitelist):")

			// get instance and operate
			output, err := helm2.GetRepo("", "").GetWhitelist("")
			if err != nil {
				logger.Errorf("%w", err)
				return
			}
			// display it
			list.PrettyPrintTable(output)
			return
		}

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// get instance and operate
		output, err := helm2.GetRepo("", "").List("local", helmHost, logger)
		if err != nil {
			logger.Errorf("%w", err)
			return
		}
		// no action is needed based on the number of row
		rowCount := list.CountNbLine(output)
		if rowCount == 1 {
			return
		}
		list.PrettyPrintTable(output)
	},
}

func init() {
	listCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos (default)")
	listCmd.Flags().BoolP("whitelist", "w", false, "show installable Helm repos (organization whitelist helm repos)")
}
