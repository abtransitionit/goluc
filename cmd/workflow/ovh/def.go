/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ovh

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	coreurl "github.com/abtransitionit/gocore/url"
)

// OVH variables
var (
	createApp      = "https://api.ovh.com/createApp/"
	ovhGithubGoLib = "https://github.com/ovh/go-ovh"
	ovhApiDoc      = "https://help.ovhcloud.com/csm/fr-api-getting-started-ovhcloud-api?id=kb_article_view&sysparm_article=KB0042789"
	listPrestatire = "https://partner.ovhcloud.com/fr/directory/"
	ovhEndpoint01  = "https://api.ovh.com/ "
	ovhEndpoint02  = "https://eu.api.ovh.com/"
	ovhApiIhm      = "https://eu.api.ovh.com/console/"
)

// OVH custom
var (
	credentialFilePath = "~/wkspc/.config/ovh/credential"
)

// Package variables
var (
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
	targets []corephase.Target
)

// Package variables : confifg1
var (
	cmdName = "ovh"
	SDesc   = "This is the OVH workflow."
)

// Package variables : confifg2
var (
	// vmList = []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
	vmList = []string{"o1u"}
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
		corephase.NewPhase("listInfo", "List infos on OVH API(s)", listInfo, nil),
		// corephase.NewPhase("listApplication", "List application(s)", listInfo, nil),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}

func listInfo(ctx context.Context, logger logx.Logger, targets []corephase.Target, cmd ...string) (string, error) {
	logx.Infof("🔹 Ovh API Endpoint : %s", ovhEndpoint01)
	logx.Info(coreurl.Display("🔹 Docs", ovhApiDoc))
	logx.Info(coreurl.Display("🔹 Ovh Go lib on Github", ovhGithubGoLib))
	logx.Info(coreurl.Display("🔹 Liste ESN", listPrestatire))
	logx.Info(coreurl.Display("🔹 Explore OVH API(s)", ovhEndpoint02))
	logx.Info(coreurl.Display("🔹 IHM API", ovhApiIhm))
	logx.Infof("🔹 Credential file path: %s", credentialFilePath)
	logx.Infof("🔹 APi: create App via gui: %s", createApp)
	return "", nil
}
