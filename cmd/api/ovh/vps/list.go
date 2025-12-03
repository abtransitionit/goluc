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
var listSDesc = "List all OVH VPS name id."
var listLDesc = listSDesc + ` xxx.`

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define the logger
		logger := logx.GetLogger()

		// log
		logger.Infof(getSDesc)

		// api get the list of vps name id
		vpsList, err := ovh.VpsGetList(context.Background(), logger)
		if err != nil {
			logger.Errorf("get vps:list failed: %v", err)
			os.Exit(1)
		}

		jsonx.PrettyPrintColor(vpsList)

	},
}

func init() {
	listCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	listCmd.Flags().BoolP("list", "l", false, "List all available phases")
	listCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	listCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
