/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"context"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/spf13/cobra"
)

var (
	allFlag   bool
	fieldFlag string
)

// Description
var getSDesc = "Api get the list of OVH VPS."
var getLDesc = getSDesc + `
- This command is used to Api get the list of OVH VPS name id.
`

// root Command
var getCmd = &cobra.Command{
	Use:   "get [vpsNameId]",
	Short: getSDesc,
	Long:  getLDesc,
	Args:  cobra.MaximumNArgs(1), // Require 0 or 1 arg: the VPS id
	Run: func(cmd *cobra.Command, args []string) {

		// define the logger
		logger := logx.GetLogger()

		// log
		logger.Infof(getSDesc)

		// case 1: --all
		if allFlag {
			// api get the list of vps name id
			vpsList, err := ovh.VpsGetList(context.Background(), logger)
			if err != nil {
				logger.Errorf("get vps:list failed: %v", err)
				return
			}
			// loop over all vps name id
			for _, vpsNameId := range vpsList {
				// api get the infos of a vps
				vpsDetail, err := ovh.VpsGetDetail(context.Background(), logger, vpsNameId)
				if err != nil {
					logger.Errorf("get detail for %s failed: %v", vpsNameId, err)
					continue
				}
				// get the infos of a vps
				vpsDetail, err = ovh.GetVpsDetailFiltered(context.Background(), logger, vpsDetail, fieldFlag)
				if err != nil {
					logger.Errorf("get vps:detail for vps %s failed: %v", vpsNameId, err)
					return
				}
				// display
				jsonx.PrettyPrintColor(vpsDetail)

			}
			return
		}
		// case 2: single VPS
		if len(args) == 0 {
			logger.Warn("missing 1 arg: the VPS name ID, or use --all to fetch detail infos for all VPSes")
			cmd.Help()
			return
		}

		// define the vps
		vpsNameId := args[0]

		// api get the infos of a vps
		vpsDetail, err := ovh.VpsGetDetail(context.Background(), logger, vpsNameId)
		if err != nil {
			logger.Errorf("get detail for %s failed: %v", vpsNameId, err)
			return
		}

		// get the infos of a vps
		vpsDetail, err = ovh.GetVpsDetailFiltered(context.Background(), logger, vpsDetail, fieldFlag)
		if err != nil {
			logger.Errorf("get vps:detail for vps %s failed: %v", vpsNameId, err)
			return
		}

		jsonx.PrettyPrintColor(vpsDetail)

	},
}

func init() {
	getCmd.Flags().BoolVar(&allFlag, "all", false, "Get all VPSs with details")
	getCmd.Flags().StringVar(&fieldFlag, "field", "", "Display only a specific field (e.g. displayName, memoryLimit)")
}
