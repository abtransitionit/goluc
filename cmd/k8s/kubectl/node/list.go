/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package node

import (
	// kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/spf13/cobra"
)

// Description
var listSDesc = "list all nodes."
var listLDesc = listSDesc

// root Command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get list
		output, err := kubectl.NodeSvc.List("local", HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

	},
}
