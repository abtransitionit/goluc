/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package api

import (
	"github.com/abtransitionit/goluc/cmd/api/ovh"
	"github.com/spf13/cobra"
)

var forceFlag bool

// Description
var epSDesc = "request some API(s)"
var epLDesc = epSDesc + `
- This command is used to run action(s) on specific resource(s) using a specific client'srequest.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "api",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(ovh.EpCmd)
}
