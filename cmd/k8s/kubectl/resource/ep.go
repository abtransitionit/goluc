/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package resource

import (
	"context"
	"os"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/spf13/cobra"
)

// Description
var tokenSDesc = "manage k8s resourse(s)."
var tokenLDesc = tokenSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "resource",
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
			_, err := ovh.RefreshToken(context.Background(), logger)
			if err != nil {
				logger.Errorf("%v", err)
				os.Exit(1)
			}
		}

	},
}

func init() {
	EpCmd.Flags().BoolP("check", "c", false, "check if the token exists in the credential file")
	EpCmd.Flags().BoolP("refresh", "r", false, "Api Refresh the token (make an API call to get a new token and replace the existing one in the credential file)")
}
