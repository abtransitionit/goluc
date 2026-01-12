/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/spf13/cobra"
)

// root command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "display the kbe cluster confgiguration",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Print summary
		logger.Infof("Command Name: %s", cmdName)
		logger.Infof("Description: %s", sDesc)
		logger.Info("Nodes:")
		logger.Infof("Control Plane: %v", cfg.Node.ControlPlane)
		logger.Infof("Worker: %v", cfg.Node.Worker)
		logger.Infof("All: %v", cfg.Node.All)
		logger.Info("Go CLI:")
		// Cluster
		listName := []string{}
		for _, v := range cfg.GoCli.Cluster {
			listName = append(listName, v.Name)
		}
		logger.Infof("  Cluster: %v", listName)
		// ControlPlane
		listName = []string{}
		for _, v := range cfg.GoCli.ControlPlane {
			listName = append(listName, v.Name)
		}
		logger.Infof("  Control Plane: %v", listName)

		logger.Info("DA Packages:")
		// Repos
		listName = []string{}
		for _, v := range cfg.Da.Repo.Node {
			listName = append(listName, v.Name)
		}
		logger.Infof("  Repos: %v", listName)
		// ControlPlane
		cpNames := []string{}
		for _, v := range cfg.Da.Pkg.ControlPlane {
			cpNames = append(cpNames, v.Name)
		}
		logger.Infof("  Control Plane: %v", cpNames)
		// Nodes
		nodeNames := []string{}
		for _, v := range cfg.Da.Pkg.Node {
			nodeNames = append(nodeNames, v.Name)
		}
		logger.Infof("  Nodes: %v", nodeNames)
		// Required
		reqNames := []string{}
		for _, v := range cfg.Da.Pkg.Required {
			reqNames = append(reqNames, v.Name)
		}
		logger.Infof("  Required: %v", reqNames)

		logger.Info("Helm Repos:")
		for _, repo := range cfg.Helm.Repo {
			logger.Infof("  - %s", repo.Name)
		}
		logger.Info("Helm Releases:")
		for _, release := range cfg.Helm.Release {
			logger.Infof("  - %s (%s) in namespace %s", release.Name, release.Chart, release.Namespace)
		}

		logger.Info("Cluster:")
		logger.Infof("Pod CIDR: %s", cfg.Cluster.PodCidr)
		logger.Infof("Service CIDR: %s", cfg.Cluster.ServiceCidr)
		logger.Infof("CR socket: %s", cfg.Cluster.CrSocketName)
		logger.Infof("Version: %s (%s)", k8sLongVersion, k8sShortVersion)

		logger.Info("Conf displayed successfully")
		return nil
	},
}
