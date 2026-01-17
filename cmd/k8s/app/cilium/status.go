/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cilium

import (
	"fmt"

	// cilium "github.com/abtransitionit/gocore/k8s-cilium"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8sapp/cilium"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var statusSDesc = "display the status of CNI:cilium components."

var statusLDesc = statusSDesc

// root Command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: statusSDesc,
	Long:  statusLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()

		// get status
		// - get instance and operate
		output, err := cilium.CilumSvc.GetStatus("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		fmt.Println(output)
	},
}

func init() {
	statusCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	statusCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
