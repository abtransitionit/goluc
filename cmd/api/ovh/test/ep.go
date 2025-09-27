/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"context"
	"fmt"
	"os"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/spf13/cobra"
)

// Description
var testSDesc = "test some API actions."
var tokenLDesc = testSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "test",
	Short: testSDesc,
	Long:  tokenLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// log
		logger.Infof(testSDesc)

		// get ssh key id
		keyId, err := ovh.GetSshKeyIdFromFileCached()
		if err != nil {
			logger.Errorf("failed to get ssh key id %v", err)
			os.Exit(1)
		}

		// get VPS:Ssh:PublicKey
		sshPubKey, err := ovh.SshKeyGetPublicKeyCached(context.Background(), logger, keyId)
		if err != nil {
			logger.Errorf("failed to get ssh public key %v", err)
			os.Exit(1)
		}
		fmt.Println(sshPubKey)

		// // api get ssh key detail
		// sshKeyDetail, err := ovh.SshKeyGetDetail(context.Background(), logger, keyId)
		// if err != nil {
		// 	logger.Errorf("failed to get ssh key detail %v", err)
		// 	os.Exit(1)
		// }
		// // filter the detail
		// field := "key"
		// sshPublicKey, err := jsonx.GetFilteredJson(context.Background(), logger, sshKeyDetail, field)
		// if err != nil {
		// 	logger.Errorf("failed to get ssh public key %v", err)
		// 	os.Exit(1)
		// }
		// fmt.Println(sshPublicKey)
		// jsonx.PrettyPrintColor(sshKeyDetail)
		os.Exit(1)

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
