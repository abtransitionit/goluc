/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package list

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/cm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/deploy"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ds"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ep"
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

type cmd struct {
	Src   *cobra.Command
	Use   string
	Short string
}

var mapCmd = []cmd{
	{cm.ListCmd, "cm", "list configMap"},
	{crd.ListCmd, "crd", "list CRD"},
	{deploy.ListCmd, "deploy", "list Deployment"},
	{ds.ListCmd, "ds", "list DaemonSet"},
	{ep.ListCmd, "ep", "list Enrypoint"},
	{node.ListCmd, "node", "list node"},
	{ns.ListCmd, "ns", "list namespace"},
	{ns.ResCmd, "ns-r", "List a namespace's resource"},
	{pod.ListCmd, "pod", "list pod"},
	{pv.ListCmd, "pv", "list PV"},
	{pvc.ListCmd, "pvc", "list PVC"},
	{res.ListCmd, "res", "display API server object type"},
	{sa.ListCmd, "sa", "list serviceAccount"},
	{sc.ListCmd, "sc", "list SC"},
}

func init() {
	for _, item := range mapCmd {
		cmd := *item.Src
		cmd.Use = item.Use
		cmd.Short = item.Short

		EpCmd.AddCommand(&cmd)
	}
}
