/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "Test some code."
var epLDesc = epSDesc + `
This command is used to run different kinds of tests on your code or environment.
It can be used for quick checks during development, or more complete scenarios
depending on the subcommand invoked.

Typical use cases:
  - Run unit tests for local code
  - Execute integration tests in a container or remote environment
  - Validate configurations before deploying

`

// root Command
var EpCmd = &cobra.Command{
	Use:   "test",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(playCmd)
	EpCmd.AddCommand(commitCmd)
	EpCmd.AddCommand(delgolucCmd)
}
