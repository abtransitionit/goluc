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
var (
	flagAuth bool
)

// Description
var listSDesc = "list authorized/installed repos."
var listLDesc = listSDesc

// root Command
var ListRCmd = &cobra.Command{
	Use:   "listr",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// 1 - get instance
		i := helm.Resource{Type: helm.ResRepo}

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
	ListRCmd.Flags().BoolVar(&flagAuth, "auth", false, "list repos authorized to be installed")
}
