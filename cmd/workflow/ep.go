/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package workflow

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/workflow/gotc"
	"github.com/abtransitionit/goluc/cmd/workflow/kbe"
	"github.com/abtransitionit/goluc/cmd/workflow/kind"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description

var workflowSDesc = "Manage the provisioning process of systems, software and tools."
var workflowLDesc = workflowSDesc + "\n" + `
This command allows you to install various systems, software or tools, on your local
machine or on one or more remote hosts, VMs or containers.

The installation process is driven by a workflow and a custom configuration YAML file.
The structure of this YAML file depends on the system, software or tool being installed,
and may define:
  - Target hosts, VMs or containers
  - Installation parameters and options
  - Environment-specific settings
`

// root Command
var WorkflowCmd = &cobra.Command{
	Use:   "workflow",
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
	WorkflowCmd.AddCommand(gotc.EpCmd)
	WorkflowCmd.AddCommand(kbe.EpCmd)
	WorkflowCmd.AddCommand(kind.EpCmd)
}
