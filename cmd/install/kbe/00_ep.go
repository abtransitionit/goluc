/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var kbeSDesc = "provision a Kubernetes clusters."
var kbeLDesc = kbeSDesc + ` xxx.`

// var kbeSequence = phase.NewPhaseList(
// 	phase.SetPhase("show", internal.SetupFunc, "display the desired KBE Cluster's configuration."),
// 	phase.SetPhase("checklist", internal.BuildFunc, "check VMs are SSH reachable."),
// 	phase.SetPhase("cpluc", internal.TestFunc, "provision LUC CLI"),
// 	phase.SetPhase("upgrade", internal.TestFunc, "provision OS nodes with latest dnfapt packages and repositories."),
// 	phase.SetPhase("dapack1", internal.TestFunc, "provision standard/required/missing OS CLI (via dnfapt  packages)."),
// 	phase.SetPhase("darepo", internal.TestFunc, "provision dnfapt repositories."),

// phase.SetPhase("dapack", internal.TestFunc, "darepo: Executes unit and integration tests."),
// phase.SetPhase("gocli", internal.TestFunc, "darepo: Executes unit and integration tests."),
// phase.SetPhase("update", internal.TestFunc, "darepo: Executes unit and integration tests."),
// phase.SetPhase("reboot", internal.TestFunc, "darepo: Executes unit and integration tests."),
// )

// root Command
var KbeCmd = &cobra.Command{
	Use:   "kbe",
	Short: kbeSDesc,
	Long:  kbeLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", kbeSDesc)
		// get the logger
		// log := logx.GetLogger()
		// Show the sequence of phases before running the sequence.
		// kbeSequence.Show(log)

	},
}

func init() {
	KbeCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
