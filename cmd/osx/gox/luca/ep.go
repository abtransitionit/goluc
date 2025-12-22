/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package luca

import (
	"fmt"

	"github.com/abtransitionit/golinux/mock/run"
	"github.com/spf13/cobra"
)

var localFlag bool

// Description
var epSDesc = "build luca for the KBE project."
var epLDesc = epSDesc

// root Command
var EpCmd = &cobra.Command{
	Use:   "luca",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// build CLI
		cli := `GOOS=linux GOARCH=amd64 go build -C /Users/max/wkspc/git/goluc -ldflags="-s -w" -o /tmp/luca`
		fmt.Printf("building luca with CLI: %s\n", cli)
		_, err := run.RunCli("local", cli, nil)
		if err != nil {
			cmd.PrintErrln("building luca, err:", err)
			return
		}
		// handle success
		cmd.Println("luca built successfully")
	},
}

// func init() {
// 	// EpCmd.PersistentFlags().BoolVarP(&remoteFlag, "remote", "r", false, "uses by default the local Helm client unless the flag is provided (it will use the remote Helm client)")
// 	EpCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "Use the local Helm client if the flag is set; otherwise, use the remote Helm client")
// 	EpCmd.AddCommand(addCmd)
// 	EpCmd.AddCommand(listCmd)
// }
