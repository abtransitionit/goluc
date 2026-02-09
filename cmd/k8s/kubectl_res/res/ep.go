/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package res

import (
	delete "github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/_delete"
	desc "github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/_desc"
	list "github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/_list"
	res "github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/_res"
	yaml "github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/_yaml"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/crd"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/deploy"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/ds"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/ep"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/mnf"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/node"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/ns"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/pod"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/pv"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/pvc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/sa"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/sc"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/secret"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/svc"
	"github.com/spf13/cobra"
)

// var HelmHost = shared.HelmHost

// Description
var epSDesc = "manage k8s resources using kubectl."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "res",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	// EpCmd.PersistentFlags().StringVar(&shared.HelmHost, "helm-host", "", "Helm host to use")

	EpCmd.AddCommand(list.EpCmd)
	EpCmd.AddCommand(desc.EpCmd)
	EpCmd.AddCommand(yaml.EpCmd)
	EpCmd.AddCommand(res.EpCmd)
	EpCmd.AddCommand(delete.EpCmd)

	// std resource
	EpCmd.AddCommand(crd.EpCmd)
	EpCmd.AddCommand(deploy.EpCmd)
	EpCmd.AddCommand(ds.EpCmd)
	EpCmd.AddCommand(ep.EpCmd)
	EpCmd.AddCommand(node.EpCmd)
	EpCmd.AddCommand(mnf.EpCmd)
	EpCmd.AddCommand(ns.EpCmd)
	EpCmd.AddCommand(pod.EpCmd)
	EpCmd.AddCommand(pv.EpCmd)
	EpCmd.AddCommand(pvc.EpCmd)
	EpCmd.AddCommand(sa.EpCmd)
	EpCmd.AddCommand(sc.EpCmd)
	EpCmd.AddCommand(svc.EpCmd)
	EpCmd.AddCommand(secret.EpCmd)
}
