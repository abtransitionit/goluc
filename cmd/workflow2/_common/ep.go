/*
Copyright © 2025 AB TRANSITION IT
*/
package common

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/spf13/cobra"
)

var FunctionRegistry = phase2.GetFnRegistry()

// Description: returns a cobra.Command
func GetEpCmd(cmdPathName, shortDesc string) *cobra.Command {
	cmdName := filepath.Base(cmdPathName)
	cobraCmd := &cobra.Command{
		Use:   cmdName,
		Short: shortDesc,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Default action: show help
			cmd.Help()
			return nil
		},
	}

	return cobraCmd
}
