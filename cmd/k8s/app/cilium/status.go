/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package cilium

import (
	"fmt"

	// cilium "github.com/abtransitionit/gocore/k8s-cilium"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8sapp/cilium"
	helm2 "github.com/abtransitionit/golinux/mock/k8scli/helm"
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

		// get helm host
		helmHost, err := helm2.GetHelmHost("local")
		if err != nil {
			logger.Errorf("%w", err)
			return
		}

		// display status
		output, err := cilium.CilumSvc.DisplayStatus("local", helmHost, logger)
		if err != nil {
			logger.Errorf("failed to get cilium status via the cilium cli > %v", err)
			return
		}
		// handle success
		fmt.Println(output)
	},
}

func init() {
	statusCmd.Flags().BoolP("installed", "i", false, "show installed Helm repos")
	statusCmd.Flags().BoolP("whitelist", "w", false, "show installable repos (organization whitelist)")
}
