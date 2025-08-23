/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// Package variables
var retainPhases []int
var skipPhases []int
var force bool
var dryRun bool

// root Command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "execute the workflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Check --skip-phase or --retain-phase are mutually exclusive
		if len(skipPhases) > 0 && len(retainPhases) > 0 {
			logger.Info("flags --skip-phase and --keep-phase cannot be used together")
			return nil
		}

		// The dry-run flag does not need the --force flag, as it is a non-destructive action.
		if dryRun {
			logger.Info("Executing in dry-run mode.")
			// Call the new Plan method from the gocore library
			return wkf.DryRun(cmd.Context(), logger, skipPhases, retainPhases)
		}

		// The --force flag is a security gate; it must be present for any execution.
		if !force {
			logger.Info("The --force flag is required to execute the workflow.")
			return cmd.Help()
		}

		// Check for the dry-run flag first
		if dryRun {
			logger.Info("execute workflow in dry-run mode.")
			// Call the DryRun method
			return wkf.DryRun(cmd.Context(), logger, skipPhases, retainPhases)
		}

		if len(skipPhases) != 0 {
			logger.Infof("execute workflow while skipping phases: %v", skipPhases)
			// define a context: allow usr to cancel the workflow execution with CTRL+C
			ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
			defer cancel()

			// execute the workflow
			if err := wkf.Execute(ctx, logger, skipPhases, retainPhases); err != nil {
				// if err := wkf.Execute(ctx, logger, skipPhases, dryRun); err != nil {
				logger.ErrorWithStack(err, "failed to execute workflow")
				return err
			}
			// success
			return nil

		}

		if len(retainPhases) != 0 {
			logger.Infof("execute workflow restricted to selected phases %v", retainPhases)
			return nil
		}

		// Default action : here we have only flag --force
		// define a context: allow usr to cancel the workflow execution with CTRL+C
		ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
		defer cancel()

		// execute the workflow with the correct parameters (in that case, skipPhases is empty)
		if err := wkf.Execute(ctx, logger, skipPhases, retainPhases); err != nil {
			// if err := wkf.Execute(ctx, logger, skipPhases, dryRun); err != nil {
			logger.ErrorWithStack(err, "failed to execute workflow")
			return err
		}
		// success
		return nil

	},
}

func init() {
	runCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "phase(s) to skip by ID during execution")
	runCmd.Flags().IntSliceVarP(&retainPhases, "retain-phase", "r", []int{}, "phase(s) to retain by ID during execution")
	runCmd.Flags().BoolVar(&force, "force", false, "security flag, needed to execute the workflow")
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "show the execution plan without executing any phases")

}
