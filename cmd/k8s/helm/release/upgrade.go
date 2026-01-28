/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package release

import (
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/golinux/mock/k8scli/kubectl"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/spf13/cobra"
)

// Description
var upgradeSDesc = "upgrade a release from a chart of an authorized repo."
var upgradeLDesc = upgradeSDesc

// root Command
var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: upgradeSDesc,
	Long:  upgradeLDesc,
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(upgradeSDesc)
		// ctx := context.Background()

		// list installed release
		// - get instance and operate
		i := helm.Resource{Type: helm.ResRelease}
		output, err := i.List("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			logger.Infof("none installed")
		} else {

		}

		// list configured repos
		// - get instance and operate
		i = helm.Resource{Type: helm.ResRepo}
		output, err = i.List("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
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
		repoName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property from ID: %s > %v", id, err)
			return
		}

		// list the repo's charts
		// - get instance and operate
		i = helm.Resource{Type: helm.ResChart, Repo: repoName}
		output, err = i.List("local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
		}

		// Ask user which ID (to choose) from the printed list
		id, err = ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}
		// define resource property from user choice
		chartQName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}
		// define resource property from user choice
		chartVersion, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// log
		logger.Infof("selected item: %s - %s", chartQName, chartVersion)
		// list ns
		// - get instance and operate
		output, err = kubectl.List(kubectl.ResNS, "local", shared.HelmHost, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// - print
		if list.CountNbLine(output) == 1 {
			return
		} else {
			list.PrettyPrintTable(output)
		}

		// Ask user which ID (to choose) from the printed list
		id, err = ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}
		// define resource property from user choice
		releaseNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}
		// log
		logger.Infof("selected item: %s : %s ", releaseNs, chartQName)

		// Ask user the release prefix
		releasePrefix := ui.AskUserString("\ndefine a release prefix (name is <prefix>-" + repoName + "): ")
		// log
		logger.Infof("upgrading	release: %s-%s from chart %s version %s in namespace %s", releasePrefix, repoName, chartQName, chartVersion, releaseNs)

	},
}
