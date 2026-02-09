/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package mnf

import (
	"regexp"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var ListSDesc = "list authorized yaml/manifest."
var ListLDesc = ListSDesc

// root Command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: ListSDesc,
	Long:  ListLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResManifest}
		output, err := i.ListAuth("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// customize display (kepp only url:filename)
		output2 := regexp.MustCompile(`https://\S+/(\S+)`).ReplaceAllString(output, "$1")
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output2)
		}
	},
}
