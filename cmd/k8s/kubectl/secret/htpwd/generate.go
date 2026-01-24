/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package htpwd

import (
	"strings"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var generateSDesc = "generate and display tha htpwd."
var generateLDesc = generateSDesc + "\n" + `
The follwoing information are needed
- namespace (choose from the list)
- user name
`

// root Command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: generateSDesc,
	Long:  generateLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// list authorized manifest
		// - get instance and operate
		output, err := kubectl.List(kubectl.ResNS, "local", shared.HelmHost, logger)
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
		resNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get res name from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s ", resNs)

		// Ask user the user name
		resUserName := ui.AskUserString("\ndefine the secret user name (often same as namespace): ")
		// logger.Infof("creating secret in selected ns:%s with UserName:%s", resName, resUserName)

		logger.Infof("user name is: %s ", resUserName)
		// - get instance and operate
		i := kubectl.Resource{Type: kubectl.ResSecret, Ns: resNs, UserName: resUserName}
		output, err = i.Generate("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		logger.Info(strings.TrimSpace(output))
	},
}
