/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/phase2"
)

func GetPrinttCmd(cmdName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "printt",
		Short: "display the worflow tiers and phases",
		RunE: func(cmd *cobra.Command, args []string) error {
			// define logger
			// logger := logx.GetLogger()

			// get the workflow yaml
			workflow, err := phase2.GetWorkflow(cmdName)
			if err != nil {
				return fmt.Errorf("getting workflow: %w", err)
			}

			// get table
			table, err := workflow.GetTableTier()
			if err != nil {
				return fmt.Errorf("getting table: %w", err)
			}

			// print
			list.PrettyPrintTable(table)

			// success
			return nil
		},
	}

	return cmd
}
