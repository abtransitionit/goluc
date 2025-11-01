/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/abtransitionit/gocore/phase"
)

func GetPrintwCmd(cmdName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "printw",
		Short: "display the worflow",
		RunE: func(cmd *cobra.Command, args []string) error {
			// define logger
			// logger := logx.GetLogger()

			// get the workflow yaml
			workflow, err := phase.GetWorkflow(cmdName)
			if err != nil {
				return fmt.Errorf("getting workflow: %w", err)
			}

			// print
			workflow.Print()
			return nil
		},
	}

	return cmd
}
