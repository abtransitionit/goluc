/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"context"
	"fmt"

	helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description
var deleteSDesc = "delete a repo."
var deleteLDesc = deleteSDesc + `
- This command delete the helm repo by just updating the Helm client configuration file in the user's home directory.
`

// root Command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: deleteSDesc,
	Long:  deleteLDesc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("❌ you must pass exactly 1 argument, the name of the repository")
		}
		return nil
	},
	Example: fmt.Sprintf(`
  # add helm repo
  %[1]s repo add myrepo https://charts.bitnami.com/bitnami	
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(deleteSDesc)
		ctx := context.Background()

		// define var needed by cli
		repo := helm.HelmRepo{
			Name: args[0],
		}
		var err error

		// define cli
		cli, err := repo.Delete(ctx, logger)
		if err != nil {
			logger.Errorf("failed to build helm command: %v", err)
			return
		}

		// run cli on local or remote
		var output string
		if localFlag {
			logger.Debugf("running on local helm client: %s", cli)
			output, err = helm.QueryHelm("", cli, logger)
		} else {
			remoteHelmHost := "o1u"
			logger.Debugf("running on remote helm client: %s : %s", remoteHelmHost, cli)
			output, err = helm.QueryHelm(remoteHelmHost, cli, logger)
		}

		if err != nil {
			logger.Errorf("failed to run helm command: %s: %w", cli, err)
			return
		}

		fmt.Println(output)
	},
}
