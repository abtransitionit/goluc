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
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pv"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/res"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sa"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sc"
	"github.com/spf13/cobra"
)

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
	cmListCmd := *cm.ListCmd
	cmListCmd.Use = "cm"
	cmListCmd.Short = "list configMaps"
	crdListCmd := *crd.ListCmd

	crdListCmd.Use = "crd"
	crdListCmd.Short = "list CRDs"

	nodeListCmd := *node.ListCmd
	nodeListCmd.Use = "node"
	nodeListCmd.Short = "list nodes"

	nsListCmd := *ns.ListCmd
	nsListCmd.Use = "ns"
	nsListCmd.Short = "list namespaces"

	podListCmd := *pod.ListCmd
	podListCmd.Use = "pod"
	podListCmd.Short = "list pods"

	pvListCmd := *pv.ListCmd
	pvListCmd.Use = "pv"
	pvListCmd.Short = "list PVs"

	pvcListCmd := *pvc.ListCmd
	pvcListCmd.Use = "pvc"
	pvcListCmd.Short = "list PVCs"

	resListCmd := *res.ListCmd
	resListCmd.Use = "res"
	resListCmd.Short = "display generics API resources"

	saListCmd := *sa.ListCmd
	saListCmd.Use = "sa"
	saListCmd.Short = "list serviceAccounts"

	scListCmd := *sc.ListCmd
	scListCmd.Use = "sc"
	scListCmd.Short = "list SCs"

	EpCmd.AddCommand(&crdListCmd)
	EpCmd.AddCommand(&cmListCmd)
	EpCmd.AddCommand(&nsListCmd)
	EpCmd.AddCommand(&nodeListCmd)
	EpCmd.AddCommand(&podListCmd)
	EpCmd.AddCommand(&pvListCmd)
	EpCmd.AddCommand(&pvcListCmd)
	EpCmd.AddCommand(&resListCmd)
	EpCmd.AddCommand(&saListCmd)
	EpCmd.AddCommand(&scListCmd)
}
