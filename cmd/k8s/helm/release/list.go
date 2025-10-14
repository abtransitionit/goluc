/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"github.com/spf13/cobra"
)

// Description
var listSDesc = "list helm repos."
var listLDesc = listSDesc

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
