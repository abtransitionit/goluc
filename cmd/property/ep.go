/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package property

import (
	"fmt"

	linuxproperty "github.com/abtransitionit/golinux/property"
	"github.com/spf13/cobra"
)

var remote string
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
var EpCmd = &cobra.Command{
	Use:   "property",
	Short: SDesc,
	// Args:  cobra.ExactArgs(1), // Require exactly one argument: the property name
	Args: cobra.MinimumNArgs(1), // Require at least one argument: the property name
	RunE: func(cmd *cobra.Command, args []string) error {
		property := args[0]
		params := args[1:] // slice of extra params

		var value string
		var err error
		value, err = linuxproperty.GetProperty(vmName, property, params...)
		if err != nil {
			return fmt.Errorf("%v", err)
		}
		fmt.Println(value)
		return nil
	},
}

func init() {
	// EpCmd.Flags().StringVarP(&remote, "remote", "r", "", "provide the name of the remote VM")
	EpCmd.Flags().StringVar(&vmName, "vm", "", "VM name to query remotely")

}
