package luc

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/gocore/syncx"
)

func cleanTmpOnSingleVm(logger logx.Logger, vmName string) (string, error) {

	cli := "sudo rm -rf /tmp/*"
	_, err := run.RunCliSsh(vmName, cli)
	if err != nil {
		return "", err
	}
	return "", nil
}

// NAme: createSliceCleanTmp
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
func createSliceCleanTmp(logger logx.Logger, targets []phase.Target) []syncx.Func {
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
			if _, err := cleanTmpOnSingleVm(logger, vmCopy.Name()); err != nil {
				logger.Errorf("🅣 Failed to do action on VM %s: %v", vmCopy.Name(), err)
				return err
			}

			logger.Infof("🅣 VM %s cleaned /tmp successfully", vmCopy.Name())
			return nil
		})
	}

	return tasks
}

// name: CleanTmpOnVm
//
// description: the overall task.
//
// Notes:
// - Each target must implement the Target interface.
func CleanTmpOnVm(ctx context.Context, logger logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	appx := "CleanTmpOnVm"
	logger.Infof("🅣 Starting phase: %s", appx)
	// check paramaters
	if len(targets) == 0 {
		logger.Warnf("🅣 No targets provided to phase: %s", appx)
		return "", nil
	}

	// Build slice of functions
	tasks := createSliceCleanTmp(logger, targets)

	// Log number of tasks
	logger.Infof("🅣 Phase %s has %d concurent tasks", appx, len(tasks))

	// Run tasks in the slice concurrently
	if errs := syncx.RunConcurrently(ctx, tasks); errs != nil {
		return "", errs[0] // return first error encountered
	}

	return "", nil
}
