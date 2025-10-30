/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package funcn

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

// root Command
var EpCmd = &cobra.Command{
	Use:   cmdName,
	Short: shortDesc,
	RunE: func(cmd *cobra.Command, args []string) error {

		// define logger
		logger := logx.GetLogger()

		// get config (package+global+local)
		v, err := viperx.GetConfig(cmdName)
		if err != nil {
			return err
		}

		// Bind flags and env vars
		viperx.BindFlags(cmd, v, cmdName)

		// log
		logger.Infof("%s", cmd.Short)

		// Default action
		cmd.Help()
		return nil
	},
}

func init() {
	EpCmd.AddCommand(printcCmd)
	EpCmd.AddCommand(printwCmd)
	EpCmd.AddCommand(runCmd)
}
