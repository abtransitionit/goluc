/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/spf13/cobra"
)

// Description
var testSDesc = "Test some code."
var testLDesc = testSDesc + `
This command is used to run different kinds of tests on your code or environment.
It can be used for quick checks during development, or more complete scenarios
depending on the subcommand invoked.

Typical use cases:
  - Run unit tests for local code
  - Execute integration tests in a container or remote environment
  - Validate configurations before deploying

`

// root Command
var TestCmd = &cobra.Command{
	Use:   "test",
	Short: testSDesc,
	Long:  testLDesc,
}

func init() {
	TestCmd.AddCommand(playCmd)
}
