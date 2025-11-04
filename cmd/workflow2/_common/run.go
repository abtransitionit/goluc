/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

func GetRunCmd(cmdPathName string) *cobra.Command {

	cobraCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the workflow",
		RunE: func(cobraCmd *cobra.Command, args []string) error {

			// define logger and ctx
			logger := logx.GetLogger()
			ctx := context.Background()

			// get the workflow configuration file into a struct
			config, err := viperx.GetViperx("wkf.conf.yaml", "workflow", cmdPathName, logger)
			if err != nil {
				return fmt.Errorf("getting workflow: %w", err)
			}

			// get the workflow itself into a struct
			workflow, err := phase2.GetWorkflow("wkf.phase.yaml", cmdPathName, logger)
			if err != nil {
				return fmt.Errorf("getting workflow: %w", err)
			}

			// execute the workflow
			err = workflow.Execute(ctx, config, FunctionRegistry, logger)
			if err != nil {
				return fmt.Errorf("executing workflow: %w", err)
			}

			// success
			return nil
		},
	}

	return cobraCmd
}
