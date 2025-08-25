/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"context"
	"log"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/goluc/internal"
)

var myWorkflow *phase.Workflow

func init() {
	var err error
	myWorkflow, err = phase.NewWorkflowFromPhases(
		phase.NewPhase("check_status", "Checks if the system is ready", internal.CheckSystemStatus, nil),
		phase.NewPhase("fetch_data", "Fetches data from an external source", internal.FetchLatestData, nil),
		phase.NewPhase("process_data", "Processes the fetched data", internal.ProcessData, nil),
		phase.NewPhase("generate_report", "Generates a final report", internal.GenerateReport, nil),
	)
	if err != nil {
		log.Fatalf("failed to build workflow: %v", err)
	}
}

// func init() {
// 	var err error
// 	myWorkflow, err = phase.NewWorkflowFromPhases(
// 		phase.NewPhase("check_status", "Checks if the system is ready", internal.CheckSystemStatus, nil),
// 		phase.NewPhase("fetch_data", "Fetches data", internal.FetchLatestData, []string{"check_status"}),
// 		phase.NewPhase("process_data", "Processes data", internal.ProcessData, []string{"fetch_data"}),
// 		phase.NewPhase("generate_report", "Generates report", internal.GenerateReport, []string{"process_data"}),
// 	)
// 	if err != nil {
// 		log.Fatalf("failed to build workflow: %v", err)
// 	}
// }

// testPhase is the function that defines and runs the workflow.
func testPhase() {

	var targets = []phase.Target{
		&phase.Vm{NameStr: "o1u"},
		&phase.Vm{NameStr: "o2a"},
		&phase.Vm{NameStr: "o3r"},
		&phase.Vm{NameStr: "o4f"},
		&phase.Vm{NameStr: "o5d"},
	}

	// Show the sequence of phases before running the sequence.
	myWorkflow.Show(logx.GetLogger())

	// Create a context for the workflow
	ctx := context.Background()

	// Execute the workflow
	if err := myWorkflow.Execute(ctx, logx.GetLogger(), targets, nil, nil); err != nil {
		log.Fatalf("Workflow execution failed: %v", err)
	}
}
