/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package workflow2

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	// om "github.com/abtransitionit/goluc/cmd/workflow2/dep"
	"github.com/abtransitionit/goluc/cmd/workflow2/dep"
	"github.com/abtransitionit/goluc/cmd/workflow2/file"
	"github.com/abtransitionit/goluc/cmd/workflow2/funcx"
	"github.com/abtransitionit/goluc/cmd/workflow2/kbe"
	"github.com/abtransitionit/goluc/cmd/workflow2/kex"
	"github.com/abtransitionit/goluc/cmd/workflow2/kind"
	"github.com/abtransitionit/goluc/cmd/workflow2/kobe"
	"github.com/abtransitionit/goluc/cmd/workflow2/kse"
	"github.com/abtransitionit/goluc/cmd/workflow2/om"
	"github.com/abtransitionit/goluc/cmd/workflow2/os"
	"github.com/abtransitionit/goluc/cmd/workflow2/ovh"
	"github.com/abtransitionit/goluc/cmd/workflow2/ssh"
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
	Use:   "wkf2",
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
	EpCmd.AddCommand(kex.EpCmd)
	EpCmd.AddCommand(kobe.EpCmd)
	EpCmd.AddCommand(kse.EpCmd)
	EpCmd.AddCommand(kind.EpCmd)
	EpCmd.AddCommand(om.EpCmd)
	EpCmd.AddCommand(ovh.EpCmd)
	// building block
	EpCmd.AddCommand(file.EpCmd)
	EpCmd.AddCommand(os.EpCmd)
	EpCmd.AddCommand(ssh.EpCmd)
	EpCmd.AddCommand(dep.EpCmd)
	EpCmd.AddCommand(funcx.EpCmd)
}
