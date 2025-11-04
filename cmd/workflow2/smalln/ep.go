/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package smalln

import (
	"context"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/onpm"
	common "github.com/abtransitionit/goluc/cmd/workflow/_common"
)

// root Command
var EpCmd = common.GetEpCmd(
	cmdName,
	shortDesc,
)

func init() {
	// sub cde
	EpCmd.AddCommand(common.GetPrintCmd(cmdName))
	EpCmd.AddCommand(common.GetRunCmd(cmdName))

	// function mapping
	common.FunctionRegistry.Add(
		"UpgradePkg",
		func(ctx context.Context, params any, logger logx.Logger) error {
			out, err := onpm.UpgradePkg()
			logger.Infof("Result: %s", out)
			return err
		})

	common.FunctionRegistry.Add(
		"UpgradeOs",
		func(ctx context.Context, params any, logger logx.Logger) error {
			out := onpm.UpgradeOs()
			logger.Infof("Result: %s", out)
			return nil
		})
	common.FunctionRegistry.Add(
		"InstallRepo",
		func(ctx context.Context, params any, logger logx.Logger) error {
			p := params.(onpm.ParamInstallRepo)
			out, err := onpm.InstallRepo(p.NodeName, p.RepoList)
			logger.Infof("Result: %s", out)
			return err
		})
	common.FunctionRegistry.Add(
		"LoadKModule",
		func(ctx context.Context, params any, logger logx.Logger) error {
			p := params.(onpm.ParamLoadKModule)
			out, err := onpm.LoadKModule(p.KModuleList, p.KFilePath)
			logger.Infof("Result: %s", out)
			return err
		})

	common.FunctionRegistry.Add(
		"SetPath",
		func(ctx context.Context, params any, logger logx.Logger) error {
			p := params.(onpm.ParamSetPath)
			return onpm.SetPath(p.BasePath, p.CustomRcFilename)
		})

	common.FunctionRegistry.Add(
		"CheckSshAccess",
		func(ctx context.Context, params any, logger logx.Logger) error {
			p := params.(onpm.ParamCheckSshAccess)
			return onpm.CheckSshAccess(p.NodeName)
		})
}

// // Test calling functions
// 	registry.Get("InstallRepo").Func(ctx, InstallRepoParams{
// 		NodeName: "node01",
// 		ListRepo: ListRepo{{Name: "repo1"}, {Name: "repo2"}},
// 	}, logger)

// 	registry.Get("SetPath").Func(ctx, SetPathParams{
// 		BinFolderPath:    "/usr/local/bin",
// 		CustomRcFileName: ".profile.luc",
// 	}, logger)
