/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package workflow

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

var worflowSDesc = "manage the provisioning process of systems, softwares and tools."

var worflowLDesc = worflowSDesc + "\n" + `
This command allows you to install various systems, software or tools either on your local
machine or on one or more remote hosts/VMs/containers.

The installation process is driven by a worflow and a custom configuration YAML file.
The structure of this YAML file depends on the system/software/tool being installed,
and may define:
  - Target hosts, VMs or containers
  - Installation parameters or options
  - Other Environment-specific settings
`

// root Command
var WorkflowCmd = &cobra.Command{
	Use:   "install",
	Short: worflowSDesc,
	Long:  worflowLDesc,
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
		logx.Infof("%s", worflowSDesc)
		// manage argument
		if len(args) == 0 {
			cmd.Help()
			return
		}

	},
}

func init() {
	WorkflowCmd.AddCommand(gotc.GotcCmd)
	WorkflowCmd.AddCommand(kbe.KbeCmd)
	WorkflowCmd.AddCommand(kind.KindCmd)
}
