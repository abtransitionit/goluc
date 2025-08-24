/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"os"
	"os/signal"

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

		// Define the context that will be used for all modes. allow user to cancel with ctrl+c
		ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
		defer cancel()

		// mode: dry-run
		if dryRun {
			//logger.Info("Executing workflow in dry-run mode")
			if err := wkf.DryRun(ctx, logger, skipPhases, retainPhases); err != nil {
				logger.ErrorWithNoStack(err, "failed to dry-run workflow")
			}
			return nil
		}

		// // Actual execution
		// if len(skipPhases) > 0 {
		// 	logger.Infof("Executing workflow with skipped phases: %v", skipPhases)
		// } else if len(retainPhases) > 0 {
		// 	logger.Infof("Executing workflow restricted to retained phases: %v", retainPhases)
		// } else {
		// 	logger.Info("Executing workflow with all phases")
		// }

		// mode: with or without skip/retain
		if err := wkf.Execute(ctx, logger, skipPhases, retainPhases); err != nil {
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
	runCmd.Flags().BoolVar(&force, "force", false, "Security flag required to execute the workflow")
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show the execution plan without executing any phases")
}
