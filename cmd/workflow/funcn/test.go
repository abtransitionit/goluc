/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package funcn

import (
	"fmt"

	"github.com/spf13/cobra"
)

var myKey = "MyFunc"

func MyFunc(p1 string, p2 int) {
	fmt.Println("MyFunc called with", p1, "and", p2)
}

var registryFn = map[string]func(string, int){
	"MyFunc": MyFunc}

// root Command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test code",
	RunE: func(cmd *cobra.Command, args []string) error {

		// get function
		fn, ok := registryFn[myKey]
		if !ok {
			return fmt.Errorf("function %q not found", myKey)
		}

		fn("Hello", 42)
		return nil

	},
}

func executeFn() {
	fmt.Println("Execute function")
}
func resolveFn() {
	fmt.Println("Resolve function")
}
