/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kubectl_debug

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_debug/ishell"
	"github.com/spf13/cobra"
)

// var HelmHost = shared.HelmHost

// Description
var epSDesc = "debug a k8s cluster with kubectl"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "debug",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	// EpCmd.PersistentFlags().StringVar(&shared.HelmHost, "helm-host", "", "Helm host to use")

	// std resource
	EpCmd.AddCommand(ishell.EpCmd)
}
