/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"github.com/spf13/cobra"
)

// root Command
var EpCmd = &cobra.Command{
	Use:   cmdName,
	Short: sDesc,
	RunE: func(cmd *cobra.Command, args []string) error {
		logger.Infof("%s", sDesc)

		// Default action
		cmd.Help()
		return nil
	},
}

func init() {
	EpCmd.AddCommand(runCmd)
	EpCmd.AddCommand(showCmd)
	EpCmd.AddCommand(confCmd)
}

// PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
// 	// viper.SetConfigFile("/Users/max/wkspc/git/goluc/cmd/workflow/kbe/conf.yaml")
// 	// if err := viper.ReadInConfig(); err != nil {
// 	// 	return fmt.Errorf("failed to read config: %w", err)
// 	// }
// 	// // Unmarshal into struct
// 	// if err := viper.Unmarshal(&cfg); err != nil {
// 	// 	return fmt.Errorf("failed to parse config: %w", err)
// 	// }
// 	// // This runs before ANY subcommand (run/show/conf) - Load YAML file
// 	// configFile := "/Users/max/wkspc/git/goluc/cmd/workflow/kbe/conf.yaml" // adjust path if needed
// 	// logger.Debugf("Loading config... from %s", configFile)
// 	// var err error
// 	// cfg, err = loadConfig(configFile)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to load config: %v", err)
// 	// }
// 	return nil
// },
