/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"os"
	"os/signal"

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
var force bool
var sorted bool
var filtered bool
var show bool

// root Command
var KindCmd = &cobra.Command{
	Use:   "kind",
	Short: kindSDesc,
	Long:  kindLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		var logger = logx.GetLogger()
		if filtered {
			// get phases topoSorted
			PhaseSortedByTier, err := kindWkf.TopoSort(cmd.Context())
			if err != nil {
				logx.ErrorWithStack(err, "failed to sort phases")
				return err
			}

			// show the phases
			logx.Info("list of worflow phases")
			kindWkf.Show(logger)

			// show the sorted phases
			logx.Info("list of worflow phases")
			PhaseSortedByTier.Show(logger)

			// filter them
			logx.Info("filtered the tiers")
			PhaseFilteredByTier := PhaseSortedByTier.Filter(logx.GetLogger(), skipPhases)

			// show them
			logx.Info("list of filtered phases")
			PhaseFilteredByTier.Show(logger)
			return nil

			// // Then filter out the phases to be skipped
			// filteredTiers, err := kindWkf.FilterPhases(sortedTiers, skipPhases)
			// if err != nil {
			// 	logx.ErrorWithStack(err, "failed to filter phases")
			// 	return err
			// }

			// // Show the filtered and sorted list
			// kindWkf.ShowPhaseList(filteredTiers, logx.GetLogger())
			// return nil
		}

		if sorted {
			// get phases sorted by tier
			PhaseSortedByTier, err := kindWkf.TopoSort(cmd.Context())
			if err != nil {
				logx.ErrorWithStack(err, "failed to sort phases")
				return err
			}

			// show them
			logx.Info("list of sorted phases")
			PhaseSortedByTier.Show(logger)
			return nil
		}

		if force {
			// define a context: allow usr to cancel the workflow execution with CTRL+C
			ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
			defer cancel()

			// run the workflow that recieve the context
			if err := kindWkf.Execute(ctx, logger, skipPhases); err != nil {
				logx.ErrorWithStack(err, "failed to execute workflow")
				return err
			}
			return nil
		}

		if show {

			// show the phases of the workflow
			logx.Info("list of worflow phases")
			kindWkf.Show(logger)
			return nil
		}

		// Default action
		logx.Info("%s", kindSDesc) // log info
		kindWkf.Show(logger)       // show the phases
		return nil
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
	)
	if err != nil {
		logx.ErrorWithStack(err, "failed to build workflow: %v")
	}

	KindCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "phase(s) to skip by ID during execution")
	KindCmd.Flags().BoolVar(&force, "force", false, "force execution of workflow")
	KindCmd.Flags().BoolVar(&sorted, "sorted", false, "show phases of a worflow (in topological order)")
	KindCmd.Flags().BoolVar(&filtered, "filtered", false, "show phases of a workflow (in a topological order and filetered)")
	KindCmd.Flags().BoolVar(&show, "show", false, "show phases of a workflow")
	KindCmd.AddCommand(provisionCmd)
}

// // manage argument
// if len(args) == 0 {
// 	cmd.Help()
// 	return
// }
