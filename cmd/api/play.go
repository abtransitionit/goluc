/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package api

import (
	"context"
	"time"

	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var playSDesc = "Play some client request."
var playLDesc = playSDesc + ` xxx.`

// root Command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: playSDesc,
	Long:  playLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logx.GetLogger()
		logger.Info(playSDesc)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel() // Always call cancel to release resources

		test04(ctx, logger)
		// test03(ctx, logger)
		// test01(logger)

	},
}

var forceFlag bool

func init() {
	playCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	playCmd.Flags().BoolP("list", "l", false, "List all available phases")
	playCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	playCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
