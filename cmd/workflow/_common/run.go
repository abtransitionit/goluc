/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gocore/viperx"
)

func GetRunCmd(cmdName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "execute the workflow",
		RunE: func(cmd *cobra.Command, args []string) error {
			// define logger
			logger := logx.GetLogger()

			// get the config
			config, err := viperx.GetConfig("wkf.conf.yaml", "workflow", cmdName)
			if err != nil {
				return err
			}

			// get the workflow yaml
			workflow, err := phase2.GetWorkflow(cmdName)
			if err != nil {
				return fmt.Errorf("getting workflow: %w", err)
			}

			// execute the workflow
			err = workflow.Execute(config, logger)
			if err != nil {
				return fmt.Errorf("executing workflow: %w", err)
			}

			// success
			return nil
		},
	}

	return cmd
}
