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

var (
	allFlag   bool
	fieldFlag string
)

// Short description
var getSDesc = "Get information on any or all OVH VPS."

// Long description
var getLDesc = getSDesc + `
This command sends a GET request to the OVH API to retrieve information about your VPS instances,
including their name and ID.
`

// // Description
// var getSDesc = "Get informations on any or all OVH VPS."
// var getLDesc = getSDesc + `
// - This command is used to Api get the list of OVH VPS name id.
// `

// root Command
var getCmd = &cobra.Command{
	Use:   "get [vpsNameId]",
	Short: getSDesc,
	Long:  getLDesc,
	Args:  cobra.MaximumNArgs(1), // Require 0 or 1 arg: the VPS id
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		ctx := context.Background()
		logger := logx.GetLogger()
		logger.Infof(getSDesc)

		// define VPS(s) to process
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

		// process each vps
		for _, id := range vpsNameIdSlice {
			// api get the VPS:detail
			vpsDetail, err := ovh.VpsGetDetail(ctx, logger, id)
			if err != nil {
				logger.Errorf("failed to get detail for VPS: %s: %v", id, err)
				os.Exit(1)
			}
			// filter the VPS:detail
			vpsInfo, err := jsonx.GetFilteredJson(ctx, logger, vpsDetail, fieldFlag)
			if err != nil {
				logger.Errorf("failed to get detail for VPS: %s: %v", id, err)
				// if --all is not set, stop
				if !allFlag {
					return
				}
				continue
			}

			// Display VPS detail
			jsonx.PrettyPrintColor(vpsInfo)
		}
	},
}

// fetchFilterDisplayVps fetches VPS detail, applies field filtering, and prints the result
func init() {
	getCmd.Flags().BoolVar(&allFlag, "all", false, "Get detailed information for all VPSs")
	getCmd.Flags().StringVar(&fieldFlag, "field", "", "Display only a specific field (e.g. displayName, memoryLimit)")
}
