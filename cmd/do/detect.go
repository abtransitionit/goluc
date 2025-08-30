/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package do

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/abtransitionit/gocore/filex"
	"github.com/spf13/cobra"
)

// Package variables for CLI flags.

var detectSDesc = "guess to detect the binary type of a file"
var detectLDesc = detectSDesc + "\n" + `
Example usage.

go run . do download <URL>
`

// the command
var detectCmd = &cobra.Command{
	Use:   "detect",
	Short: detectSDesc,
	Long:  detectLDesc,
	Args:  cobra.ExactArgs(1), // Require exactly one argument: the file location
	RunE: func(cmd *cobra.Command, args []string) error {
		filePath := args[0]
		// Create a context to allow user to ctrl-c the download.
		ctx := context.Background()
		ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
		defer cancel()
		// logger := logx.GetLogger()
		// logger.Infof("%s", cmd.Short)
		fileType, err := filex.DetectBinaryType(ctx, filePath)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		fmt.Println(fileType)
		return nil
	},
}
