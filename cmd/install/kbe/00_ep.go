/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var kbeSDesc = "provision a Kubernetes clusters."
var kbeLDesc = kbeSDesc + ` xxx.`
var kbeWkf *phase.Workflow

// root Command
var KbeCmd = &cobra.Command{
	Use:   "kbe",
	Short: kbeSDesc,
	Long:  kbeLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", kbeSDesc)
		// Show the sequence of phases before running the sequence.
		kbeWkf.Show(logx.GetLogger())
		// Run the workflow
		err := kbeWkf.Execute(cmd.Context(), logx.GetLogger())
		if err != nil {
			logx.ErrorWithStack(err, "failed to execute workflow")
		}

	},
}

func init() {
	var err error
	kbeWkf, err = phase.NewWorkflowFromPhases(
		phase.NewPhase("show", "display the desired KBE Cluster's configuration.", internal.CheckSystemStatus, nil),
		phase.NewPhase("checklist", "check VMs are SSH reachable.", internal.FetchLatestData, nil),
		phase.NewPhase("cpluc", "provision LUC CLI", internal.ProcessData, nil),
		phase.NewPhase("upgrade", "provision OS nodes with latest dnfapt packages and repositories.", internal.GenerateReport, []string{"cpluc"}),
		phase.NewPhase("dapack1", "provision standard/required/missing OS CLI (via dnfapt  packages).", internal.CheckSystemStatus, []string{"upgrade"}),
		phase.NewPhase("dapack2", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"dapack1"}),
		phase.NewPhase("darepo", "provision dnfapt repositories.", internal.GenerateReport, []string{"dapack1"}),
		// phase.NewPhase("dapack", internal.GenerateReport, "darepo: Executes unit and integration tests."),
		// phase.NewPhase("gocli", internal.GenerateReport, "darepo: Executes unit and integration tests."),
		// phase.NewPhase("update", internal.GenerateReport, "darepo: Executes unit and integration tests."),
		// phase.NewPhase("reboot", internal.GenerateReport, "darepo: Executes unit and integration tests."),
	)
	if err != nil {
		logx.ErrorWithStack(err, "failed to build workflow: %v")
	}

	KbeCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
