/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"github.com/abtransitionit/gotask/workflow"
	"github.com/spf13/cobra"
)

// Package variables
var phase bool
var tier bool
var filter bool

// root Command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the workflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		if phase {
			logger.Info("show all phases of a workflow")
			err := workflow.ShowPhase(wkf, logger)
			if err != nil {
				logger.ErrorWithStack(err, "failed to get workflow")
				return err
			}
			return nil
		}

		if tier {
			logger.Info("show all tiers of a workflow (in topological order)")
			_, err := workflow.ShowTier(wkf, logger)
			if err != nil {
				logger.ErrorWithStack(err, "failed to get workflow")
				return err
			}
			return nil

		}

		if filter {
			logger.Info("Show all tiers of a workflow, including skipped or retained phases, in topological order and filtered.")
			return nil

		}

		// Default action
		cmd.Help()
		return nil

	},
}

func init() {
	showCmd.Flags().BoolVar(&phase, "phase", false, "show all phases of a worflow")
	showCmd.Flags().BoolVar(&tier, "tier", false, "show all tiers of a worflow (in topological order)")
	showCmd.Flags().BoolVar(&filter, "filter", false, "Show all tiers of a workflow, including skipped or retained phases, in topological order and filtered.")
}
