/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package goprj

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags

// root command
var vetCmd = &cobra.Command{
	Use:   "vet",
	Short: "check the syntax of a go project",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logx.GetLogger()
		logger.Infof("%s", cmd.Short)

		return nil
	},
}
