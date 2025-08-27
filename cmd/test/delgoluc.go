/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
	"github.com/spf13/cobra"
)

var (
	golucPath = "/usr/local/bin/goluc"
	vmSlice   = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
)

var (
	defgolucSDesc = fmt.Sprintf("On all selected VMs: delte the file %s.", golucPath)
	delgolucLDesc = defgolucSDesc + ` xxx.`
)

// root Command
var delgolucCmd = &cobra.Command{
	Use:   "delgoluc",
	Short: defgolucSDesc,
	Long:  delgolucLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logger := logx.GetLogger()
		logger.Info(defgolucSDesc)
		command := fmt.Sprintf("sudo rm -f %s", golucPath)
		// define a slice of string

		// iterate over the slice
		for _, vmName := range vmSlice {
			// print the value
			logx.Infof("deleting : %s:%s with remote command: %s", vmName, golucPath, command)
			_, err := run.RunOnVm(vmName, command)
			if err != nil {
				logger.Errorf("failed deleting '%s:%s': %v", vmName, golucPath, err)
			}
			logx.Infof("%s : done", vmName)
		}
	},
}
