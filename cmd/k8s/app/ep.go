/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package app

import (
	"github.com/abtransitionit/goluc/cmd/k8s/app/cilium"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "manage resources of specific k8s application."
var epLDesc = epSDesc + `
- This command is used to run action(s) on resources of specific k8s application.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "app",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(cilium.EpCmd)
}
