/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cmd

import (
	"os"

	"github.com/abtransitionit/goluc/cmd/goprj"
	"github.com/abtransitionit/goluc/cmd/property"
	"github.com/abtransitionit/goluc/cmd/test"
	"github.com/abtransitionit/goluc/cmd/workflow"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var rootSDesc = "LUC (aka. Linux Unified CLI) is a user-friendly, auto-documented command-line interface."
var rootLDesc = rootSDesc + ` It simplifies daily tasks for DevOps engineers and developers by providing a unified and consistent CLI experience. LUC can, for example:
	→ Manage containers and container images,
	→ Manage Linux OS packages and repositories using a unified interface — no need to worry about whether it's apt or dnf nor if it's debian, fedora or ubuntu
	→ Manage remote VM objects,
	→ Simplify the creation and management of Kubernetes clusters across virtual machines,
	→ ...and much more.

As a Linux cross-distribution CLI, LUC is also well-suited and ready for full automation and integration into any CI/CD pipelines.
`

// root Command
var rootCmd = &cobra.Command{
	Use:   internal.CliName,
	Short: rootSDesc,
	Long:  rootLDesc,
}

// called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(goprj.GoPrjCmd)
	rootCmd.AddCommand(workflow.WorkflowCmd)
	rootCmd.AddCommand(property.PropertyCmd)
	rootCmd.AddCommand(test.TestCmd)
	// rootCmd.AddCommand(prop.PropCmd)
}
