/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package do

import (
	"github.com/spf13/cobra"

	"github.com/abtransitionit/gocore/logx"
)

var prefix string
var vmName string

// Description
var SDesc = "Do actions locally or on remote host."
var LDesc = SDesc + "\n" + `
This command allows to use Go functions via the a remote host as if it was the local host. I cabn be used to for example:

- Download file
- Detect file type
- ... and many more
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "do",
	Short: SDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger := logx.GetLogger()
		logger.Infof("%s", SDesc)

		// success
		cmd.Help()
		return nil

	},
}

func init() {
	EpCmd.AddCommand(downloadCmd)
	EpCmd.AddCommand(detectCmd)
	EpCmd.Flags().StringVarP(&vmName, "vm", "v", "", "VM name on which to do the download")
	// EpCmd.Flags().AddCommand(resolveUrl)

}
