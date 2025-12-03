/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"context"
	"os"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/spf13/cobra"
)

// Description
var getliSDesc = "Get the list of availbale OVH VPS OS images."
var getliLDesc = getliSDesc + `
- This command is used to get the list all the OVH VPS OS images available.
`

// root Command
var getliCmd = &cobra.Command{
	Use:   "getli",
	Short: getliSDesc,
	Long:  getliLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define the logger
		logger := logx.GetLogger()
		ctx := context.Background()

		// log
		logger.Infof(getliSDesc)

		// get the infos of all vps (ie. read the local VPS config file)
		vpsList, err := ovh.ImageGetList(ctx, logger)
		if err != nil {
			logger.Errorf("getting available image:list : %v", err)
			os.Exit(1)
		}

		jsonx.PrettyPrintColor(vpsList)

	},
}

func init() {
	getliCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	getliCmd.Flags().BoolP("list", "l", false, "List all available phases")
	getliCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	getliCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
