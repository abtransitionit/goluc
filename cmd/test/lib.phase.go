/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gotest/internal"
)

var mySequence = phase.NewPhaseList(
	phase.SetPhase("Setup", internal.SetupFunc, "Prepares the environment for the build."),
	phase.SetPhase("Build", internal.BuildFunc, "Compiles the source code into a binary."),
	phase.SetPhase("Test", internal.TestFunc, "Executes unit and integration tests."),
)

func testPhase() {
	// Get the global logger instance.
	log := logx.GetLogger()

	// Show the sequence of phases before running the sequence.
	mySequence.Show(log)

	// Run the sequence.
	if err := mySequence.Run(log); err != nil {
		log.ErrorWithNoStack(err, "Workflow execution failed.")
		return
	}
}
