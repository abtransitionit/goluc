/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pod

import (
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	kubectl "github.com/abtransitionit/gocore/k8s-kubectl"
	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/gocore/ui"
	"github.com/spf13/cobra"
)

// Description
var describeSDesc = "list all pods."
var describeLDesc = describeSDesc

// root Command
var DescribeCmd = &cobra.Command{
	Use:   "desc",
	Short: describeSDesc,
	Long:  describeLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()

		// get list
		output, err := kubectl.ListPod(localFlag, "o1u", logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// output
		list.PrettyPrintTable(output)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose pod (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		podName, err := list.GetFieldByID(output, id, 1)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}
		podNs, err := list.GetFieldByID(output, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod namespace from ID: %s: %v", id, err)
			return
		}

		// define cli
		cli, err := kubectl.Resource{Type: "pod", Name: podName, Ns: podNs}.Describe()
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// play cli
		output, err = run.ExecuteCliQuery(cli, logger, localFlag, "o1u", helm.HandleHelmError)
		if err != nil {
			logger.Errorf("failed to run command: %s: %w", cli, err)
			return
		}

		fmt.Println(output)

	},
}

func init() {
	DescribeCmd.PersistentFlags().BoolVarP(&localFlag, "local", "l", false, "uses by default the remote Helm client unless the flag is provided (it will use the local Helm client)")
}
