/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gotc

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var gotcSDesc = "Install the GO toolchain to start coding."
var gotcLDesc = gotcSDesc + ` xxx.`
var gotcWkf *phase.Workflow

// root Command
var GotcCmd = &cobra.Command{
	Use:   "gotc",
	Short: gotcSDesc,
	Long:  gotcLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", gotcSDesc)
		// Show the sequence of phases before running the sequence.
		gotcWkf.Show(logx.GetLogger())
		// Run the workflow
		// Run the workflow
		err := gotcWkf.Execute(cmd.Context(), logx.GetLogger())
		if err != nil {
			logx.ErrorWithStack(err, "failed to execute workflow")
		}
	},
}

func init() {
	var err error
	gotcWkf, err = phase.NewWorkflowFromPhases(
		phase.NewPhase("setup", "display the desired KIND Cluster's configuration", internal.CheckSystemStatus, nil),
		phase.NewPhase("build", "check VMs are SSH reachable.", internal.FetchLatestData, nil),
		phase.NewPhase("test", "provision LUC CLI", internal.ProcessData, nil),
		phase.NewPhase("upgrade", "provision OS nodes with latest dnfapt packages and repositories.", internal.GenerateReport, []string{"cpluc"}),
		phase.NewPhase("dapack1", "provision standard/required/missing OS CLI (via dnfapt  packages).", internal.CheckSystemStatus, []string{"upgrade"}),
		phase.NewPhase("dapack2", "provision OS dnfapt package(s) on VM(s).", internal.CheckSystemStatus, []string{"upgrade"}),
		phase.NewPhase("gocli", "provision Go toolchain", internal.GenerateReport, []string{"dapack1"}),
		phase.NewPhase("service", "configure OS services on Kind VMs.", internal.GenerateReport, []string{"dapack1"}),
		phase.NewPhase("linger", "Allow non root user to run OS services.", internal.GenerateReport, []string{"dapack1"}),
		phase.NewPhase("path", "configure OS PATH envvar.", internal.GenerateReport, []string{"dapack1"}),
		phase.NewPhase("rc", "Add a line to non-root user RC file.", internal.GenerateReport, []string{"dapack1"}),
		// phase.NewPhase("rcc", internal.GenerateReport, "darepo: Executes unit and integration tests."),
		// phase.NewPhase("create", internal.GenerateReport, "darepo: Executes unit and integration tests."),
		// phase.NewPhase("check", internal.GenerateReport, "darepo: Executes unit and integration tests."),
	)
	if err != nil {
		logx.ErrorWithStack(err, "failed to build workflow: %v")
	}

	GotcCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
