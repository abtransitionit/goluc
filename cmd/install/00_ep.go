/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package install

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/goluc/cmd/install/gotc"
	"github.com/abtransitionit/goluc/cmd/install/kbe"
	"github.com/abtransitionit/goluc/cmd/install/kind"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

// Description

var installSDesc = "install softwares, tools locally or on remote host(s)."

var installLDesc = installSDesc + "\n" + `
This command allows you to install various tools or software either on your local
machine or on one or more remote hosts/VMs/containers.

The installation process is usually driven by a custom YAML configuration file.
The structure of this YAML file depends on the tool/software being installed,
and may define:
  - Target hosts or VMs
  - Installation parameters or options
  - Environment-specific settings
`

// root Command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: installSDesc,
	Long:  installLDesc,
	Example: fmt.Sprintf(`
  # install kind locally using default configuration
  %[1]s  install kind <tool>
	
  # Install kind locally using default configuration
  %[1]s  install kind --remote o1u <tool>

  # provision a Kubernetes cluster on remote VMs using a default configuration
  %[1]s  install kbe <tool>

  # provision a Kubernetes cluster on remote VMs using a specific configuration
  %[1]s  install kbe -f config.yaml <tool>
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info("%s", installSDesc)
		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}

func init() {
	InstallCmd.AddCommand(gotc.GotcCmd)
	InstallCmd.AddCommand(kbe.KbeCmd)
	InstallCmd.AddCommand(kind.KindCmd)
}
