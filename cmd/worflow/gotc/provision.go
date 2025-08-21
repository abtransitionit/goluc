/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gotc

import (
	"github.com/spf13/cobra"
)

// Description
var provisionSDesc = "deploy a Kind cluster on remote Linux VM(s)."
var provisionLDesc = provisionSDesc + ` xxx.`

// provision Command
var provisionCmd = &cobra.Command{
	Use:   "provision [phase name]",
	Short: provisionSDesc,
	Long:  provisionLDesc,
	// define the set of phases for this cmd
	// Run: phase.CmdRun(kind.ProvisionPhases, provisionSDesc),
}

// func init() {
// 	phase.CmdInit(provisionCmd)
// }
