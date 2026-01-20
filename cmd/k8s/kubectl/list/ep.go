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
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/res"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sa"
	"github.com/spf13/cobra"
)

// var
var localFlag bool

// Description
var epSDesc = "list resources."
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
	resListCmd := *res.ListCmd
	resListCmd.Use = "res"
	resListCmd.Short = "display generics API resources"

	crdListCmd := *crd.ListCmd
	crdListCmd.Use = "crd"
	crdListCmd.Short = "list CRDs"

	nodeListCmd := *node.ListCmd
	nodeListCmd.Use = "node"
	nodeListCmd.Short = "list nodes"

	podListCmd := *pod.ListCmd
	podListCmd.Use = "pod"
	podListCmd.Short = "list pods"

	nsListCmd := *ns.ListCmd
	nsListCmd.Use = "ns"
	nsListCmd.Short = "list namespaces"

	saListCmd := *sa.ListCmd
	saListCmd.Use = "sa"
	saListCmd.Short = "list serviceAccounts"

	cmListCmd := *cm.ListCmd
	cmListCmd.Use = "cm"
	cmListCmd.Short = "list configMaps"

	EpCmd.AddCommand(&crdListCmd)
	EpCmd.AddCommand(&resListCmd)
	EpCmd.AddCommand(&nodeListCmd)
	EpCmd.AddCommand(&podListCmd)
	EpCmd.AddCommand(&nsListCmd)
	EpCmd.AddCommand(&saListCmd)
	EpCmd.AddCommand(&cmListCmd)
}
