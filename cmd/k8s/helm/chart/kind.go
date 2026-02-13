/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package chart

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/mock/k8scli/helm"
	"github.com/abtransitionit/goluc/cmd/k8s/shared"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var kindSDesc = "list the kind of K8s resources a chart will create into the K8s cluster"
var kindLDesc = kindSDesc + `
manage following use case:
	- chart is on the local FS where the helm client lives
	- chart is part of a chart repository configured in the helm client configuration files
`

// root Command
var kindCmd = &cobra.Command{
	Use:   "kind",
	Short: kindSDesc,
	Long:  kindLDesc,
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s repo add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(kindSDesc)
		// ctx := context.Background()

		// - get instance and operate
		i := helm.Resource{Type: helm.ResRepo}
		output, err := i.List("local", shared.HelmHost, logger)
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
		resName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get property repo:name from ID: %s > %w", id, err)
			return
		}

		// - get instance and operate
		i = helm.Resource{Type: helm.ResChart, Repo: resName}
		output, err = i.List("local", shared.HelmHost, logger)
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
		id, err = ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}
		// define resource property from user choice
		resQName, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get resource property from ID: %s: %v", id, err)
			return
		}

		// list the chart resource kind
		// - get instance and operate
		i = helm.Resource{Type: helm.ResChart, QName: resQName}
		output, err = i.ListResKind("local", shared.HelmHost, logger)
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

		// list the chart resource
		// - get instance and operate
		output, err = i.ListResName("local", shared.HelmHost, logger)
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

	},
}
