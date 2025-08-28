// file: gocore/cli/cli.go
package goprj

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var SDesc = "Manage go projects."
var LDesc = SDesc + "\n" + `
This command allows to act on a go project: vet, test, build, ...
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "goprj",
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
	// define the entry point for each workflow
	EpCmd.AddCommand(testCmd)
	EpCmd.AddCommand(vetCmd)
	EpCmd.AddCommand(buildCmd)
	EpCmd.AddCommand(deployCmd)
}
