/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gotc

import (
	"os"
	"os/signal"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var gotcSDesc = "Install the GO toolchain to start coding."
var gotcLDesc = gotcSDesc + ` xxx.`
var gotcWkf *phase.Workflow
var skipPhases []int
var force bool
var sorted bool
var filtered bool

// root Command
var GotcCmd = &cobra.Command{
	Use:   "gotc",
	Short: gotcSDesc,
	Long:  gotcLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		if filtered {
			// First, get sorted phases
			sortedTiers, err := gotcWkf.SortedPhases(cmd.Context())
			if err != nil {
				logx.ErrorWithStack(err, "failed to sort phases")
				return err
			}

			// Then filter out the phases to be skipped
			filteredTiers, err := gotcWkf.FilterPhases(sortedTiers, skipPhases)
			if err != nil {
				logx.ErrorWithStack(err, "failed to filter phases")
				return err
			}

			// Show the filtered and sorted list
			gotcWkf.ShowPhaseList(filteredTiers, logx.GetLogger())
			return nil
		}

		if sorted {
			sortedTiers, err := gotcWkf.SortedPhases(cmd.Context())
			if err != nil {
				logx.ErrorWithStack(err, "failed to sort phases")
				return err
			}
			gotcWkf.ShowPhaseList(sortedTiers, logx.GetLogger())
			return nil
		}

		if force {
			// Run the workflow with this context
			ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
			defer cancel()

			if err := gotcWkf.Execute(ctx, logx.GetLogger(), skipPhases); err != nil {
				logx.ErrorWithStack(err, "failed to execute workflow")
				return err
			}
			return nil
		}

		// Default: just show
		logx.Info("%s", gotcSDesc)
		gotcWkf.Show(logx.GetLogger())
		return nil
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

	GotcCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "phase(s) to skip by ID during execution")
	GotcCmd.Flags().BoolVar(&force, "force", false, "force execution of workflow")
	GotcCmd.Flags().BoolVar(&sorted, "sorted", false, "show phases in topological order")
	GotcCmd.Flags().BoolVar(&filtered, "filtered", false, "show phases in topological order and filetered")
	GotcCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
