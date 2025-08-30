/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package do

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	coreurl "github.com/abtransitionit/gocore/url"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags.

var downloadSDesc = "Download a file denoted by an URL"
var downloadLDesc = downloadSDesc + "\n" + `
Example usage.

go run . do download <URL>
`

// the command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: downloadSDesc,
	Long:  downloadLDesc,
	Args:  cobra.ExactArgs(1), // Require exactly one argument: the url
	RunE: func(cmd *cobra.Command, args []string) error {
		url := args[0]
		// Create a context to allow user to ctrl-c the download.
		ctx := context.Background()
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()
		// logger := logx.GetLogger()
		// logger.Infof("%s", cmd.Short)
		fileName, err := coreurl.DownloadArtifact(ctx, vmName, url, prefix)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		fmt.Println(fileName)
		return nil
	},
}

func init() {
	// downloadCmd.Flags().StringVarP(&vmName, "vm", "v", "", "VM name on which to do the download")
	downloadCmd.Flags().StringVarP(&prefix, "prefix", "p", "tmp", "a custom prefix for the downloaded artifact's path")

}
