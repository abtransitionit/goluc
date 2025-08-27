/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package prop

import (
	"fmt"

	linuxproperty "github.com/abtransitionit/golinux/property"
	"github.com/spf13/cobra"
)

var vmName string

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
var PropCmd = &cobra.Command{
	Use:   "prop",
	Short: SDesc,
	Args:  cobra.ExactArgs(1), // Require exactly one argument: the property name
	RunE: func(cmd *cobra.Command, args []string) error {
		property := args[0]

		var value string
		var err error
		value, err = linuxproperty.GetProperty(vmName, property)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		// Print the value (both for user and logs)
		fmt.Println(value)
		// logger.Infof("Property '%s' = %s", property, value)
		return nil

	},
}

func init() {
	// Register the --vm flag for the prop command
	PropCmd.Flags().StringVar(&vmName, "vm", "", "VM name to query remotely")
}
