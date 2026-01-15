/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package module

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var forceFlag bool
var localFlag bool

// Description
var epSDesc = "manage linux os kernel moduless."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "module",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(epSDesc)
		cmd.Help()
	},
}

func init() {
	// repoListCmd := *repo.DescribeCmd
	// repoListCmd.Use = "list"

	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
	EpCmd.AddCommand(describeCmd)
	EpCmd.AddCommand(createCmd)
}
