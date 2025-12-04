/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package pm

import (
	"github.com/abtransitionit/goluc/cmd/osx/pm/pkg"
	"github.com/abtransitionit/goluc/cmd/osx/pm/repo"
	"github.com/spf13/cobra"
)

// Description
var epSDesc = "manage linux os packages repositories and packages using native package managers."
var epLDesc = epSDesc + `
This command supports the following Linux families
  - Debian-based (Ubuntu, Debian)
  - Rhel-based (RHEL, CentOS, AlmaLinux, RockyLinux, Fedora). 

It abstracts the native package managers (apt, dnf) and provide a unified
way to manage repositories and packages.`

// root Command
var EpCmd = &cobra.Command{
	Use:   "pm",
	Short: epSDesc,
	Long:  epLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	EpCmd.AddCommand(repo.EpCmd)
	EpCmd.AddCommand(pkg.EpCmd)
}
