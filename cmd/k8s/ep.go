/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package k8s

import (
	"github.com/abtransitionit/goluc/cmd/k8s/app"
	"github.com/abtransitionit/goluc/cmd/k8s/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_debug"
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "manage k8s resources using kubectl or helm."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "k8s",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(helm.EpCmd)
	EpCmd.AddCommand(kubectl_res.EpCmd)
	EpCmd.AddCommand(kubectl_debug.EpCmd)
	EpCmd.AddCommand(app.EpCmd)
}
