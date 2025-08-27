/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package goprj

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/cli"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags.
var (
	localArtifactPath string
	VmName            string
	remoteFilePath    string
	force             bool
	dryRun            bool
)

var deploySDesc = "deploy an artifcat on a remote host"
var deployLDesc = deploySDesc + "\n" + `
Example usage.

go run . goprj deploy --force -l /tmp/goluc-linux -r /usr/local/bin/goluc -v o1u
clear && go  run . goprj build -o /tmp -p /Users/max/wkspc/git/goluc/ && go run . goprj deploy --force -l /tmp/goluc-linux -r/usr/local/bin/goluc -vo1u
`

// the command
var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: deploySDesc,
	Long:  deployLDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logx.GetLogger()

		// check requirement: either --dry-run OR --force must be used
		if !force && !dryRun {
			logger.Info("The --force flag is required to execute the workflow")
			return cmd.Help()
		}
		// check requirement: either --dry-run OR --force must be used

		// here --fore or dry-run is used
		// convert cli inputs to function inputs
		hostFilePath := fmt.Sprintf("%s:%s", VmName, remoteFilePath)

		logger.Infof("Deploying %s to %s", localArtifactPath, hostFilePath)
		// deploy the artifact : ie. scp fie to remote
		deployOk, err := cli.DeployGoArtifactAsSudo(logger, localArtifactPath, hostFilePath)
		if err != nil {
			logger.Errorf("%v", err)
			return err
		}

		if !deployOk {
			logger.Error("Deployment failed")
			return fmt.Errorf("deployment failed")
		}

		// success
		logger.Info("Deployment successful")
		return nil
	},
}

func init() {
	// Add flags to the command.
	deployCmd.Flags().StringVarP(&localArtifactPath, "lpath", "l", "", "Full path to the local artifact to deploy")
	deployCmd.Flags().StringVarP(&VmName, "Vm", "v", "", "Name of the VM")
	deployCmd.Flags().StringVarP(&remoteFilePath, "rpath", "r", "", "Full path of the remote file to create on the VM")
	deployCmd.Flags().BoolVar(&force, "force", false, "Security flag required to execute the deployment")
	deployCmd.Flags().BoolVarP(&dryRun, "dry-run", "d", false, "Show the execution plan without dploying the artifact")

	// Mark the 'path' flag as required.
	deployCmd.MarkFlagRequired("lpath")
	deployCmd.MarkFlagRequired("Vm")
	deployCmd.MarkFlagRequired("rpath")
}
