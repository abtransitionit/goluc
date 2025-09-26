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
var getlSDesc = "Get the list of OVH VPS define in the local configuration file (There is no API call)."
var getlLDesc = getlSDesc + `
- This command is used to get the list of OVH VPS info statically defined in the local configuration file.
- This command also add a dynamic (computated) field.
`

// root Command
var getlCmd = &cobra.Command{
	Use:   "getl",
	Short: getlSDesc,
	Long:  getlLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define the logger
		logger := logx.GetLogger()

		// log
		logger.Infof(getlSDesc)

		// get the infos of all vps (ie. read the locl VPS config file)
		vpsList, err := ovh.GetListVpsFromFile()
		if err != nil {
			logger.Errorf("get vps:list from configuration file failed: %v", err)
			os.Exit(1)
		}

		jsonx.PrettyPrintColor(vpsList)

	},
}

func init() {
	getlCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	getlCmd.Flags().BoolP("list", "l", false, "List all available phases")
	getlCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	getlCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
