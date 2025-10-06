/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package realease

import (
	"github.com/spf13/cobra"
)

// Description
var addSDesc = "add a helm repo."
var addLDesc = addSDesc

// root Command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: addSDesc,
	Long:  addLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
