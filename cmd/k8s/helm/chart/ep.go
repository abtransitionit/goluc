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
	chartListCmd.Use = "list"
	chartListCmd.Short = "List all charts of all repo configured in the helm client's config."

	repoListCmd := *repo.DescribeCmd
	repoListCmd.Use = "repo"
	repoListCmd.Short = "List charts of a specific repo configured in the helm client's config."

	// repoListCmd := *repo.DescribeCmd
	// repoListCmd.Use = "list"

	// EpCmd.AddCommand(buildCmd)
	EpCmd.AddCommand(buildCmd)
	EpCmd.AddCommand(kindCmd)
	EpCmd.AddCommand(sListCmd)
	EpCmd.AddCommand(&chartListCmd)
	EpCmd.AddCommand(&repoListCmd)
	EpCmd.AddCommand(ReadmeCmd)
	EpCmd.AddCommand(pushCmd)
	// EpCmd.AddCommand(ListCmd)
	// EpCmd.AddCommand(&repoListCmd)
}
