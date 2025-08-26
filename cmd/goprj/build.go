/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package goprj

import (
	"github.com/abtransitionit/gocore/cli"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags.
var (
	projectPath string
	outputDir   string
	platform    string
)

// the command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the artifcat of a Go project",
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logx.GetLogger()
		if err := cli.BuildGoProject(logger, projectPath, outputDir); err != nil {
			logger.Errorf("%v", err)
			return err
		}

		return nil
	},
}

func init() {
	// Add flags to the command.
	buildCmd.Flags().StringVarP(&projectPath, "path", "p", "", "Full path to the Go project to build")
	buildCmd.Flags().StringVarP(&outputDir, "output-dir", "o", "/tmp", "Directory to save the built artifact")
	buildCmd.Flags().StringVarP(&platform, "platform", "", "", "Target platform for cross-compilation (e.g., 'linux/amd64').")

	// Mark the 'path' flag as required.
	buildCmd.MarkFlagRequired("path")
}
