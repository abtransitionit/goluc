/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package common

import (
	"fmt"
	"path"
	"reflect"
	"runtime"
	"strings"

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

			// at least one 	flag required
			if !showConfigTxt && !showConfigTable && !showPhase && !showTier && !showFunction {
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
					// get tiers
					tiers, err := workflow.TopoSortByTier(logger)
					if err != nil {
						return fmt.Errorf("cannot sort tiers: %w", err)
					}

					tierView, err := workflow.GetTierView(tiers, logger)
					if err != nil {
						return fmt.Errorf("getting tier table: %w", err)
					}

					logger.Infof("Workflow %s (Tier View) to %s", workflow.Name, workflow.Description)
					list.PrettyPrintTable(tierView)
				}

			}

			// --- PHASE VIEW---
			if showFunction {
				registry := phase2.GetFnRegistry()

				keys := registry.List()

				// Build the table as a raw TSV string:
				// Header
				var b strings.Builder
				b.WriteString("Key\tPackage\tFunction\n")

				for _, key := range keys {
					fn, ok := registry.Get(key)
					if !ok {
						continue
					}

					fnVal := reflect.ValueOf(fn)
					ptr := fnVal.Pointer()
					rf := runtime.FuncForPC(ptr)

					pkg := "<??>"
					fnName := "<??>"

					if rf != nil {
						fullName := rf.Name() // github.com/.../node.CheckSshConf

						// Remove prefix
						prefix := "github.com/abtransitionit/"
						fullName = strings.TrimPrefix(fullName, prefix)
						// define var
						pkg = path.Dir(fullName)
						fnName = path.Base(fullName)
					}

					fmt.Fprintf(&b, "%s\t%s\t%s\n", key, pkg, fnName)
				}

				logger.Info("Function Registry (key + package + function)")

				list.PrettyPrintTable(b.String())
			}

			// success
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
