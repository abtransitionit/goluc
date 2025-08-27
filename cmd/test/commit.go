/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
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
		logx.Info(commitSDesc)
		command:='git checkout main ; git merge --no-edit dev; git push origin main; git checkout dev'
		// define a slice of string
		reploFolder := "/Users/max/wkspc/git"
		repoName := []string{"gocore", "golinux", "gotask", "goluc"}

		// iterate over the slice
		for _, v := range repoName {
			repoPath := fmt.Sprintf("%s/%s", reploFolder, v)
			// print the value
			logx.Infof("git actions on : %s", repoPath)
		}
	},
}
