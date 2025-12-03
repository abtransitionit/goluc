/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package workflow

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"

	// "github.com/abtransitionit/goluc/cmd/workflow/gotc"
	"github.com/abtransitionit/goluc/cmd/workflow/kbe"
	"github.com/abtransitionit/goluc/cmd/workflow/kind"
	"github.com/abtransitionit/goluc/cmd/workflow/om"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

var workflowSDesc = "Manage systems, software or tools using a workflow."
var workflowLDesc = workflowSDesc + "\n" + `
This command allows you to act on various systems, software or tools, on your local
machine or on remote hosts (ie. VMs or containers).
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
	// EpCmd.AddCommand(gotc.EpCmd)
	EpCmd.AddCommand(kbe.EpCmd)
	EpCmd.AddCommand(kind.EpCmd)
	EpCmd.AddCommand(om.EpCmd)
}
