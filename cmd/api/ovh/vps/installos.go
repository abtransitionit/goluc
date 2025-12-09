/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"context"
	"os"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/abtransitionit/gocore/ui"
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
		// define var
		var vpsName string
		// define ctx and logger
		ctx := context.Background()
		logger := logx.GetLogger()
		logger.Infof(getSDesc)

		// 1 - define VPS(s) to process
		if !allFlag {
			// 1 - get the list of VPS - ie. read the local VPS config file
			vpsList, err := ovh.GetVpsList()
			if err != nil {
				logger.Errorf("getting vps:list from configuration file: %v", err)
				os.Exit(1)
			}
			pList := ovh.GetPrintableVpsList(vpsList)
			list.PrettyPrintTable(pList)

			// Ask user which ID (to choose) from the printed list
			id, err := ui.AskUserInt("\nchoose vps (enter ID): ")
			if err != nil {
				logger.Errorf("invalid ID: %v", err)
				return
			}
			// 2 - define resource property from user choice (ID and output)
			// 21
			vpsName, err = list.GetFieldByID(pList, id, 0)
			if err != nil {
				logger.Errorf("failed to get item from ID: %s: %v", id, err)
				return
			}
			vpsDistroCid, err := list.GetFieldByID(pList, id, 2)
			if err != nil {
				logger.Errorf("failed to get item from ID: %s: %v", id, err)
				return
			}

			// log
			// logger.Debugf("%s (%s): Re-installing image: %s (%s - %s)", vpsName, vpsId, vpsDistroCid, distroName, vpsDistroId)
			logger.Debugf("%s : Re-installing image: %s ", vpsName, vpsDistroCid)
			// re-install vps

		} else {
			logger.Debug("installing all vps > Todo")
		}

		// 3 - re-install the VPS
		jsonResponse, err := ovh.VpsReinstall(ctx, vpsName, logger)
		if err != nil {
			logger.Errorf("failed to re-install VPS: %v", err)
			os.Exit(1)
		}
		jsonx.PrettyPrintColor(jsonResponse)
	},
}

// fetchFilterDisplayVps fetches VPS detail, applies field filtering, and prints the result
func init() {
	installCmd.Flags().BoolVar(&allFlag, "all", false, "Get all VPSs with details")
	installCmd.Flags().StringVar(&fieldFlag, "field", "", "Display only a specific field (e.g. displayName, memoryLimit)")
}

// // 22
// vpsId, err := list.GetFieldByID(pList, id, 1)
// if err != nil {
// 	logger.Errorf("failed to get item from ID: %s: %v", id, err)
// 	return
// }
// // 23
// // 3 - get the VPS:Distro:Name - ie. read the local VPS config file
// distroName, err := ovh.GetDistroName(vpsDistroCid)
// if err != nil {
// 	logger.Errorf("getting distro name from id %q: %v", vpsDistroCid, err)
// 	os.Exit(1)
// }
// // 4 - get the list of available image for this VPS
// imageDetailList, err := ovh.ImageAvailableGetList(ctx, vpsName, logger)
// if err != nil {
// 	logger.Errorf("getting vps:distro:list for vps %s > %v", vpsName, err)
// 	return
// }
// // jsonx.PrettyPrintColor(imageDetailList)
// // 5 - get the image id from the distro name
// var vpsDistroId string
// for _, item := range imageDetailList {
// 	if item.Name == distroName {
// 		vpsDistroId = item.Id
// 	}
// }
