/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package secret

import (
	"github.com/abtransitionit/goluc/cmd/k8s/kubectl_res/res/secret/htpwd"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage secrets."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "secret",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(ListCmd)
	EpCmd.AddCommand(htpwd.EpCmd)
	EpCmd.AddCommand(deleteCmd)
	EpCmd.AddCommand(describeCmd)
	EpCmd.AddCommand(yamlCmd)
}
