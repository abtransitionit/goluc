/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"os"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/spf13/cobra"
)

// Description
var getlvSDesc = "Get the list of OVH VPS define in the local configuration file (There is no API call)."
var getlvLDesc = getlvSDesc + `
- This command is used to get the list of OVH VPS info statically defined in the local configuration file.
- This command also add a dynamic (computated) field.
`

// root Command
var getlvCmd = &cobra.Command{
	Use:   "getlv",
	Short: getlvSDesc,
	Long:  getlvLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define the logger
		logger := logx.GetLogger()

		// log
		logger.Infof(getlvSDesc)

		// get the infos of all vps (ie. read the local VPS config file)
		vpsList, err := ovh.GetListVps()
		if err != nil {
			logger.Errorf("get vps:list from configuration file failed: %v", err)
			os.Exit(1)
		}

		jsonx.PrettyPrintColor(vpsList)

	},
}

func init() {
	getlvCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	getlvCmd.Flags().BoolP("list", "l", false, "List all available phases")
	getlvCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	getlvCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
