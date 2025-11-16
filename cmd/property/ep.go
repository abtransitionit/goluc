/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package property

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	lproperty "github.com/abtransitionit/golinux/mock/property"
	"github.com/spf13/cobra"
)

var targetName string
var propFlags = map[string]*bool{} // holds which flag was used - *string because some handlers need params

// Description
var SDesc = "Get a property from a target machine (locally or remotely)."
var LDesc = SDesc + "\n" + `
This command allows to get several kind of OS properties like.

- OsFamily
- OsType
- Kernel version
- ... and many more
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "property",
	Short: SDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		// define vars
		var propertyName string
		logger := logx.GetLogger()

		// 1 - get the selected property as flag
		count := 0
		for flag, ptr := range propFlags {
			if *ptr {
				propertyName = flag
				count++
			}
		}
		// 2 - check flag
		if count == 0 {
			// 21 - check missing flag
			return fmt.Errorf("you must specify one property flag (e.g., --cpu, --path)")
		} else if count > 1 {
			// 22 - check mulmtiple flag
			return fmt.Errorf("you must specify ONLY one property flag")
		}

		// 3 - get parameters if any
		propertyParams := args

		// log
		// logger.Debugf("selected property: %s", property)
		// logger.Debugf("params: %s", params)

		// 4 - get property
		value, err := lproperty.GetProperty(logger, targetName, propertyName, propertyParams...)
		// 4 - handle system error
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// 5 - handle success
		fmt.Println(value)
		return nil
	},
}

func init() {
	// static flags
	EpCmd.Flags().StringVarP(&targetName, "target", "t", "local", "Name of the target to query")

	// dynamic flags - create automatically one bool flag per property
	for prop := range lproperty.PropertyMap {
		var flag bool
		propFlags[prop] = &flag
		EpCmd.Flags().BoolVar(propFlags[prop], prop, false, "")
	}
}

// value, err = linuxproperty.GetProperty(targetName, property, params...)
