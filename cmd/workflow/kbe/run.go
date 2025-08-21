/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"os"
	"os/signal"

	"github.com/spf13/cobra"
)

// Package variables
var keepPhases []int
var skipPhases []int
var force bool

// root Command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "execute the workflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		if len(skipPhases) != 0 {
			logger.Infof("execute workflow while skipping phases: %v", skipPhases)
			return nil
		}

		if len(keepPhases) != 0 {
			logger.Infof("execute workflow restricted to selected phases %v", keepPhases)
			return nil
		}

		if force {
			logger.Info("execute workflow with all phases (force enabled)")
			// define a context: allow usr to cancel the workflow execution with CTRL+C
			ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt)
			defer cancel()

			// execute the workflow
			if err := wkf.Execute(ctx, logger, skipPhases); err != nil {
				logger.ErrorWithStack(err, "failed to execute workflow")
				return err
			}

			return nil
		}

		// Default action
		cmd.Help()
		return nil

	},
}

func init() {
	runCmd.Flags().IntSliceVarP(&skipPhases, "skip-phase", "s", []int{}, "phase(s) to skip by ID during execution")
	runCmd.Flags().IntSliceVarP(&keepPhases, "keep-phase", "k", []int{}, "phase(s) to keep by ID during execution")
	runCmd.Flags().BoolVar(&force, "force", false, "security flag, needed to execute the workflow")
}
