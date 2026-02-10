/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package lucm

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/mock/run"
	"github.com/spf13/cobra"
)

var localFlag bool

// Description
var epSDesc = "build luc for mac (aka lucm)."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "lucm",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		// var
		OutFilePath := "/tmp/lucm"
		// cli
		cli := fmt.Sprintf(`GOOS=darwin GOARCH=arm64 go build -C /Users/max/wkspc/git/goluc -ldflags="-s -w" -o %s`, OutFilePath)
		// log
		logger.Info("building luca - luc for the KBE project for mac worstation")
		logger.Debugf("CLI: %s\n", cli)
		// build CLI
		_, err := run.RunCli("local", cli, nil)
		if err != nil {
			cmd.PrintErrln("building luca, err:", err)
			return
		}
		// handle success
		cmd.Println("luca built successfully")
	},
}
