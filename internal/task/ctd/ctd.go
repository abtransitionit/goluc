// file gotask/gocli/icli.go
package ctd

import (
	"context"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/gocore/syncx"
)

func SetContainerdOnSingleVm(ctx context.Context, logger logx.Logger, vmName string) (string, error) {
	cmds := []string{
		"echo CNI_PATH=$CNI_PATH",
		"echo PWD=$(pwd)",
		"containerd-rootless-setuptool.sh install",
	}
	cmd := strings.Join(cmds, " && ")
	logger.Debugf("playing %s: %s", vmName, cmd)

	// play cli
	output, err := run.RunCliSsh(vmName, cmd)
	if err != nil {
		return "", fmt.Errorf("failed to play cli on vm: '%s': '%s' : %w", vmName, cmd, err)
	}

	// success
	fmt.Printf("%s\n", output)

	logger.Debugf("%s: played %s", vmName, cmd)
	return "", nil
}

func createSliceFuncForContainerd(ctx context.Context, logger logx.Logger, targets []phase.Target) []syncx.Func {
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
		// define the job of the task and add it to the slice
		tasks = append(tasks, func() error {
			if _, err := SetContainerdOnSingleVm(ctx, logger, vmCopy.Name()); err != nil {
				logger.Errorf("🅣 Failed to execute task on VM %s: %v", vmCopy.Name(), err)
				return err
			}
			logger.Infof("🅣 task on VM %s succeded", vmCopy.Name())
			return nil
		})
	}

	return tasks
}

func SetContainerd(ctx context.Context, logger logx.Logger, targets []phase.Target, cmd ...string) (string, error) {
	appx := "SetContainerd"
	logger.Infof("🅣 Starting phase: %s", appx)
	// check paramaters
	if len(targets) == 0 {
		logger.Warnf("🅣 No targets provided to phase: %s", appx)
		return "", nil
	}

	// Build slice of functions
	tasks := createSliceFuncForContainerd(ctx, logger, targets)

	// Log number of tasks
	logger.Infof("🅣 Phase : %s : has %d concurent tasks", appx, len(tasks))

	// Run tasks in the slice concurrently
	if errs := syncx.RunConcurrently(ctx, tasks); errs != nil {
		return "", errs[0] // return first error encountered
	}

	return "", nil
}
