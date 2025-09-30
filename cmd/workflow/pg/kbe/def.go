/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kbe

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
)

// Package variables
var (
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
	targets []corephase.Target
)

// Package variables : confifg1
var (
	cmdName = "kbe"
	SDesc   = "KBE production grade and security workflow."
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
		corephase.NewPhase("firewall", "allows only allowed ports between nodes using nftables firewall", Todo, nil),
		corephase.NewPhase("firewall", "deny all direct access to nodes via ports", Todo, nil),
		corephase.NewPhase("firewall", "deny all direct access to nodes", Todo, nil),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}

func Todo(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {
	logger.Infof("To implement")
	return "", nil
}
