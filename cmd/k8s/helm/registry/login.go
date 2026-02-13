/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package registry

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/spf13/cobra"
)

// Description
var loginSDesc = "login to a Helm OCI registry (mandatory before push or pull)"
var loginDesc = loginSDesc

// root Command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: loginSDesc,
	Long:  loginDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(loginSDesc)
		// ctx := context.Background()

		// get instance and operate
		i := helm.Resource{Type: helm.ResRegistry, Name: "ghcr"}
		err := i.Login("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
