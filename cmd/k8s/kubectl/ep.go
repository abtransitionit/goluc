/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kubectl

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/cm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/desc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/list"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/mnf"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/res"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sa"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/sc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl/secret"

	// "github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage k8s resources using kubectl."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "kubectl",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	// EpCmd.PersistentFlags().StringVar(&shared.HelmHost, "helm-host", "", "Helm host to use")
	EpCmd.AddCommand(list.EpCmd)
	EpCmd.AddCommand(desc.EpCmd)

	EpCmd.AddCommand(cm.EpCmd)
	EpCmd.AddCommand(crd.EpCmd)
	EpCmd.AddCommand(node.EpCmd)
	EpCmd.AddCommand(mnf.EpCmd)
	EpCmd.AddCommand(ns.EpCmd)
	EpCmd.AddCommand(pod.EpCmd)
	EpCmd.AddCommand(res.EpCmd)
	EpCmd.AddCommand(sa.EpCmd)
	EpCmd.AddCommand(sc.EpCmd)
	EpCmd.AddCommand(secret.EpCmd)
}
