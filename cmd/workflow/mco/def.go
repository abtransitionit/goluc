/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package mco

import (
	"context"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/goluc/internal/task/luc"
)

// Package variables
var (
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
	targets []corephase.Target
)

// Package variables : confifg1
var (
	cmdName = "mco"
	SDesc   = "This is the LUC maintenace workflow."
)

// Package variables : confifg2
var (
	vmList = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
)

func init() {
	// create the targets slice from vmList
	for _, vmName := range vmList {
		targets = append(targets, &corephase.Vm{NameStr: vmName})
	}

	// create the workflow
	var err error
	wkf, err = corephase.NewWorkflowFromPhases(
		corephase.NewPhase("git", "merge dev to main and push all 4 project", git, nil),
		corephase.NewPhase("buildLuc", "build LUC for curent and linux platform", build, nil),
		corephase.NewPhase("deployLuc", "deploy LUC on all VMs", luc.DeployOnVm, nil),
		corephase.NewPhase("deleteLuc", "delete LUC on all VMs", luc.DeleteOnVm, nil),
		corephase.NewPhase("deployLuc2", "deploy LUC on all VMs", luc.DeployOnVm, []string{"buildLuc"}),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}

func build(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {
	cli := "goluc goprj build -o /tmp -p /Users/max/wkspc/git/goluc/"
	_, err := run.RunOnLocal(cli)
	if err != nil {
		return "", err
	}
	return "", nil
}
func git(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {

	// define var
	const repoFolder = "/Users/max/wkspc/git"
	var cmds []string
	// loopt over repo name
	for _, repo := range []string{"gocore", "golinux", "gotask", "goluc"} {

		// define repo path
		repoPath := fmt.Sprintf("%s/%s", repoFolder, repo)

		// define cli(s) to play on repo
		cmds = append(cmds, fmt.Sprintf("cd %s", repoPath))
		cmds = append(cmds, "git checkout main")
		cmds = append(cmds, "git merge --no-edit dev")
		cmds = append(cmds, "git push origin main")
		cmds = append(cmds, "git checkout dev")

		// Join commands with && to run them sequentially
		cli := strings.Join(cmds, " && ")
		_, err := run.RunOnLocal(cli)
		if err != nil {
			return "", err
		}
	}
	return "", nil
}
