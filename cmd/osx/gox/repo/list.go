/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/abtransitionit/gocore/ui"
	"github.com/abtransitionit/golinux/da"
	lproperty "github.com/abtransitionit/golinux/mock/property"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

var (
	showInstalled bool
	showWhitelist bool
)

// Description
var listSDesc = "list native os package repositories installed (default) or in the whitelist (ie. authorized to install)."
var listLDesc = listSDesc

// root Command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: listSDesc,
	Long:  listLDesc,
	Example: fmt.Sprintf(`
# list installed repos
  %[1]s repo list -i

  # list installable repos (whitelist)
  %[1]s repo list -w
  `, internal.CliName),
	Run: func(cmd *cobra.Command, args []string) {

		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listSDesc)
		// ctx := context.Background()

		// default behavior: show installed repo
		if !showInstalled && !showWhitelist {
			showInstalled = true
		}

		// 1 - list whitelist repos
		if showWhitelist {
			logger.Info("Installable repositories (ie. organization whitelist):")

			// print the list
			rawAndtruncatedString := da.MapRepoReference.ConvertToStringTruncated()
			list.PrettyPrintTable(rawAndtruncatedString)
			return
		}

		// 1 - get list of OVH VPS names
		// vpsSliceName, := []string{"o1u", "o2a", "o3r", "o4f", "o5d"}
		vpsSliceName, err := ovh.GetVpsListName()
		if err != nil {
			logger.Errorf("getting vps:list from configuration file: %v", err)
			os.Exit(1)
		}

		// print list
		pList := "Name\n" + strings.Join(vpsSliceName, "\n")
		list.PrettyPrintTable(pList)

		// Ask user which ID (to choose) from the printed list
		id, err := ui.AskUserInt("\nchoose item (enter ID): ")
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// 2 - define resource property from user choice (ID and output)
		vmName, err := list.GetFieldByID(pList, id, 0)
		if err != nil {
			logger.Errorf("failed to get item from ID: %d > %v", id, err)
			return
		}
		osFamily, err := lproperty.GetProperty(logger, vmName, "osFamily")
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		osDistro, err := lproperty.GetProperty(logger, vmName, "osDistro")
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

		// get the  package manager
		repo := &da.Repo{Name: "some-repo"} // TODO: from flag/args
		err = repo.GetCliBuilder(osFamily, osDistro)
		if err != nil {
			fmt.Printf("Failed to get CLI builder: %v\n", err)
			return
		}

		// Get list repo
		output, err := repo.List(context.Background(), false, vmName, logger)
		if err != nil {
			logger.Errorf("failed to list repos: %v", err)
			return
		}
		fmt.Println(output)

	},
}

func init() {
	listCmd.Flags().BoolVarP(&showInstalled, "installed", "i", false, "show installed Helm repos (default)")
	listCmd.Flags().BoolVarP(&showWhitelist, "whitelist", "w", false, "show installable Helm repos (organization whitelist)")
}
