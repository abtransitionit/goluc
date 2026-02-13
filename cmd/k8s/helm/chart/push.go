/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var pushSDesc = "push a chart's artifact to an OCI registry"
var pushLDesc = pushSDesc + `
manage following use case:
	- the chart artifact (targz) is on the local FS where the helm client is installed
`

// root Command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: pushSDesc,
	Long:  pushLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s build add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(pushSDesc)
		// ctx := context.Background()

		// TODO: Hard coded now - idea: pick an artifact from a list
		// 1. Define the parameters in a map
		param := map[string]string{
			"folderSrcRoot": "$HOME/wkspc/git/k8s-manifest/helm-chart/",
		}
		// 2 - Ask user
		resName := ui.AskUserString(fmt.Sprintf(`which chart folder to build (inside %s):`, param["folderSrcRoot"]))
		// log
		// logger.Debugf("cli is : %s", cli)

		// 3- get instance and operate
		i := helm.Resource{Type: helm.ResChart, Name: resName, Param: param}
		err := i.Build("local", "local", logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

	},
}
