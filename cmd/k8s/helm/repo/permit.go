/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var permitSDesc = "list repos in the whitelist/autorized to be installed."
var permitLDesc = permitSDesc

// root Command
var PermitCmd = &cobra.Command{
	Use:   "permit",
	Short: permitSDesc,
	Long:  permitLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get instance
		i := helm.Resource{Type: helm.ResRepo}
		// operate
		output, err := i.ListPermit("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

		// print
		list.PrettyPrintTable(output)

	},
}
