package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
)

func getWorkflow() (*Workflow, error) {

	// 1. Define YAML workflow file path
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		return nil, fmt.Errorf("could not get caller information")
	}
	workflowPath := filepath.Join(path.Dir(file), "phase.yaml")

	// 2. Load the yaml using the generic function from lib.go
	workflow, err := LoadFile[Workflow](workflowPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load workflow from %s: %w", workflowPath, err)
	}

	return workflow, nil
}
