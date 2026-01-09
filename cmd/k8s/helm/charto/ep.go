/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/k8s/helm/repo"
	"github.com/spf13/cobra"
)

// var localFlag bool

// Description
var epSDesc = "managing helm chart."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "chart",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(epSDesc)
		cmd.Help()
	},
}

func init() {
	repoListCmd := *repo.DescribeCmd
	repoListCmd.Use = "list"

	// EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
	// EpCmd.AddCommand(describeCmd)
	// EpCmd.AddCommand(createCmd)
	EpCmd.AddCommand(&repoListCmd)
	EpCmd.AddCommand(valueCmd)
	EpCmd.AddCommand(kindCmd)
	EpCmd.AddCommand(readmeCmd)
}
