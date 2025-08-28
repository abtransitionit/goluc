package luc

import (
	"context"
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/gocore/syncx"
)

func deleteOnSingleVm(logger logx.Logger, vmName string) (string, error) {

	cli := "sudo rm -f /usr/local/bin/goluc"
	_, err := run.RunCliSsh(vmName, cli)
	if err != nil {
		return "", err
	}
	return "", nil
}

// NAme: createSliceFunc
//
// Description: create the slice of tasks
//
// Parameters:
// - l: logger
// - targets: list of targets
//
// Returns:
//
// - slice of syncx.Func
//
// Notes:
//
// - as many tasks as there are VMs
// - Only VM targets are included; others are skipped with a warning.
func createSliceDelete(logger logx.Logger, targets []phase.Target) []syncx.Func {
	var tasks []syncx.Func

	for _, t := range targets {
		if t.Type() != "Vm" {
			continue
		}

		vm, ok := t.(*phase.Vm)
		if !ok {
			logger.Warnf("🅣 Target %s is not a VM, skipping", t.Name())
			continue
		}

		vmCopy := vm // capture for closure
		tasks = append(tasks, func() error {
			if _, err := deleteOnSingleVm(logger, vmCopy.Name()); err != nil {
				logger.Errorf("🅣 Failed to delete LUC on VM %s: %v", vmCopy.Name(), err)
				return err
			}

			logger.Infof("🅣 VM %s deleted LUC successfully", vmCopy.Name())
			return nil
		})
	}

	return tasks
}

// name: UpgradeVmOs
//
// description: the overall task.
//
// Notes:
// - Each target must implement the Target interface.
func DeleteOnVm(ctx context.Context, logger logx.Logger, targets []phase.Target, cmd ...string) (string, error) {

	logger.Info("🅣 Starting phase: DeleteOnVm")
	// check paramaters
	if len(targets) == 0 {
		logger.Warn("🅣 No targets provided to : DeleteOnVm")
		return "", nil
	}

	// Build slice of functions
	tasks := createSliceDelete(logger, targets)

	// Log number of tasks
	logger.Infof("🅣 Phase DeleteOnVm has %d concurent tasks", len(tasks))

	// Run tasks in the slice concurrently
	if errs := syncx.RunConcurrently(ctx, tasks); errs != nil {
		return "", errs[0] // return first error encountered
	}

	return fmt.Sprintf("🅣 Terminated phase DeleteOnVm on %d VM(s)", len(tasks)), nil
}
