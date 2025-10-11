/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package list

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pod"
	"github.com/spf13/cobra"
)

// var
var localFlag bool

// Description
var epSDesc = "list k8s resources."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "list",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	nodeListCmd := *node.ListCmd
	nodeListCmd.Use = "node"
	nodeListCmd.Short = "List nodes"

	podListCmd := *pod.ListCmd
	podListCmd.Use = "pod"
	podListCmd.Short = "List pods"

	nsListCmd := *ns.ListCmd
	nsListCmd.Use = "ns"
	nsListCmd.Short = "List namespaces"

	EpCmd.AddCommand(&nodeListCmd)
	EpCmd.AddCommand(&podListCmd)
	EpCmd.AddCommand(&nsListCmd)
}
