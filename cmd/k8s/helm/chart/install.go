/*
Copyright © 2025 Amar BELGACEM abtransitionit@hotmail.com
*/
package chart

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

var installSDesc = "Install a [Helm] chart (ie. helm install --upgrade ...)"
var (
	chartVersionFlag string
	chartNameFlag    string
	k8sNamespaceFlag string
	chartReleaseFlag string
	repoNameFlag     string
	fileConfFlag     string
)

// Parent command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: installSDesc,
	Example: `
		install
			--repo		kdashb
			--name		kubernetes-dashboard
			--release	kbe-kdashb 
			--version	7.12
			--ns		kdashb

		install
			--repo		cilium
			--name		cilium
			--release	kbe-cilium
			--version	1.17
			--ns		kube-system
			--file		kbe-cilium.yaml

		install
			--repo		ingnginx
			--name		ingress-nginx
			--release	kbe-ingress-nginx
			--version	4.12
			--ns		ingnginx

	`,
	Run: func(cmd *cobra.Command, args []string) {
		// define ctx and logger
		logger := logx.GetLogger()
		logger.Info(installSDesc)
		cmd.Help()

		// fmt.Printf("🚀 %s\n", installSDesc)
		// // confirm action
		// if !forceFlag && !ui.AskForConfirmation() {
		// 	fmt.Fprintln(os.Stderr, "⚠️ Operation canceled")
		// } else {
		// 	installAction()
		// }
	},
}

func init() {
	// Flag - only for this command
	installCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass  confirmation and force installation")
	installCmd.Flags().StringVar(&chartVersionFlag, "version", "", "The version of the chart")
	installCmd.Flags().StringVar(&chartNameFlag, "name", "", "The [formal] name of the chart to install")
	installCmd.Flags().StringVar(&chartReleaseFlag, "release", "", "A [custom] release name for the chart being installed")
	installCmd.Flags().StringVar(&k8sNamespaceFlag, "namespace", "", "The k8s namespace where the chart must lives")
	installCmd.Flags().StringVar(&repoNameFlag, "repo", "", "The [custom] name of a Helm repo")
	installCmd.Flags().StringVar(&fileConfFlag, "file", "", "A configuration files")
}

// func installAction() {
// 	cli := ""
// 	if fileConfFlag != "" {
// 		cli = fmt.Sprintf(`helm upgrade --install %s %s/%s --version %s --namespace %s -f %s`, chartReleaseFlag, repoNameFlag, chartNameFlag, chartVersionFlag, k8sNamespaceFlag, fileConfFlag)
// 	} else {
// 		cli = fmt.Sprintf(`helm upgrade --install %s %s/%s --version %s --namespace %s`, chartReleaseFlag, repoNameFlag, chartNameFlag, chartVersionFlag, k8sNamespaceFlag)
// 	}
// 	// Play CLI
// 	res, cerr, err := config.PlayQueryHelm(cli)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, cerr)
// 		os.Exit(1)
// 	}
// 	fmt.Printf(res)
// }

// //Todo : handle errors
