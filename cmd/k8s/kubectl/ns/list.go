/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package ns

import (
	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var listSDesc = "list all cluster namespaces."
var listLDesc = listSDesc

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get list
		output, err := kubectl.ListNs(localFlag, "o1u", logger)
		// cli, err := kubectl.Resource{Type: "ns"}.List()
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}
		// print list
		list.PrettyPrintTable(output)
	},
}

func init() {
	listCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "uses by default the remote Helm client unless the flag is provided (it will use the local Helm client)")
}
