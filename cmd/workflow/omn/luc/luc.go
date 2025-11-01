/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package luc

import (
	"context"
	"fmt"
	"strings"

	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/golinux/filex"
	"github.com/abtransitionit/goluc/internal/task/luc"
)

// Package variables
var (
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
	targets []corephase.Target
)

// Package variables : confifg2
var (
	vmList = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	// vmList = []string{"o1u", "o2a"}
)

func init() {
	// create the targets slice from vmList
	for _, vmName := range vmList {
		targets = append(targets, &corephase.Vm{NameStr: vmName})
	}

	// create the workflow
	var err error
	wkf, err = corephase.NewWorkflowFromPhases(
		// corephase.NewPhase("cleanuserbin", "empty the /usr/local/bin folder on the VMs of a list LUC specifics binaries", CleanUserBinForLucOnVm, nil),
		corephase.NewPhase("cleanTmp", "empty the /tmp folder on the VMs", luc.CleanTmpOnVm, nil),
		corephase.NewPhase("gitDevToMain", "merge dev to main and push all 4 projects : gocore, golinux, gotask, goluc", gitDevToMain, nil),
		corephase.NewPhase("gitDevToPremain", "merge dev to premain and push all 4 projects : gocore, golinux, gotask, goluc", gitTodo, nil),
		corephase.NewPhase("gitPremainToMain", "merge premain to main and push all 4 projects : gocore, golinux, gotask, goluc", gitTodo, nil),
		corephase.NewPhase("buildLuc", "build LUC for curent and linux platform", build, nil),
		corephase.NewPhase("deployLuc", "deploy LUC on all VMs", luc.DeployOnVm, nil),
		corephase.NewPhase("deployLucLocal", "deploy LUC Localy (ie. /usr/local/bin)", DeployLucLocally, []string{"buildLuc"}),
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
func gitTodo(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {
	logger.Infof("To implement")
	return "", nil
}
func DeployLucLocally(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {
	logger.Infof("Deploying goluc to /usr/local/bin")

	// get cli
	cli, err := filex.CpAsSudo(ctx, logger, "/tmp/goluc", "/usr/local/bin/goluc")
	if err != nil {
		return "", err
	}

	// play it localy
	_, err = run.RunOnLocal(cli)
	if err != nil {
		return "", err
	}

	// success
	logger.Infof("deployed goluc locally to /usr/local/bin")
	return "", nil
}
func gitDevToMain(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {

	// define var
	const repoFolder = "/Users/max/wkspc/git"
	var cmds []string
	// loopt over repo name
	for _, repoName := range []string{"gocore", "golinux", "gotask", "goluc"} {

		// define repo path
		repoPath := fmt.Sprintf("%s/%s", repoFolder, repoName)

		// define cli(s) to play on repo
		cmds = append(cmds, fmt.Sprintf("cd %s", repoPath))
		// cmds = append(cmds, "git checkout dev") // we sure we are on dev
		cmds = append(cmds, "git add .")                                           // stage all files
		cmds = append(cmds, "git diff --cached --quiet || git commit -m 'update'") // commit if there is changes
		cmds = append(cmds, "git push origin dev")                                 // push
		cmds = append(cmds, "git checkout main")
		cmds = append(cmds, "git merge --no-edit dev")
		cmds = append(cmds, "git push origin main")
		cmds = append(cmds, "git checkout dev")

		// Join commands with && to run them sequentially
		cli := strings.Join(cmds, " && ")
		_, err := run.RunOnLocal(cli)
		if err != nil {
			return "error while updating repo", err
		}
		logger.Debugf("updated repo %s", repoName)
	}
	return "repo updated", nil
}
