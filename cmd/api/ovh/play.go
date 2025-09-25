/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

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
		// define the logger
		logger := logx.GetLogger()

		// log
		logger.Info(playSDesc)

		// define a ctx with timeout
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// sshKeyList(ctx, logger)
		// DetailMe(ctx, logger)
		// listVps(ctx, logger)
		// installVps(ctx, logger)
		vpsGetList(ctx, logger)
	},
}

func init() {
	playCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	playCmd.Flags().BoolP("list", "l", false, "List all available phases")
	playCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	playCmd.MarkFlagsMutuallyExclusive("list", "runall")
}

// test04(ctx, logger)
// test03(ctx, logger)
// test01(logger)
