package main

import (
	"fmt"

	"github.com/test/gocore/callfn"
)

// description register in the registry all functions of the workflow
func RegisterWkfFunction() {

	RegisterSingleFunc("vm.CheckVmSshAccess", callfn.CheckVmSshAccess)
	RegisterSingleFunc("luc.DeployLuc", callfn.DeployLuc)
	RegisterSingleFunc("dnfapt.UpgradeVmOs", callfn.UpgradeVmOs)
	RegisterSingleFunc("dnfapt.UpdateVmOsApp", callfn.UpdateVmOsApp)

	// log
	fmt.Println("completed Registration succesfully")
}

// Description: a private registry
//
// Notes:
// - map a YAML function denote as a string to a real Go functions.
// - private and not exported.
var registry = make(map[string]YamlFunc)

func RegisterSingleFunc(name string, f YamlFunc) {
	if _, exists := registry[name]; exists {
		fmt.Printf("⚠️  overwriting registration for %q\n", name)
	}
	registry[name] = f
}
