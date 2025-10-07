/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package conf

import (
	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var envSDesc = "list helm envars."
var envLDesc = envSDesc

// root Command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: envSDesc,
	Long:  envLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(epSDesc)
		// define cli
		cli, err := helm.GetEnv()
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
		list.PrettyPrintKvpair(output)

	},
}
