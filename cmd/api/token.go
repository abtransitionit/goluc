/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package api

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/spf13/cobra"
)

// Description
var tokenSDesc = "manage the OAuth2 OVH token related to the Servcie Account (aka. Client)."
var tokenLDesc = tokenSDesc + ` xxx.`

// root Command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: tokenSDesc,
	Long:  tokenLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define the logger
		logger := logx.GetLogger()
		// get flags
		check, _ := cmd.Flags().GetBool("check")
		refresh, _ := cmd.Flags().GetBool("refresh")

		// if no flag is provided, show help
		if !check && !refresh {
			cmd.Help()
			return
		}

		// handle flags
		if check {
			ovh.CheckTokenExist(context.Background(), logger)
		}
		if refresh {
			ovh.RefreshToken(context.Background(), logger)
		}

	},
}

func init() {
	tokenCmd.Flags().BoolP("check", "c", false, "display infos on token")
	tokenCmd.Flags().BoolP("refresh", "r", false, "Refresh a token (replace the existing one in the credential file)")
}
