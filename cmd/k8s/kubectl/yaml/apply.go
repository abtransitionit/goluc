/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package yaml

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var ApplySDesc = "apply a yaml or manifest to a cluster."
var ApplyLDesc = ApplySDesc

// root Command
var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: ApplySDesc,
	Long:  ApplyLDesc,
	// Args: func(cmd *cobra.Command, args []string) error {
	// 	if len(args) != 1 {
	// 		return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the node, got %d", len(args))
	// 	}
	// 	return nil
	// },
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info("apply a yaml or manifest to a cluster")
	},
}
