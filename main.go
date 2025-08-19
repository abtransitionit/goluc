/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package main

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd"
)

func main() {
	// init: make logx GALI (Global Application Logger Instance).
	logx.Init()
	cmd.Execute()
}
