/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package registry

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
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

		// - get instance and operate
		i := helm.Resource{Type: helm.ResRegistry}
		output, err := i.List("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
		}

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from user choice
		resName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}

		// get instance and operate
		i = helm.Resource{Type: helm.ResRegistry, Name: resName}
		err = i.Login("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
