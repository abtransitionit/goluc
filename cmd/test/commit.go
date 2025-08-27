/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
	"github.com/spf13/cobra"
)

// Description
var commitSDesc = "for all 4 go project: merge dev into main and push main."
var commitLDesc = commitSDesc + ` xxx.`

// root Command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: commitSDesc,
	Long:  commitLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logx.GetLogger()
		logger.Info(commitSDesc)
		command := "git checkout main ; git merge --no-edit dev; git push origin main; git checkout dev"
		// define a slice of string
		reploFolder := "/Users/max/wkspc/git"
		repoSlice := []string{"gocore", "golinux", "gotask", "goluc"}
		// repoSlice := []string{"gocore", "golinux", "gotask"}
		// repoSlice := []string{"gocore"}

		// iterate over the slice
		for _, repoName := range repoSlice {
			repoPath := fmt.Sprintf("%s/%s", reploFolder, repoName)
			// print the value
			logx.Infof("Initiating git actions on : %s", repoPath)
			output, err := run.RunOnLocal(command)
			if err != nil {
				logger.Errorf("failed to initiate git actions for '%s': %v > %s", repoName, err, output)
			}
			logger.Infof("git actions done on: %s", repoName)

		}
	},
}
