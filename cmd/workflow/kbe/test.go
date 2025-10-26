/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

// root command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test some code",
	RunE: func(cmd *cobra.Command, args []string) error {

		logger.Info("Loading config...")

		// Load YAML file
		configFile := "/Users/max/wkspc/git/goluc/cmd/workflow/kbe/conf.yaml" // adjust path if needed
		data, err := os.ReadFile(configFile)
		if err != nil {
			return fmt.Errorf("failed to read config file: %v", err)
		}

		var cfg Config
		if err := yaml.Unmarshal(data, &cfg); err != nil {
			return fmt.Errorf("failed to unmarshal YAML: %v", err)
		}

		// Print summary
		logger.Infof("Command Name: %s", cfg.CmdName)
		logger.Infof("Description: %s", cfg.Description)
		logger.Infof("K8s Version: %s (%s)", cfg.K8sVersion, cfg.K8sVersionShort)
		logger.Infof("Nodes:")
		logger.Infof("  Control Plane: %v", cfg.Nodes.ControlPlane)
		logger.Infof("  Worker: %v", cfg.Nodes.Worker)

		logger.Infof("Go CLI tools:")
		for _, cli := range cfg.Go.Cli {
			logger.Infof("  - %s %s", cli.Name, cli.Version)
		}

		logger.Infof("DA Packages:")

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

		logger.Infof("Helm Repos:")
		for _, repo := range cfg.Helm.Repo {
			logger.Infof("  - %s", repo.Name)
		}
		logger.Infof("Helm Releases:")
		for _, release := range cfg.Helm.Release {
			logger.Infof("  - %s (%s) in namespace %s", release.Name, release.Chart, release.Namespace)
		}

		logger.Info("Cluster:")
		logger.Infof("  Pod CIDR: %s", cfg.Cluster.PodCidr)
		logger.Infof("  Service CIDR: %s", cfg.Cluster.ServiceCidr)
		logger.Infof("  CR socket: %s", cfg.Cluster.CrSocketName)

		logger.Info("Test executed successfully")
		return nil
	},
}
