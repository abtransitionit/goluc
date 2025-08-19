/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/gocore/logx"
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
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", provisionSDesc)
	},
}

// func init() {
// 	phase.CmdInit(provisionCmd)
// }

// define the set of phases for this cmd
// Run: phase.CmdRun(kind.ProvisionPhases, provisionSDesc),
