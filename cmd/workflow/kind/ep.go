/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kind

import (
	"github.com/spf13/cobra"
)

// root Command
var EpCmd = &cobra.Command{
	Use:   cmdName,
	Short: sDesc,
	RunE: func(cmd *cobra.Command, args []string) error {

		logger.Infof("%s", sDesc)

		// Default action
		cmd.Help()
		return nil

	},
}

func init() {
	EpCmd.AddCommand(runCmd)
	EpCmd.AddCommand(showCmd)
}
