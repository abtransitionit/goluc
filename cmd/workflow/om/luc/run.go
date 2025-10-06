/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package luc

import (
	"context"
	"os"
	"os/signal"

	corectx "github.com/abtransitionit/gocore/ctx"
	"github.com/abtransitionit/gocore/list"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags
var retainPhaseRange []string
var skipPhaseRange []string
var force bool
var dryRun bool

// root command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Execute the workflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		var err error
		var retainPhaseSlice, skipPhaseSlice []int
		// convert retainPhaseRange into  retailPhaseSlice (ie. []int)
		if len(retainPhaseRange) > 0 {
			retainPhaseSlice, err = list.ConvertRangeToSliceInt(retainPhaseRange)
			if err != nil {
				logger.ErrorWithNoStack(err, "invalid --retain-phase range")
				return err
			}
		}
		// convert skipPhaseRange into  skipPhaseSlice (ie. []int)
		if len(skipPhaseRange) > 0 {
			skipPhaseSlice, err = list.ConvertRangeToSliceInt(skipPhaseRange)
			if err != nil {
				logger.ErrorWithNoStack(err, "invalid --skip-phase range")
				return err
			}
		}

		// check requirement: --skip-phase and --retain-phase are mutually exclusive
		if len(skipPhaseSlice) > 0 && len(retainPhaseSlice) > 0 {
			logger.Info("flags --skip-phase and --retain-phase cannot be used together")
			return nil
		}

		// check requirement: either --dry-run OR --force must be used
		if !force && !dryRun {
			logger.Info("The --force flag is required to execute the workflow")
			return cmd.Help()
		}

		// Define the context that will be used by all phases of tyhe workflow
		ctx := context.Background()
		// allowuser to cancel with ctrl+c
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()
		// add var 01
		ctx = context.WithValue(ctx, corectx.StringKeyId, "exec-123")
		ctx = context.WithValue(ctx, corectx.WorkflowKeyId, wkf)

		// mode: dry-run
		if dryRun {
			if err := wkf.DryRun(ctx, logger, skipPhaseSlice, retainPhaseSlice); err != nil {
				logger.ErrorWithNoStack(err, "failed to dry-run workflow")
			}
			return nil
		}

		// mode: with or without skip/retain
		if err := wkf.Execute(ctx, logger, targets, skipPhaseSlice, retainPhaseSlice); err != nil {
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

	runCmd.Flags().StringSliceVarP(&skipPhaseRange, "skip-phase", "s", []string{}, "Phase(s) to skip by ID during execution")
	runCmd.Flags().StringSliceVarP(&retainPhaseRange, "retain-phase", "r", []string{}, "Phase(s) to retain by ID during execution")
	runCmd.Flags().BoolVarP(&force, "force", "f", false, "Security flag required to execute the workflow")
	runCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show the execution plan without executing any phases")
}

// skipPhases or retainPhases is a string that will be converted to an int slice by The Execute functions
// possible value "2-3", "2-", "-5", "1,2,3"
