/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package repo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/abtransitionit/gocore/list"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/golinux/da"
	"github.com/abtransitionit/golinux/property"
	"github.com/abtransitionit/goluc/internal"
	"github.com/spf13/cobra"
)

var (
	showInstalled bool
	showWhitelist bool
)

// Description
var listSDesc = "list native os package repositories installed or in the whitelist (ie. authorized)."
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

		// 1 - list installed repos
		// 11 - define list Vms
		vmNames := []string{"o1u", "o2a", "o3r", "o4f", "o5d"}

		// no action is needed based on the number of row
		formatedString := strings.Join(vmNames, "\n")
		rowCount := list.CountNbLine(formatedString)
		if rowCount == 1 {
			return
		}

		// 12 - print the list
		formatedString = "NAME\n" + strings.Join(vmNames, "\n")
		list.PrettyPrintTable(formatedString)

		// Ask user which ID (to choose) from the printed list
		fmt.Print("\nWhich item (enter ID): ")

		// convert user input to int
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		id, err := strconv.Atoi(input)
		if err != nil {
			logger.Errorf("invalid ID: %v", err)
			return
		}

		// define resource property from ID and output
		vmName, err := list.GetFieldByID(formatedString, id, 0)
		if err != nil {
			logger.Errorf("failed to get pod name from ID: %s: %v", id, err)
			return
		}

		// get property
		osFamily, err := property.GetProperty(vmName, "osfamily")
		if err != nil {
			logger.Errorf("%v", err)
			return
		}
		// get property
		osDistro, err := property.GetProperty(vmName, "osdistro")
		if err != nil {
			logger.Errorf("%v", err)
			return
		}

		logger.Debugf("family/distro: %s / %s", osFamily, osDistro)

		// 13 - get the  package manager
		repo := da.Repo{Name: "some-repo"} // TODO: from flag/args
		out, err := repo.GetCliBuilder(osFamily, osDistro)
		if err != nil {
			logger.Errorf("failed to get package manager: %v", err)
			return
		}
		logger.Errorf("%v", out)
		// switch osFamily {
		// case "debian":
		// 	repo.Cbd = &da.AptManager{Repo: &repo}
		// case "rhel", "fedora":
		// 	repo.Cbd = &da.DnfManager{Repo: &repo, Distro: osDistro}
		// default:
		// 	logger.Errorf("unsupported OS family/distro: %s", osFamily, osDistro)
		// 	return
		// }

		// // Get list repo
		// output, err := repo.List(context.Background(), false, vmName, logger)
		// if err != nil {
		// 	logger.Errorf("failed to list repos: %v", err)
		// 	return
		// }
		// fmt.Println(output)

	},
}

func init() {
	listCmd.Flags().BoolVarP(&showInstalled, "installed", "i", false, "show installed Helm repos (default)")
	listCmd.Flags().BoolVarP(&showWhitelist, "whitelist", "w", false, "show installable Helm repos (organization whitelist)")
}
