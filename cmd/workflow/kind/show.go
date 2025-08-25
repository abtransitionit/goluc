// File goluc/cmd/workflow/kind/show.go
/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"context"
	"os"
	"os/signal"

	corectx "github.com/abtransitionit/gocore/ctx"
	"github.com/abtransitionit/gotask/workflow"
	"github.com/spf13/cobra"
)

// Package variables
var phase bool
var tier bool
var filter bool
var testctx bool

// root Command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the workflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		if testctx {
			logger.Info("testing passing a var to the context and use it by a phase in a library")
			// create a context
			ctx := context.Background()
			// allow user to ctrl-c via ctx
			ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
			defer cancel()

			// add the var to the ctx
			ctx = context.WithValue(ctx, corectx.WorkflowKeyId, wkf)

			// call any phase (they mandatory have all the same signature)
			result, err := workflow.ShowWorkflow(ctx, logger)
			if err != nil {
				logger.ErrorWithStack(err, "failed to get workflow")
				return err
			}

			// show the result
			logger.Infof("result of the call: %s", result)
			logger.Infof("var pass to the context: %s", result)
			logger.Info("the test is to pass the worflow var and make it use by a phase in a library")
			return nil

		}
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
	showCmd.Flags().BoolVar(&testctx, "tctx", false, "test passing a vat to the context.")
}
