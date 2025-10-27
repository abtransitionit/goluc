/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kindn

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/viperx"
	"github.com/spf13/cobra"
)

// root Command
var EpCmd = &cobra.Command{
	Use:   "kindn",
	Short: "sDesc",
	RunE: func(cmd *cobra.Command, args []string) error {

		// define logger
		logger := logx.GetLogger()

		// get config (package+global+local)
		v, err := viperx.GetSection("kindn")
		if err != nil {
			return err
		}

		// Bind flags and env vars
		viperx.BindFlags(cmd, v, "kindn")

		// log
		logger.Infof("%s", cmd.Short)

		// Default action
		cmd.Help()
		return nil
	},
}

func init() {
	EpCmd.AddCommand(testCmd)
}

// log current package
// pc, file, _, _ := runtime.Caller(0)
// pkg := path.Dir(runtime.FuncForPC(pc).Name())
// logger.Infof("Package: %s", pkg)
// logger.Infof("Package: %s", path.Dir(file))

// pc, _, _, _ := runtime.Caller(0)
// fn := runtime.FuncForPC(pc).Name() // e.g. "github.com/.../cmd/workflow/kindn.init"
// pkg := path.Base(path.Dir(fn))
// logger.Infof("Package: %s", pkg)
