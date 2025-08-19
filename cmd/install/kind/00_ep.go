/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var kindSDesc = "create a Kind clusters."
var kindLDesc = kindSDesc + ` xxx.`

// var kindSequence = phase.NewPhaseList(
// 	phase.SetPhase("show", internal.SetupFunc, "display the desired KIND Cluster's configuration"),
// 	phase.SetPhase("checkssh", internal.BuildFunc, "check VMs are SSH reachable."),
// 	phase.SetPhase("cpluc", internal.TestFunc, "provision LUC CLI"),
// 	phase.SetPhase("upgrade", internal.TestFunc, "provision OS nodes with latest dnfapt packages and repositories."),
// 	phase.SetPhase("dapack", internal.TestFunc, "provision OS dnfapt package(s) on VM(s)."),
// 	phase.SetPhase("gocli", internal.TestFunc, "provision Go toolchain"),
// 	phase.SetPhase("service", internal.TestFunc, "configure OS services on Kind VMs."),
// 	phase.SetPhase("linger", internal.TestFunc, "Allow non root user to run OS services."),
// 	phase.SetPhase("path", internal.TestFunc, "configure OS PATH envvar."),
// 	phase.SetPhase("rc", internal.TestFunc, "Add a line to non-root user RC file."),
// 	// phase.SetPhase("rcc", internal.TestFunc, "Add lines to non-root user custom RC file."),
// 	// phase.SetPhase("create", internal.TestFunc, "create KIND cluster."),
// 	// phase.SetPhase("check", internal.TestFunc, "check KIND clusters."),
// )

// root Command
var KindCmd = &cobra.Command{
	Use:   "kind",
	Short: kindSDesc,
	Long:  kindLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", kindSDesc)
		// get the logger
		// log := logx.GetLogger()
		// Show the sequence of phases before running the sequence.
		// kindSequence.Show(log)

	},
}

// func init() {
// 	KindCmd.AddCommand(provisionCmd)
// }

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
