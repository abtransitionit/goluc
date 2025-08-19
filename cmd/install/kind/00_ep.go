/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var kindSDesc = "create a Kind clusters."
var kindLDesc = kindSDesc + ` xxx.`
var kindWkf *phase.Workflow
var skipPhases []int

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
		// Show the sequence of phases before running the sequence.
		kindWkf.Show(logx.GetLogger())

	},
}

func init() {
	var err error
	kindWkf, err = phase.NewWorkflowFromPhases(
		phase.NewPhase("show", "display the desired KIND Cluster's configuration", internal.CheckSystemStatus, nil),
		phase.NewPhase("checklist", "check VMs are SSH reachable.", internal.FetchLatestData, nil),
		phase.NewPhase("cpluc", "provision LUC CLI", internal.ProcessData, nil),
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

	KindCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "phase(s) to skip by ID during execution")
	KindCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
