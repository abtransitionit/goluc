/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package yaml

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var ListSDesc = "list applied/authorized yaml or manifest."
var ListLDesc = ListSDesc

// root Command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: ListSDesc,
	Long:  ListLDesc,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) != 1 {
	// 		return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the node, got %d", len(args))
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info("list applied/authorized yaml or manifest")
	},
}
