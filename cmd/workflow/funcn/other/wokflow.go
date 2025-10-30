package main

import (
	"fmt"
)

// Description: Execute the function of the workflow
func (w *Workflow) Run() error {
	fmt.Printf("👉 Starting workflow: %s\n", w.Name)

	// loop over all phases of the workflow
	for name, phase := range w.Phases {
		fmt.Printf("▶️  Phase: %s (%s)\n", name, phase.Fn)
		// execute the phase:function
		if err := ExecFunction(phase.Fn); err != nil {
			// return the original error
			return fmt.Errorf("phase %q failed: %w", name, err)
		}
	}

	fmt.Printf("👉 completed successfully Workflow %s \n", w.Name)
	return nil
}
