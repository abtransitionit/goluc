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

func GetViewCmd(cmdPathName string) *cobra.Command {
	var (
		showConfigTxt   bool
		showConfigTable bool
		showPhase       bool
		showTier        bool
	)

	cobraCmd := &cobra.Command{
		Use:   "view",
		Short: "Display workflow views (config, phases, or tiers)",
		RunE: func(cobraCmd *cobra.Command, args []string) error {

			// define logger
			logger := logx.GetLogger()

			// at least one flag required
			if !showConfigTxt && !showConfigTable && !showPhase && !showTier {
				return fmt.Errorf("please specify one of: --config, --phase, or --tier")
			}

			// --- WORKFLOW CONFIG AS TABLE ---
			if showConfigTable {
				config, err := viperx.GetViperx("wkf.conf.yaml", "workflow", cmdPathName, logger)
				if err != nil {
					return fmt.Errorf("getting config: %w", err)
				}

				configContent, err := config.GetContentAsTable()
				if err != nil {
					return fmt.Errorf("getting config content: %w", err)
				}

				logger.Info("Workflow Config (table view)")
				list.PrettyPrintTable(configContent)
			}
			// --- WORKFLOW CONFIG AS FILE ---
			if showConfigTxt {
				config, err := viperx.GetViperx("wkf.conf.yaml", "workflow", cmdPathName, logger)
				if err != nil {
					return fmt.Errorf("getting config: %w", err)
				}

				configContent, err := config.GetContentAsString()
				if err != nil {
					return fmt.Errorf("getting config content: %w", err)
				}

				logger.Info("Workflow Config (file view)")
				fmt.Println(configContent)
			}

			// --- WORKFLOW ---
			if showPhase || showTier {
				workflow, err := phase2.GetWorkflow("wkf.phase.yaml", cmdPathName, logger)
				if err != nil {
					return fmt.Errorf("getting workflow: %w", err)
				}

				// --- PHASE VIEW---
				if showPhase {
					phaseView, err := workflow.GetPhaseView()
					if err != nil {
						return fmt.Errorf("getting phase table: %w", err)
					}

					logger.Infof("Workflow %s (Phase View) to %s", workflow.Name, workflow.Description)
					list.PrettyPrintTable(phaseView)
				}

				// --- TIER VIEW---
				if showTier {
					tierView, err := workflow.GetTierView()
					if err != nil {
						return fmt.Errorf("getting tier table: %w", err)
					}

					logger.Infof("Workflow %s (Tier View) to %s", workflow.Name, workflow.Description)
					list.PrettyPrintTable(tierView)
				}
			}

			// success
			return nil
		},
	}

	// define flags
	cobraCmd.Flags().BoolVar(&showConfigTable, "cfgTable", false, "Display workflow config file in a table")
	cobraCmd.Flags().BoolVar(&showConfigTxt, "cfgTxt", false, "Display workflow config file as a Txt file")
	cobraCmd.Flags().BoolVar(&showPhase, "wphase", false, "Display workflow phases")
	cobraCmd.Flags().BoolVar(&showTier, "wtier", false, "Display workflow tiers")

	return cobraCmd
}
