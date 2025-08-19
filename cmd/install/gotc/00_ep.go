/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gotc

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var gotcSDesc = "Install the GO toolchain to start coding."
var gotcLDesc = gotcSDesc + ` xxx.`

// var gotcSequence = phase.NewPhaseList(
// 	phase.SetPhase("Setup", internal.SetupFunc, "Prepares the environment for the build."),
// 	phase.SetPhase("Build", internal.BuildFunc, "Compiles the source code into a binary."),
// 	phase.SetPhase("Test", internal.TestFunc, "Executes unit and integration tests."),
// )

// root Command
var GotcCmd = &cobra.Command{
	Use:   "gotc",
	Short: gotcSDesc,
	Long:  gotcLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", gotcSDesc)
		// get the logger
		// log := logx.GetLogger()
		// Show the sequence of phases before running the sequence.
		// gotcSequence.Show(log)
	},
}

func init() {
	GotcCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
