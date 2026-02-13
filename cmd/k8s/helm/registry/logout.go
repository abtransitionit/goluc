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
var logoutSDesc = "logout from a Helm OCI registry"
var logoutDesc = logoutSDesc

// root Command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: logoutSDesc,
	Long:  logoutDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(loginSDesc)
		// ctx := context.Background()

		// 1. Define the parameters in a map
		param := map[string]string{
			"DnsOrIp": "ghcr.io",
		}

		// 2 - get instance and operate
		i := helm.Resource{Type: helm.ResRegistry, Param: param}
		// _, err := i.Login("local", "shared.HelmHost", logger)
		_, err := i.Logout("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
