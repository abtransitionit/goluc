/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kubectl_res

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res"

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

	// std resource
	EpCmd.AddCommand(res.EpCmd)
}
