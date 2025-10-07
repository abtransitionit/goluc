/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package node

import (
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var describeSDesc = "list k8s nodes."
var describeLDesc = describeSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: describeSDesc,
	Long:  describeLDesc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the node, got %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get node:name key from args
		nodeName := args[0]

		// define cli
		cli, err := kubectl.Resource{Type: "node", Name: nodeName}.Describe()
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// run cli on local or remote
		var output string
		if localFlag {
			logger.Debugf("running on local kubectl client: %s", cli)
			output, err = helm.QueryHelm("", cli, logger)
		} else {
			remoteHelmHost := "o1u"
			logger.Debugf("running on remote kubectl client: %s : %s", remoteHelmHost, cli)
			output, err = helm.QueryHelm(remoteHelmHost, cli, logger)
		}

		if err != nil {
			logger.Errorf("failed to run helm command: %s: %w", cli, err)
			return
		}

		fmt.Println(output)
	},
}

func init() {
	DescribeCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "uses by default the remote Helm client unless the flag is provided (it will use the local Helm client)")
}
