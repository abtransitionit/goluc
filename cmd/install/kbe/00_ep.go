// Copyright © 2025 Amar BELGACEM abtransitionit@hotmail.com
// Fille goluc/cmd/install/kbe/00_ep.go
package kbe

import (
	"os"
	"os/signal"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var kbeSDesc = "provision a Kubernetes clusters."
var kbeLDesc = kbeSDesc + ` xxx.`
var kbeWkf *phase.Workflow
var skipPhases []int
var force bool
var sorted bool
var filtered bool

// root Command
var KbeCmd = &cobra.Command{
	Use:   "kbe",
	Short: kbeSDesc,
	Long:  kbeLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		var logger = logx.GetLogger()
		if filtered {
			// get phases topoSorted
			PhaseSortedByTier, err := kbeWkf.TopoSort(cmd.Context())
			if err != nil {
				logx.ErrorWithStack(err, "failed to sort phases")
				return err
			}
			// filter them
			logx.Info("filtered the tiers")
			PhaseFilteredByTier := PhaseSortedByTier.Filter(*kbeWkf, logx.GetLogger(), skipPhases)

			// show them
			logx.Info("list of filtered phases")
			PhaseFilteredByTier.Show(logger)
			return nil
		}

		if sorted {
			sortedTiers, err := kbeWkf.TopoSort(cmd.Context())
			if err != nil {
				logx.ErrorWithStack(err, "failed to sort phases")
				return err
			}
			sortedTiers.Show(logx.GetLogger())
			// kbeWkf.ShowPhaseList(sortedTiers, logx.GetLogger())
			return nil
		}

		if force {
			// Allows the command to be gracefully canceled by a user (e.g., via Ctrl+C).
			ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
			defer cancel()

			// Pass the skipPhases variable to the Execute method.
			if err := kbeWkf.Execute(ctx, logx.GetLogger(), skipPhases); err != nil {
				logx.ErrorWithStack(err, "failed to execute workflow")
				return err
			}
			return nil
		}

		// Default: just show
		logx.Info("%s", kbeSDesc)
		kbeWkf.Show(logx.GetLogger())
		return nil
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
	)
	if err != nil {
		logx.ErrorWithStack(err, "failed to build workflow: %v")
	}
	KbeCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "phase(s) to skip by ID during execution")
	KbeCmd.Flags().BoolVar(&force, "force", false, "force execution of workflow")
	KbeCmd.Flags().BoolVar(&sorted, "sorted", false, "show phases of a worflow (in topological order)")
	KbeCmd.Flags().BoolVar(&filtered, "filtered", false, "show phases of a workflow (in a topological order and filetered)")
	KbeCmd.AddCommand(provisionCmd)
}
