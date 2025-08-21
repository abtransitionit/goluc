/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package gotc

import (
	"github.com/abtransitionit/gocore/logx"
	corephase "github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
)

// Package variables
var (
	SDesc   = "This is the GO toochain workflow."
	cmdName = "gotc"
	logger  = logx.GetLogger()
	wkf     *corephase.Workflow
)

func init() {
	var err error
	wkf, err = corephase.NewWorkflowFromPhases(
		corephase.NewPhase("show", "show execution plan.", internal.CheckSystemStatus, nil),
		corephase.NewPhase("checkP", "check prerequisites.", internal.CheckSystemStatus, nil),
		corephase.NewPhase("download", "download tarball.", internal.CheckSystemStatus, nil),
		corephase.NewPhase("copy", "temporary copy file to temp.", internal.FetchLatestData, nil),
		corephase.NewPhase("extract", "extract file to dest folder", internal.ProcessData, nil),
	)
	if err != nil {
		logger.ErrorWithStack(err, "failed to build workflow: %v")
	}
}
