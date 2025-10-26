/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"fmt"

	"github.com/spf13/cobra"
)

// root command
var confCmd = &cobra.Command{
	Use:   "conf",
	Short: "display the kbe cluster confgiguration",
	RunE: func(cmd *cobra.Command, args []string) error {

		// Load YAML file
		configFile := "/Users/max/wkspc/git/goluc/cmd/workflow/kbe/conf.yaml" // adjust path if needed
		logger.Debugf("Loading config... from %s", configFile)
		cfg, err := LoadConfig(configFile)
		if err != nil {
			return fmt.Errorf("failed to load config: %v", err)
		}

		// Print summary
		logger.Infof("Command Name: %s", cfg.CmdName)
		logger.Infof("Description: %s", cfg.Description)
		logger.Info("Nodes:")
		logger.Infof("Control Plane: %v", cfg.Nodes.ControlPlane)
		logger.Infof("Worker: %v", cfg.Nodes.Worker)

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
		for _, v := range cfg.Da.Package.ControlPlane {
			cpNames = append(cpNames, v.Name)
		}
		logger.Infof("  Control Plane: %v", cpNames)
		// Nodes
		nodeNames := []string{}
		for _, v := range cfg.Da.Package.Node {
			nodeNames = append(nodeNames, v.Name)
		}
		logger.Infof("  Nodes: %v", nodeNames)
		// Required
		reqNames := []string{}
		for _, v := range cfg.Da.Package.Required {
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
		logger.Infof("Version: %s (%s)", cfg.Cluster.Version.Long, cfg.Cluster.Version.Short)

		logger.Info("Conf displayed successfully")
		return nil
	},
}
