/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package smalln

import (
	"fmt"

	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gotask/dnfapt"
	"github.com/abtransitionit/gotask/luc"
	"github.com/abtransitionit/gotask/vm"
	"github.com/spf13/cobra"
)

// root Command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "launch the worflow",
	RunE: func(cmd *cobra.Command, args []string) error {
		// define logger
		// logger := logx.GetLogger()

		// get the workflow yaml
		workflow, err := phase.GetWorkflow()
		if err != nil {
			return fmt.Errorf("getting workflow: %w", err)
		}

		// action
		workflow.Execute()

		//return
		return nil

	},
}

func init() {

}

func registerWorkflowFn() {
	// manually register all function of the workflow
	phase.RegisterSingleFunc("vm.CheckVmSshAccess", vm.CheckVmSshAccess)
	phase.RegisterSingleFunc("luc.DeployLuc", luc.DeployLuc)
	phase.RegisterSingleFunc("dnfapt.UpgradeVmOs", dnfapt.UpgradeVmOs)
	// solve pbs of function that does not have the type PhaseFunc2 - closure for example
	// phase.RegisterSingleFunc("dnfapt.UpdateVmOsApp", dnfapt.UpdateVmOsApp)
}
