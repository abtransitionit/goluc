/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package desc

import (
	res "github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/_res"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/cm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/deploy"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/ep"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/sa"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/sc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/svc"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "display a resource status	"
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
	{svc.DescribeCmd, "svc", "describe Service"},
}

func init() {
	for _, item := range mapCmd {
		cmd := *item.Src
		cmd.Use = item.Use
		cmd.Short = item.Short

		EpCmd.AddCommand(&cmd)
	}
}
