/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kindn

import (
	"path"
	"runtime"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

// root Command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "sDesc",
	RunE: func(cmd *cobra.Command, args []string) error {

		// define logger
		logger := logx.GetLogger()

		v, err := viperx.GetSection("kindn")
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
		return nil

	},
}
