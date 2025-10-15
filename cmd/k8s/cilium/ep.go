/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cilium

import (
	"github.com/spf13/cobra"
)

var (
	forceFlag bool
	localFlag bool
)

// Description
var epSDesc = "manage k8s resources using kubectl."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "cilium",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
	EpCmd.AddCommand(statusCmd)
}
