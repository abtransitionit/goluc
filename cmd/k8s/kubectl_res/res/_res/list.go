/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package res

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var (
	flagNs   bool
	flagNoNs bool
)

// Description
var listSDesc = "list all api resources."
var listLDesc = listSDesc

// root Command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// logic for the insatnce method
		listFn := kubectl.ListNs // default (avoid declaration)
		switch {
		case flagNoNs:
			listFn = kubectl.ListNoNs
		case flagNs:
			listFn = kubectl.ListNs
		default:
			// optional: fallback to generic List
			listFn = kubectl.List
		}
		// - get instance and operate
		output, err := listFn(kubectl.ResRes, "local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
			logger.Infof("other options: --ns, --cl")
		}

	},
}

func init() {
	ListCmd.Flags().BoolVar(&flagNs, "ns", false, "list namespaced resources (default)")
	ListCmd.Flags().BoolVar(&flagNoNs, "cl", false, "list clustered resources")

	// optional but recommended: make them mutually exclusive
	ListCmd.MarkFlagsMutuallyExclusive("ns", "cl")
}
