package main

import (
	"fmt"
	"os"
)

func main() {
	// register all function of the workflow
	RegisterWkfFunction()

	// get the workflow
	workflow, err := getWorkflow()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// run the workflow
	if err := workflow.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
