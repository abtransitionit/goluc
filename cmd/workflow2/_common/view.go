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
		showFunction    bool
	)

	cobraCmd := &cobra.Command{
		Use:   "view",
		Short: "Display workflow views (config, phases, or tiers)",
		RunE: func(cobraCmd *cobra.Command, args []string) error {

			// define logger
			logger := logx.GetLogger()

			// // at least one 	flag required
			// if !showConfigTxt && !showConfigTable && !showPhase && !showTier && !showFunction {
			// 	return fmt.Errorf("please specify one of: --config, --phase, or --tier")
			// }

			// If no flags provided, enable all by default
			if !showConfigTxt && !showConfigTable && !showPhase && !showTier && !showFunction {
				showConfigTable = true
				showPhase = true
				showTier = true
				showFunction = true
			}

			// 1 - load config YAML
			config, err := viperx.GetViperx("wkf.conf.yaml", "workflow", cmdPathName, logger)
			if err != nil {
				return fmt.Errorf("getting config: %w", err)
			}

			// 1 - load workflow YAML
			workflow, err := phase2.GetWorkflow("wkf.phase.yaml", cmdPathName, logger)
			if err != nil {
				return fmt.Errorf("getting workflow: %w", err)
			}

			// 2 - workflow config as table
			if showConfigTable {
				configContent, err := config.GetContentAsTable()
				if err != nil {
					return fmt.Errorf("getting config content: %w", err)
				}

				// print
				logger.Info("Workflow Config (table view)")
				list.PrettyPrintTable(configContent)
			}

			// 3 - workflow config as text file
			if showConfigTxt {
				configContent, err := config.GetContentAsString()
				if err != nil {
					return fmt.Errorf("getting config content: %w", err)
				}

				// print
				logger.Info("Workflow Config (file view)")
				fmt.Println(configContent)
			}

			// 4 - workflow phases
			if showPhase {
				// get phases
				phaseView, err := workflow.GetPhaseView()
				if err != nil {
					return fmt.Errorf("getting phase table: %w", err)
				}

				// print
				logger.Infof("Workflow %s (Phase View) to %s", workflow.Name, workflow.Description)
				list.PrettyPrintTable(phaseView)
			}

			// 5 - workflow functions
			if showFunction {
				// get registry
				registry := phase2.GetFnRegistry()

				// get functions
				functionView, err := workflow.GetFunctionView(cmdPathName, registry)
				if err != nil {
					return fmt.Errorf("getting function table: %w", err)
				}
				// print
				logger.Info("Workflow registred Function	")
				list.PrettyPrintTable(functionView)
			}

			// 6 - workflow tiers
			if showPhase {
				// get tiers
				tiers, err := workflow.TopoSortByTier(logger)
				if err != nil {
					return fmt.Errorf("cannot sort tiers: %w", err)
				}

				tierView, err := workflow.GetTierView(tiers, logger)
				if err != nil {
					return fmt.Errorf("getting tier table: %w", err)
				}

				// print
				logger.Infof("Workflow %s (Tier View) to %s", workflow.Name, workflow.Description)
				list.PrettyPrintTable(tierView)
			}

			return nil
		},
	}

	// define flags
	cobraCmd.Flags().BoolVarP(&showConfigTable, "cfgTable", "c", false, "Display workflow config file in a table")
	cobraCmd.Flags().BoolVarP(&showConfigTxt, "cfgTxt", "x", false, "Display workflow config file as a Txt file")
	cobraCmd.Flags().BoolVarP(&showPhase, "wphase", "p", false, "Display workflow phases")
	cobraCmd.Flags().BoolVarP(&showTier, "wtier", "t", false, "Display workflow tiers")
	cobraCmd.Flags().BoolVarP(&showFunction, "func", "f", false, "Display workflow GO functions used and registered")

	return cobraCmd
}
