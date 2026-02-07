/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package yaml

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/cm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/deploy"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ep"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/svc"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "display the YAML of a resource."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "yaml",
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
	// {res.YamlCmd, "res", "API resource"},
	{crd.YamlCmd, "crd", "yaml CRD"},
	{cm.YamlCmd, "cm", "yaml configMap"},
	{deploy.YamlCmd, "deploy", "yaml Deployment"},
	{ep.YamlCmd, "deploy", "yaml Entrypoin"},
	{node.YamlCmd, "node", "yaml node"},
	{ns.YamlCmd, "ns", "yaml namespace"},
	{pod.YamlCmd, "pod", "yaml pod"},
	{pvc.YamlCmd, "pvc", "yaml pvc"},
	{svc.YamlCmd, "pvc", "yaml service"},
	// {sa.YamlCmd, "sa", "yaml serviceAccount"},
}

func init() {
	for _, item := range mapCmd {
		cmd := *item.Src
		cmd.Use = item.Use
		cmd.Short = item.Short

		EpCmd.AddCommand(&cmd)
	}

}
