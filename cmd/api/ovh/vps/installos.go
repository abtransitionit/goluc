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
var installSDesc = "Re-install the OVH VPS Os image."
var installLDesc = installSDesc + `
- This command is used to Api re-install the OVH VPS Os image.
`

// root Command
var installCmd = &cobra.Command{
	Use:   "installos [vpsNameId]",
	Short: installSDesc,
	Long:  installLDesc,
	Args:  cobra.MaximumNArgs(1), // Require 0 or 1 arg: the VPS id
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		ctx := context.Background()
		logger := logx.GetLogger()
		logger.Infof(getSDesc)

		// 1 - define VPS(s) to process
		var vpsNameIdSlice []string
		if allFlag {
			// api get the list of vps
			list, err := ovh.VpsGetList(ctx, logger)
			if err != nil {
				logger.Errorf("failed to list VPS: %v", err)
				return
			}
			// set the list
			vpsNameIdSlice = list
		} else {
			if len(args) == 0 {
				logger.Warn("missing 1 arg: the VPS name ID, or use --all")
				cmd.Help()
				return
			}
			// set the list
			vpsNameIdSlice = []string{args[0]}
		}

		// 4 - process each VPS
		for _, id := range vpsNameIdSlice {
			logger.Infof("Processing VPS: %s", id)
			jsonResponse, err := ovh.VpsReinstall(ctx, id, logger)
			if err != nil {
				logger.Errorf("failed to re-install VPS: %v", err)
				os.Exit(1)
			}

			// 5 - print the response
			jsonx.PrettyPrintColor(jsonResponse)
		}
	},
}

// fetchFilterDisplayVps fetches VPS detail, applies field filtering, and prints the result
func init() {
	installCmd.Flags().BoolVar(&allFlag, "all", false, "Get all VPSs with details")
	installCmd.Flags().StringVar(&fieldFlag, "field", "", "Display only a specific field (e.g. displayName, memoryLimit)")
}
