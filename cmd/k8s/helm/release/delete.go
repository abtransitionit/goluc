/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package realease

import (
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete a helm repo."
var deleteLDesc = deleteSDesc

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
