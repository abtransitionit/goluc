/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package htpwd

import (
	// "github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage htpasswd secrets for Basic Auth."
var epLDesc = epSDesc + "\n" + `
This kind of pwd works fine for Basic Auth in the  applications lioke:
  - Private Docker registry
  - Nginx ingresss
  - Website
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "htpwd",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(generateCmd)
	EpCmd.AddCommand(createCmd)
	EpCmd.AddCommand(deleteCmd)
}
