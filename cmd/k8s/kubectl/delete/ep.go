/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package delete

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/deploy"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ds"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pv"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sc"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "delete resources."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "delete",
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
	// {cm.DeleteCmd, "cm", "delete configMap"},
	// {crd.DeleteCmd, "crd", "delete CRD"},
	{deploy.DeleteCmd, "deploy", "delete Deployment"},
	{ds.DeleteCmd, "ds", "delete DaemonSet"},
	// {node.DeleteCmd, "node", "delete node"},
	// {ns.DeleteCmd, "ns", "delete namespace"},
	{ns.ResCmd, "ns-r", "delete a namespace's resource"},
	{pod.DeleteCmd, "pod", "delete pod"},
	{pv.DeleteCmd, "pv", "delete PV"},
	{pvc.DeleteCmd, "pvc", "delete PVC"},
	// {res.DeleteCmd, "res", "display API server object type"},
	// {sa.DeleteCmd, "sa", "delete serviceAccount"},
	{sc.DeleteCmd, "sc", "delete SC"},
}

func init() {
	for _, item := range mapCmd {
		cmd := *item.Src
		cmd.Use = item.Use
		cmd.Short = item.Short

		EpCmd.AddCommand(&cmd)
	}
}
