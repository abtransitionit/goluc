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
var addSDesc = "add a helm repo."
var addLDesc = addSDesc + `
- This command add the helm repo by just updating the Helm client configuration file in the user's home directory.
- If the repository name is not already in the Helm configuration file, it adds it.
- If the repository name is already in the Helm configuration file, it updates the URL of the repository.
`

// root Command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: addSDesc,
	Long:  addLDesc,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("❌ you must pass exactly 1 arguments, the name of the repository (in the whitelist) to add, got %d", len(args))
		}
		return nil
	},
	Example: fmt.Sprintf(`
  # add helm repo from whitelist
  %[1]s repo add bitnami
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(addSDesc)
		ctx := context.Background()

		// get repo key from args
		repoKey := args[0]

		// check provided name exists in the whitelist and get url
		repoObj, ok := helm.MapHelmRepoReference[repoKey]
		if !ok {
			logger.Errorf("repository '%s' is not in the allowed helm repository whitelist", repoKey)
			return
		}

		// define var needed by cli
		repo := helm.HelmRepo{
			Name: repoObj.Name,
			Url:  repoObj.Url,
		}
		var err error

		// define cli
		cli, err := repo.Add(ctx, logger)
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
