/*
Copyright © 2025 AB TRANSITION IT
*/
package common

import "github.com/spf13/cobra"

func SetInitSubCmd(epCmd *cobra.Command, cmdPathName string) {
	epCmd.AddCommand(
		GetViewCmd(cmdPathName),
		GetRunCmd(cmdPathName),
	)
}
