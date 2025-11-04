/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package om

import (
	"fmt"
	"path/filepath"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/workflow2/om/luc"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// root Command
var EpCmd = &cobra.Command{
	Use:   filepath.Base(cmdPathName),
	Short: shortDesc,
	Long:  longDesc,
	Example: fmt.Sprintf(`
  # manage KIND workflow
  %[1]s workflow kind
	
  # manage go toolchain workflow
  %[1]s workflow gotc

	# manage KBE workflow
  %[1]s workflow kbe
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {
		logx.Infof("%s", cmd.Short)
		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}

func init() {
	// sub cde
	EpCmd.AddCommand(luc.EpCmd)
}
