/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package workflow

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/workflow/kbe"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

var workflowSDesc = "update a system to be production grade and security proof."
var workflowLDesc = workflowSDesc + "\n" + `
This command allows to add security concerns to system to be production grade and security proof.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "wkf",
	Short: workflowSDesc,
	Long:  workflowLDesc,
	Example: fmt.Sprintf(`
  # manage KIND workflow
  %[1]s workflow kind
	
  # manage go toolchain workflow
  %[1]s workflow gotc

	# manage KBE workflow
  %[1]s workflow kbe
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {
		logx.Infof("%s", workflowSDesc)
		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}

func init() {
	// define the entry point for each workflow
	EpCmd.AddCommand(kbe.EpCmd)
}
