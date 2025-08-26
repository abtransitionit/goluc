/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package property

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/property"
	"github.com/spf13/cobra"
)

var remote string

// Description
var SDesc = "Manage a property of an OS or a VM."
var LDesc = SDesc + "\n" + `
This command allows to get several kind of OS properties like.

- OsFamily
- OsType
- Kernel version
- ... and many more
`

// root Command
var PropertyCmd = &cobra.Command{
	Use:   "property",
	Short: SDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logx.GetLogger()
		logger.Infof("%s", SDesc)
		// check parameters
		if len(args) == 0 {
			logger.Info("No property name provided.")
			cmd.Help()
			return nil
		}

		// here arg is provided: This is the first argument
		propertyName := args[0]

		if remote == "" {
			logger.Infof("look up property %s for the local machine", propertyName)
			propertyValue, err := property.GetPropertyLocal("cpu")
			if err != nil {
				logger.Infof("impossible")
				return nil
			}
			logger.Infof("value : %s", propertyValue)

			return nil
		}

		// here remote is provided
		logger.Infof("look up property %s for the remote machine > %s", propertyName, remote)
		// propertyValue, err := property.GetPropertyRemote(remote, "cpu")
		// if err != nil {
		// 	logger.Infof("impossible")
		// 	return nil
		// }
		// logger.Infof("value : %s", propertyValue)

		// // Get the property
		// if val, err := GetProperty(propertyName); err == nil {
		// 	logger.Infof("%s: %s", propertyName, val)
		// }

		// success
		return nil

	},
}

func init() {
	PropertyCmd.Flags().StringVarP(&remote, "remote", "r", "", "provide the name of the remote VM")
}
