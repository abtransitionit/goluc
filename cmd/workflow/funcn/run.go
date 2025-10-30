/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package funcn

import (
	"fmt"

	"github.com/abtransitionit/gocore/phase"
	"github.com/spf13/cobra"
)

// root Command
var runCmd = &cobra.Command{
	Use:   "print",
	Short: "display the worflow",
	RunE: func(cmd *cobra.Command, args []string) error {
		// define logger
		// logger := logx.GetLogger()

		// get the workflow yaml
		workflow, err := phase.GetWorkflow()
		if err != nil {
			return fmt.Errorf("getting workflow: %w", err)
		}

		// print it
		workflow.Print()
		return nil

	},
}
