/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"context"
	"os"
	"strings"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/abtransitionit/gocore/ui"
	"github.com/spf13/cobra"
)

// Description
var getliSDesc = "Get the list of availbale OVH VPS OS images/distros."
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

		vpsSliceName, err := ovh.GetVpsListName()
		if err != nil {
			logger.Errorf("getting vps:list from configuration file: %v", err)
			os.Exit(1)
		}

		// print list
		pList := "NameDynamic\n" + strings.Join(vpsSliceName, "\n")
		list.PrettyPrintTable(pList)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose vps (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}
		// define resource property from user choice
		vpsName, err := list.GetFieldByID(pList, id, 0)
		if err != nil {
			logger.Errorf("failed to get item from ID: %s: %v", id, err)
			return
		}
		// get the list of available distro for the selected vps
		vpsDistro, err := ovh.ImageAvailableGetList(ctx, vpsName, logger)
		if err != nil {
			logger.Errorf("getting vps:distro:list for vps %s > %v", vpsName, err)
			return
		}
		// print
		jsonx.PrettyPrintColor(vpsDistro)
	},
}

func init() {
	getliCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	getliCmd.Flags().BoolP("list", "l", false, "List all available phases")
	getliCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	getliCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
