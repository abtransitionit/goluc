/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ns

import (
	"github.com/spf13/cobra"
)

// var
var localFlag bool

// Description
var epSDesc = "manage k8s namespaces."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "ns",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "uses by default the remote Helm client unless the flag is provided (it will use the local Helm client)")
	EpCmd.AddCommand(listCmd)
}
