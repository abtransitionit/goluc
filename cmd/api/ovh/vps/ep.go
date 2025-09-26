/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var vpsSDesc = "execute action(s) on OVH VPS via OVH client's request(s)."
var epLDesc = vpsSDesc + `
- This command is used to execute action(s) on VPS using an OVH client's request(s).
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
	EpCmd.AddCommand(getlCmd)
}
