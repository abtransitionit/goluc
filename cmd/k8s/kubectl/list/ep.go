/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package list

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/cm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sa"
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
	crdListCmd := *crd.ListCmd
	crdListCmd.Use = "crd"
	crdListCmd.Short = "List CRDs"

	nodeListCmd := *node.ListCmd
	nodeListCmd.Use = "node"
	nodeListCmd.Short = "List nodes"

	podListCmd := *pod.ListCmd
	podListCmd.Use = "pod"
	podListCmd.Short = "List pods"

	nsListCmd := *ns.ListCmd
	nsListCmd.Use = "ns"
	nsListCmd.Short = "List namespaces"

	saListCmd := *sa.ListCmd
	saListCmd.Use = "sa"
	saListCmd.Short = "List serviceAccounts"

	cmListCmd := *cm.ListCmd
	cmListCmd.Use = "cm"
	cmListCmd.Short = "List configMaps"

	EpCmd.AddCommand(&crdListCmd)
	EpCmd.AddCommand(&nodeListCmd)
	EpCmd.AddCommand(&podListCmd)
	EpCmd.AddCommand(&nsListCmd)
	EpCmd.AddCommand(&saListCmd)
	EpCmd.AddCommand(&cmListCmd)
}
