/*
Copyright © 2025 Amar BELGACEM abtransitionit@hotmail.com
*/
package chart

import (

	// "github.com/abelgacem/lucg/config"
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var chartValueShortDesc = "Display possible values on a [Helm] chart of a [Helm chart] repository"

// Parent command
var valueCmd = &cobra.Command{
	Use:   "value ChartName",
	Short: chartValueShortDesc,
	Example: `
	value --ingnginx ingress-nginx
	value --cilium    cilium
	value ingress-nginx --ingngin
	value cilium        --cilium
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(listSDesc)
		cmd.Help()

		// fmt.Println("\n🟦", chartValueShortDesc)
		// // action : count the number of flag passed
		// nbFlag := 0
		// cmd.Flags().Visit(func(*pflag.Flag) { nbFlag++ })
		// // error : exit : the number of flag(Helm Repo) must be exactly 1
		// if nbFlag != 1 {
		// 	fmt.Fprintln(os.Stderr, "❌ Error: you must specify a flag (ie. a Helm chart repository).")
		// 	return
		// }
		// // error : exit : the number of args must be 1
		// if len(args) != 1 {
		// 	fmt.Fprintln(os.Stderr, "❌ Error: you must specify a chart name.")
		// 	return
		// }
		// helmChartName := args[0]
		// helmRepoName := ""
		// // map the flag to the Helm repo
		// // 👉 : *repo.FlagVar denote the user input boolean:flag
		// for _, repo := range config.ListHelmRepo {
		// 	if repo.FlagVar != nil && *repo.FlagVar {
		// 		helmRepoName = repo.Name
		// 	}
		// }
		// cli := fmt.Sprintf(`helm show values %s/%s`, helmRepoName, helmChartName)
		// output, cerr, err := config.PlayQueryHelm(cli)
		// if err != nil {
		// 	fmt.Fprintln(os.Stderr, cerr)
		// }
		// fmt.Println(output)
	},
}

// func init() {
// 	for _, repo := range config.ListHelmRepo {
// 		valueCmd.Flags().BoolVar(repo.FlagVar, repo.Name, false, repo.Description)
// 	}
// }
