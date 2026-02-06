/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package desc

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/cm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/deploy"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ep"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/res"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sa"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sc"
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
	{cm.DescribeCmd, "cm", "describe configMap"},
	{deploy.DescribeCmd, "deploy", "describe Deployment"},
	{ep.DescribeCmd, "ep", "describe Entrypoint"},
	{node.DescribeCmd, "node", "describe Node"},
	{pod.DescribeCmd, "pod", "describe Pod"},
	{pvc.DescribeCmd, "pvc", "describe pvc"},
	{ns.DescribeCmd, "ns", "describe Namespace"},
	{sa.DescribeCmd, "sa", "describe ServiceAccount"},
	{sc.DescribeCmd, "sc", "describe StoraceClass"},
}

func init() {
	for _, item := range mapCmd {
		cmd := *item.Src
		cmd.Use = item.Use
		cmd.Short = item.Short

		EpCmd.AddCommand(&cmd)
	}
}
