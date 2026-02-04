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
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/res"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sa"
	"github.com/spf13/cobra"
)

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

type Cmd struct {
	Src   *cobra.Command
	Use   string
	Short string
}

var mapCmd = []Cmd{
	{res.DescribeCmd, "res", "API resource"},
	{crd.DescribeCmd, "crd", "describe CRD"},
	{node.DescribeCmd, "node", "describe node"},
	{pod.DescribeCmd, "pod", "describe pod"},
	{pvc.DescribeCmd, "pvc", "describe pvc"},
	{ns.DescribeCmd, "ns", "describe namespace"},
	{sa.DescribeCmd, "sa", "describe serviceAccount"},
	{cm.DescribeCmd, "cm", "describe configMap"},
}

func init() {
	for _, item := range mapCmd {
		cmd := *item.Src
		cmd.Use = item.Use
		cmd.Short = item.Short

		EpCmd.AddCommand(&cmd)
	}
	// resDescCmd := *res.DescribeCmd
	// resDescCmd.Use = "res"
	// resDescCmd.Short = "API resource"

	// crdDescCmd := *crd.DescribeCmd
	// crdDescCmd.Use = "crd"
	// crdDescCmd.Short = "describe CRD"

	// nodeDescCmd := *node.DescribeCmd
	// nodeDescCmd.Use = "node"
	// nodeDescCmd.Short = "describe node"

	// podDescCmd := *pod.DescribeCmd
	// podDescCmd.Use = "pod"
	// podDescCmd.Short = "describe pod"

	// pvcDescCmd := *pvc.DescribeCmd
	// pvcDescCmd.Use = "pvc"
	// pvcDescCmd.Short = "describe pvc"

	// nsDescCmd := *ns.DescribeCmd
	// nsDescCmd.Use = "ns"
	// nsDescCmd.Short = "describe namespace"

	// saDescCmd := *sa.DescribeCmd
	// saDescCmd.Use = "sa"
	// saDescCmd.Short = "describe serviceAccount"

	// cmDescCmd := *cm.DescribeCmd
	// cmDescCmd.Use = "cm"
	// cmDescCmd.Short = "describe configMap"

	// EpCmd.AddCommand(&crdDescCmd)
	// EpCmd.AddCommand(&resDescCmd)
	// EpCmd.AddCommand(&nodeDescCmd)
	// EpCmd.AddCommand(&podDescCmd)
	// EpCmd.AddCommand(&pvcDescCmd)
	// EpCmd.AddCommand(&nsDescCmd)
	// EpCmd.AddCommand(&saDescCmd)
	// EpCmd.AddCommand(&cmDescCmd)
}
