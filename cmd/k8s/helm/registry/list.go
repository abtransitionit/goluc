/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package registry

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var (
	flagAuth bool
)

// Description
var listSDesc = "list known OCI registry."
var listDesc = listSDesc

// root Command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listSDesc)
		// ctx := context.Background()

		// 1 - get instance
		i := helm.Resource{Type: helm.ResRegistry}

		// 2 - logic - choose the method to call based on the flag
		instanceFn := i.List // default (no flag)
		switch {
		case flagAuth:
			instanceFn = i.ListAuth // case flagAuth is used
		}

		// 3 - operate on the instance
		output, err := instanceFn("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
			logger.Infof("other options: --auth")
		}

	},
}

func init() {
	ListCmd.Flags().BoolVar(&flagAuth, "auth", false, "list repos authorized to be installed")
}
