/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/gocore/logx"
	"github.com/spf13/cobra"
)

// Description
var playSDesc = "Test some code."
var playLDesc = playSDesc + ` xxx.`

// root Command
var playCmd = &cobra.Command{
	Use:   "play",
	Short: playSDesc,
	Long:  playLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		logx.Info(playSDesc)

		// logx.Info("test Phase testPhaseO1")
		// testPhaseO1()

		logx.Info("test Phase testPhase")
		testPhase()

		// logx.Info("test DeleteFile")
		// testDeleteFile()

		// logx.Info("test CanBeSudoAndIsNotRoot")
		// testCanBeSudoAndIsNotRoot()

		// logx.Info("test TouchAsSudo")
		// testTouchAsSudo()

		logx.Info("End of test")

		// test.CheckCliExits("gpg")
		// // out := test.GeRemoteProperty("o1u", "osfamily")
		// out, err := test.GetPackage1("o1u", "gpg")
		// if err != nil {
		// 	logx.L.Error("%v : %s", err, out)
		// 	return
		// }
		// logx.L.Infof("yo in test : %s", out)
		// local function tested
		// test.TouchFileLocal("/tmp", "titi")
		// test.CheckFileLocalExits("/tmp/test.txt")
		// remote function tested
		// test.TouchFileOnRemote("o1u", "/tmp", "toto")
		// test.CheckFileRemoteExists("o1u", "/tmp/toto")

		// test.TestCheckCliExistsOnremote(config.KbeListNode, "gpg")
		// test.TestCheckCliExistsOnremote(config.KbeListNode, "curl")

		// url := "https://pkgs.k8s.io/core:/stable:/v1.32/rpm/repodata/repomd.xml.key"
		// path := "/etc/apt/sources.list.d/kbe-k8s-apt-keyring.gpg"
		// vm := "o1u"
		// test.TestRemoteGetGpgFromUrl(vm, url, path, true)
		// test.TestGetGpgFromUrl(url, path, true)

		// test.TestVmAreSshReachable(config.KbeListNode)
		// test.DaAddRepoLocal("kbe-k8s")

		// createFileLocal()
		// touchFileRemote("o1u")
		// MoveFileLocal()
		// ListOvhVm()
		// ListMapKey()
		// installGoCli()
		// getPath()
		// fmt.Println(configi.KbeGoCliConfigMap)
		// addLineToFileRemote()

	},
}

var forceFlag bool

func init() {
	playCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	playCmd.Flags().BoolP("list", "l", false, "List all available phases")
	playCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	playCmd.MarkFlagsMutuallyExclusive("list", "runall")
}
