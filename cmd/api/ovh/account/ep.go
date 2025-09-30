/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package account

import (
	"github.com/abtransitionit/goluc/cmd/api/ovh/test"
	"github.com/abtransitionit/goluc/cmd/api/ovh/token"
	"github.com/abtransitionit/goluc/cmd/api/ovh/vps"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "request the OVH API."
var epLDesc = epSDesc + `
- This command is used to run action(s) on OVH resource(s) using an OVH client'srequest.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "ovh",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(vps.EpCmd)
	EpCmd.AddCommand(token.EpCmd)
	EpCmd.AddCommand(test.EpCmd)
}
