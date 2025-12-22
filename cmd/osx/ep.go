/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package osx

import (
	"github.com/abtransitionit/goluc/cmd/k8s"
	"github.com/abtransitionit/goluc/cmd/osx/gox"
	"github.com/abtransitionit/goluc/cmd/osx/kernel"
	"github.com/abtransitionit/goluc/cmd/osx/pm"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "manage Linux os resources."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "os",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(kernel.EpCmd)
	EpCmd.AddCommand(pm.EpCmd)
	EpCmd.AddCommand(gox.EpCmd)
	EpCmd.AddCommand(k8s.EpCmd)
}
