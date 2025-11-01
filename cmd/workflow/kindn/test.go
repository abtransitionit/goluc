/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kindn

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/abtransitionit/gocore/yamlx"
	"github.com/spf13/cobra"
)

// root Command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "sDesc",
	RunE: func(cmd *cobra.Command, args []string) error {

		// define logger
		logger := logx.GetLogger()

		v, err := viperx.GetConfig("wkf.conf.yaml", "workflow", cmdName)
		if err != nil {
			return err
		}

		// The keys are already scoped to "workflow.kindn" by GetSection
		nodes := v.GetStringSlice("node")
		requiredPkgs := v.GetStringSlice("da.pkg.required")
		goCliList := v.Get("goCli")
		services := v.Get("service")
		envVars := v.Get("envar")

		// log
		pc, file, _, _ := runtime.Caller(0)
		pkg := path.Dir(runtime.FuncForPC(pc).Name())
		logger.Infof("Package: %s", pkg)
		logger.Infof("Package: %s", path.Dir(file))

		// log
		logger.Infof("nodes: %v", nodes)
		logger.Infof("requiredPkgs: %v", requiredPkgs)
		logger.Infof("goCliList: %v", goCliList)
		logger.Infof("services: %v", services)
		logger.Infof("envVars: %v", envVars)

		// log
		_, file, _, _ = runtime.Caller(0) // because it is not called directly but through GetSection
		workflowPath := filepath.Join(path.Dir(file), "phase.yaml")
		logger.Infof("Package: %s", workflowPath)

		workflow, err := yamlx.LoadFile[phase.Workflow2](workflowPath)
		if err != nil {
			log.Fatal(err)
		}

		workflow.Print()
		return nil

	},
}
