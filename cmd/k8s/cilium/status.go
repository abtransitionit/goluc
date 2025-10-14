/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cilium

import (
	cilium "github.com/abtransitionit/gocore/k8s-cilium"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var statusSDesc = "display the status of CNI:cilium components."

var statusLDesc = statusSDesc

// root Command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: statusSDesc,
	Long:  statusLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()

		// display status
		output, err := cilium.DisplayStatus(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// fmt.Println(output)
		list.PrettyPrintTable(output)

	},
}

func init() {
	statusCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	statusCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
