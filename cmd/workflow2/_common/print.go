/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

func GetPrintCmd(cmdPathName string) *cobra.Command {
	var (
		showConfigFile  bool
		showConfigTable bool
		showPhase       bool
		showTier        bool
	)

	cobraCmd := &cobra.Command{
		Use:   "print",
		Short: "Display workflow details (config, phases, or tiers)",
		RunE: func(cobraCmd *cobra.Command, args []string) error {

			// define logger
			logger := logx.GetLogger()

			// at least one flag required
			if !showConfigFile && !showConfigTable && !showPhase && !showTier {
				return fmt.Errorf("please specify one of: --config, --phase, or --tier")
			}

			// --- CONFIG AS TABLE ---
			if showConfigTable {
				config, err := viperx.GetViperx("wkf.conf.yaml", "workflow", cmdPathName, logger)
				if err != nil {
					return fmt.Errorf("getting config: %w", err)
				}

				configContent, err := config.GetContentAsTable()
				if err != nil {
					return fmt.Errorf("getting config content: %w", err)
				}

				logger.Info("Workflow Config view as table")
				list.PrettyPrintTable(configContent)
			}
			// --- CONFIG AS FILE ---
			if showConfigFile {
				config, err := viperx.GetViperx("wkf.conf.yaml", "workflow", cmdPathName, logger)
				if err != nil {
					return fmt.Errorf("getting config: %w", err)
				}

				configContent, err := config.GetContentAsString()
				if err != nil {
					return fmt.Errorf("getting config content: %w", err)
				}

				logger.Info("Workflow Config view as file")
				fmt.Println(configContent)
			}

			// --- WORKFLOW ---
			if showPhase || showTier {
				workflow, err := phase2.GetWorkflow("wkf.phase.yaml", cmdPathName, logger)
				if err != nil {
					return fmt.Errorf("getting workflow: %w", err)
				}

				// --- PHASE ---
				if showPhase {
					phaseView, err := workflow.GetPhaseView()
					if err != nil {
						return fmt.Errorf("getting phase table: %w", err)
					}

					logger.Info("Workflow Phase View")
					list.PrettyPrintTable(phaseView)
				}

				// --- TIER ---
				if showTier {
					tierView, err := workflow.GetTierView()
					if err != nil {
						return fmt.Errorf("getting tier table: %w", err)
					}

					logger.Info("Workflow Tier View")
					list.PrettyPrintTable(tierView)
				}
			}

			// success
			return nil
		},
	}

	// define flags
	cobraCmd.Flags().BoolVar(&showConfigTable, "configTable", false, "Display workflow config as txt file")
	cobraCmd.Flags().BoolVar(&showConfigFile, "configFile", false, "Display workflow config as a table")
	cobraCmd.Flags().BoolVar(&showPhase, "phase", false, "Display workflow phases")
	cobraCmd.Flags().BoolVar(&showTier, "tier", false, "Display workflow tiers")

	return cobraCmd
}
