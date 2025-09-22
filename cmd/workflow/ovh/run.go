/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"context"
	"os"
	"os/signal"

	corectx "github.com/abtransitionit/gocore/ctx"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags
var retainPhases []int
var skipPhases []int
var force bool
var dryRun bool

// root command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the workflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		// check requirement: --skip-phase and --retain-phase are mutually exclusive
		if len(skipPhases) > 0 && len(retainPhases) > 0 {
			logger.Info("flags --skip-phase and --retain-phase cannot be used together")
			return nil
		}

		// check requirement: either --dry-run OR --force must be used
		if !force && !dryRun {
			logger.Info("The --force flag is required to execute the workflow")
			return cmd.Help()
		}

		// Define the context that will be used by all phases of tyhe workflow
		PhaseCtx := context.Background()
		// allowuser to cancel with ctrl+c
		PhaseCtx, cancel := signal.NotifyContext(PhaseCtx, os.Interrupt)
		defer cancel()
		// add var 01
		PhaseCtx = context.WithValue(PhaseCtx, corectx.StringKeyId, "exec-123")
		PhaseCtx = context.WithValue(PhaseCtx, corectx.WorkflowKeyId, wkf)

		// - add var to the context

		// mode: dry-run
		if dryRun {
			if err := wkf.DryRun(PhaseCtx, logger, skipPhases, retainPhases); err != nil {
				logger.ErrorWithNoStack(err, "failed to dry-run workflow")
			}
			return nil
		}

		// mode: with or without skip/retain
		if err := wkf.Execute(PhaseCtx, logger, targets, skipPhases, retainPhases); err != nil {
			logger.ErrorWithNoStack(err, "failed to execute workflow")
			return nil
		}

		logger.Info("Workflow executed successfully")
		return nil
	},
}

func init() {
	// var err error
	// wkf, err = corephase.NewWorkflowFromPhases()
	// if err != nil {
	// 	logger.ErrorWithNoStack(err, "failed to build workflow")
	// }

	runCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "Phase(s) to skip by ID during execution")
	runCmd.Flags().IntSliceVarP(&retainPhases, "retain-phase", "r", []int{}, "Phase(s) to retain by ID during execution")
	runCmd.Flags().BoolVarP(&force, "force", "f", false, "Security flag required to execute the workflow")
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show the execution plan without executing any phases")
}
