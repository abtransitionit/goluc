/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cilium

import (
	cilium "github.com/abtransitionit/gocore/k8s-cilium"
	helm "github.com/abtransitionit/gocore/k8s-helm"
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

		// define cli
		var err error
		cli, err := cilium.Status()
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
	statusCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	statusCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
