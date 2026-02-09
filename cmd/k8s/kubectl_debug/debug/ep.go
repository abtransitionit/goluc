/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package debug

import (
	ishell "github.com/abtransitionit/goluc/cmd/k8s/kubectl_debug/debug/isshel"
	"github.com/spf13/cobra"
)

// var HelmHost = shared.HelmHost

// Description
var epSDesc = "tool to debug a k8s cluster"
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "debug",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(ishell.EpCmd)
}
