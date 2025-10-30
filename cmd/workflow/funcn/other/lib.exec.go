package main

import (
	"fmt"
)

func ExecFunction(name string) error {
	if targetFunc, ok := registry[name]; ok {
		fmt.Printf("✅ manager: Start Playing function: %s\n", name)
		targetFunc() // Execute the function
		fmt.Printf("✅ manager: Played function: %s\n", name)
		return nil
	}
	return fmt.Errorf("unknown function name %q in registry", name)
}
