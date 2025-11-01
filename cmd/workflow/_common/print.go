/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"fmt"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

func GetPrintCmd(cmdName string) *cobra.Command {
	var (
		showConfig bool
		showPhase  bool
		showTier   bool
	)

	cmd := &cobra.Command{
		Use:   "print",
		Short: "Display workflow details (config, phases, or tiers)",
		RunE: func(cmd *cobra.Command, args []string) error {
			// at least one flag required
			if !showConfig && !showPhase && !showTier {
				return fmt.Errorf("please specify one of: --config, --phase, or --tier")
			}

			// --- CONFIG ---
			if showConfig {
				config, err := viperx.GetConfig("wkf.conf.yaml", "workflow", cmdName)
				if err != nil {
					return fmt.Errorf("getting config: %w", err)
				}

				table, err := config.GetTable()
				if err != nil {
					return fmt.Errorf("getting config table: %w", err)
				}

				fmt.Println("== Workflow Config ==")
				fmt.Println(table)
			}

			// --- WORKFLOW ---
			if showPhase || showTier {
				workflow, err := phase2.GetWorkflow(cmdName)
				if err != nil {
					return fmt.Errorf("getting workflow: %w", err)
				}

				// --- PHASE ---
				if showPhase {
					table, err := workflow.GetTablePhase()
					if err != nil {
						return fmt.Errorf("getting phase table: %w", err)
					}

					fmt.Println("== Workflow Phases ==")
					list.PrettyPrintTable(table)
				}

				// --- TIER ---
				if showTier {
					table, err := workflow.GetTableTier()
					if err != nil {
						return fmt.Errorf("getting tier table: %w", err)
					}

					fmt.Println("== Workflow Tiers ==")
					list.PrettyPrintTable(table)
				}
			}

			return nil
		},
	}

	// define flags
	cmd.Flags().BoolVar(&showConfig, "config", false, "Display workflow config")
	cmd.Flags().BoolVar(&showPhase, "phase", false, "Display workflow phases")
	cmd.Flags().BoolVar(&showTier, "tier", false, "Display workflow tiers")

	return cmd
}
