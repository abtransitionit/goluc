/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var vpsSDesc = "Execute actions on an OVH VPS via the OVH API."
var epLDesc = vpsSDesc + `
An action represents an operation that can be performed on a VPS through the OVH API.

Examples of actions include:
  • starting or stopping a VPS
  • retrieving information such as status or configuration
  • rebooting, reinstalling, or managing snapshots

This command sends requests to the OVH API client to run these actions on the target VPS.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "vps",
	Short: vpsSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(installCmd)
	EpCmd.AddCommand(listCmd)
	EpCmd.AddCommand(getCmd)
	EpCmd.AddCommand(getliCmd)
	EpCmd.AddCommand(getlvCmd)
}
