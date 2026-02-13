/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package helm

import (
	"github.com/abtransitionit/goluc/cmd/k8s/helm/chart"
	"github.com/abtransitionit/goluc/cmd/k8s/helm/conf"
	"github.com/abtransitionit/goluc/cmd/k8s/helm/registry"
	"github.com/abtransitionit/goluc/cmd/k8s/helm/release"
	"github.com/abtransitionit/goluc/cmd/k8s/helm/repo"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage k8s resources with helm"
var epLDesc = epSDesc + `
- This command is used to run action(s) on OVH resource(s) using an OVH client'srequest.
`

// root Command
var EpCmd = &cobra.Command{
	Use:   "helm",
	Short: epSDesc,
	Long:  epLDesc,
}

func init() {
	EpCmd.AddCommand(repo.EpCmd)
	EpCmd.AddCommand(release.EpCmd)
	EpCmd.AddCommand(chart.EpCmd)
	EpCmd.AddCommand(conf.EpCmd)
	EpCmd.AddCommand(registry.EpCmd)
}
