/*
Copyright © 2025 AB TRANSITION IT
*/
package common

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

// NewWorkflowCmd returns a base cobra.Command configured with logging, config, and default RunE.
func GetEpCmd(cmdName, shortDesc string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   cmdName,
		Short: shortDesc,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logx.GetLogger()

			// Get configuration (package + global + local)
			v, err := viperx.GetConfig("wkf.conf.yaml", "workflow", cmdName)
			if err != nil {
				return err
			}

			// Bind flags and env vars
			viperx.BindFlags(cmd, v, cmdName)

			// Log short description
			logger.Infof("%s", shortDesc)

			// Default action: show help
			cmd.Help()
			return nil
		},
	}

	return cmd
}
