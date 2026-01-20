/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package desc

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
var epSDesc = "describe resources."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "desc",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	resDescCmd := *res.DescribeCmd
	resDescCmd.Use = "res"
	resDescCmd.Short = "API resources"

	crdDescCmd := *crd.DescribeCmd
	crdDescCmd.Use = "crd"
	crdDescCmd.Short = "describe CRDs"

	nodeDescCmd := *node.DescribeCmd
	nodeDescCmd.Use = "node"
	nodeDescCmd.Short = "describe nodes"

	podDescCmd := *pod.DescribeCmd
	podDescCmd.Use = "pod"
	podDescCmd.Short = "describe pods"

	nsDescCmd := *ns.DescribeCmd
	nsDescCmd.Use = "ns"
	nsDescCmd.Short = "describe namespaces"

	saDescCmd := *sa.DescribeCmd
	saDescCmd.Use = "sa"
	saDescCmd.Short = "describe serviceAccounts"

	cmDescCmd := *cm.DescribeCmd
	cmDescCmd.Use = "cm"
	cmDescCmd.Short = "describe configMaps"

	EpCmd.AddCommand(&crdDescCmd)
	EpCmd.AddCommand(&resDescCmd)
	EpCmd.AddCommand(&nodeDescCmd)
	EpCmd.AddCommand(&podDescCmd)
	EpCmd.AddCommand(&nsDescCmd)
	EpCmd.AddCommand(&saDescCmd)
	EpCmd.AddCommand(&cmDescCmd)
}
