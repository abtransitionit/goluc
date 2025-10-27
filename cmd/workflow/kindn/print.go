/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package kindn

import (
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/abtransitionit/gocore/phase"
	"github.com/abtransitionit/gocore/yamlx"
	"github.com/spf13/cobra"
)

// root Command
var printCmd = &cobra.Command{
	Use:   "print",
	Short: "display all the phases of the worflow",
	RunE: func(cmd *cobra.Command, args []string) error {

		// define yaml workflow file
		_, file, _, _ := runtime.Caller(0) // because it is not called directly but through GetSection
		workflowPath := filepath.Join(path.Dir(file), "phase.yaml")

		// Load the yaml
		workflow, err := yamlx.LoadFile[phase.Workflow2](workflowPath)
		if err != nil {
			log.Fatal(err)
		}

		// print it
		workflow.Print()
		return nil

	},
}
