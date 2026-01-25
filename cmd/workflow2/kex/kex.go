/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kex

import (
	"path/filepath"

	"github.com/abtransitionit/gocore/phase2"
	"github.com/abtransitionit/gotask/mock/k8s"
)

// register function used in the workflow
var (
	cmdPathName = "kex"
	shortDesc   = "Add Kubernetes Extension."
)

// registred function
func registerFunctions() {
	// get objects
	registry := phase2.GetFnRegistry()        // an instance of the FnRegistry that is shared by all phases
	worflowName := filepath.Base(cmdPathName) // the workflow name

	registry.Add(worflowName, "ApplyManifest", k8s.ApplyManifest)
	registry.Add(worflowName, "CreateSecret", k8s.CreateSecret)
}
