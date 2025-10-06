/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package om

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/workflow/om/kbe"
	"github.com/abtransitionit/goluc/cmd/workflow/om/luc"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

var omSDesc = "Access to operations and maintenance (O&M) tools."
var omLDesc = omSDesc + "\n" + `
This command allows you to act on various systems, software or tools, on your local
machine or on remote hosts (ie. VMs or containers).
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "om",
	Short: omSDesc,
	Long:  omLDesc,
	Example: fmt.Sprintf(`
  # manage KIND workflow
  %[1]s workflow kind
	
  # manage go toolchain workflow
  %[1]s workflow gotc

	# manage KBE workflow
  %[1]s workflow kbe
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {
		logx.Infof("%s", omSDesc)
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
	EpCmd.AddCommand(luc.EpCmd)
}
