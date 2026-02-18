/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"github.com/abtransitionit/goluc/cmd/k8s/helm/repo"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage helm chart."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "chart",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	chartListCmd := *repo.ChartCmd
	chartListCmd.Use = "listr"
	chartListCmd.Short = (*repo.ChartCmd).Short

	repoListCmd := *repo.DescribeCmd
	repoListCmd.Short = (*repo.DescribeCmd).Short
	repoListCmd.Use = "listc"

	// repoListCmd := *repo.DescribeCmd
	// repoListCmd.Use = "list"

	// EpCmd.AddCommand(buildCmd)
	EpCmd.AddCommand(buildCmd)
	EpCmd.AddCommand(pullCmd)
	EpCmd.AddCommand(kindCmd)
	EpCmd.AddCommand(listSCmd)
	EpCmd.AddCommand(&chartListCmd)
	EpCmd.AddCommand(&repoListCmd)
	EpCmd.AddCommand(ReadmeCmd)
	EpCmd.AddCommand(pushCmd)
	// EpCmd.AddCommand(ListCmd)
	// EpCmd.AddCommand(&repoListCmd)
}
