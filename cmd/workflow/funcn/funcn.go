package funcn

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
)

// Package variables : confifg2
var (
	cmdName   = "funcn" // the app name - should also be the workflow name
	shortDesc = "workflow to test function registry."
)

type closure func(ctx context.Context, l logx.Logger) error

var fnRegistry map[string]closure

func init() {
	// fnRegistry = make(map[string]phase.PhaseFunc3)

	// define the function used by the phase of the workflow
	// fnRegistry = map[string]phase.PhaseFunc3{
	// 	"vm.CheckVmSshAccess": func(ctx context.Context, nodes []string, logger logx.Logger) error {
	// 		return callfn.CheckVmSshAccessReal(nodes)
	// 	}),
	// 	"luc.DeployLuc":      callfn.DeployLuc,
	// 	"dnfapt.UpgradeVmOs": callfn.UpgradeVmOs,
	// }
}

// phase.RegisterSingleFunc(fnRegistry, "vm.CheckVmSshAccess", callfn.CheckVmSshAccess)
// phase.RegisterSingleFunc(fnRegistry, "luc.DeployLuc", callfn.DeployLuc)
// phase.RegisterSingleFunc(fnRegistry, "dnfapt.UpgradeVmOs", callfn.UpgradeVmOs)
// phase.RegisterSingleFunc(fnRegistry, "dnfapt.UpdateVmOsApp", callfn.UpdateVmOsApp)
