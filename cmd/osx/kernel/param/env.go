/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package param

import (
	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
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

		// play cli
		output, err := run.ExecuteCliQuery(cli, logger, localFlag, "o1u", helm.HandleHelmError)
		if err != nil {
			logger.Errorf("failed to run command: %s: %w", cli, err)
			return
		}

		if err != nil {
			logger.Errorf("failed to run helm command: %s: %w", cli, err)
			return
		}
		list.PrettyPrintKvpair(output)

	},
}
