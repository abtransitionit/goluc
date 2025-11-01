/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

func GetPrintcCmd(cmdName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "printc",
		Short: "display the worflow's config",
		RunE: func(cmd *cobra.Command, args []string) error {
			// define logger
			// logger := logx.GetLogger()

			// get the config yaml
			config, err := viperx.GetConfig(cmdName)
			if err != nil {
				return err
			}

			// print config
			config.Print()
			return nil
		},
	}

	return cmd
}
