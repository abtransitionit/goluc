/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package vps

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
	"github.com/abtransitionit/gocore/syncx"
	"github.com/spf13/cobra"
)

// Description
var installSDesc = "Api re-install the an OVH VPS Os image."
var installLDesc = installSDesc + ` xxx.`

// root Command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: installSDesc,
	Long:  installLDesc,
	Run: func(cmd *cobra.Command, args []string) {
		// define the logger
		logger := logx.GetLogger()

		// log
		logger.Info(installSDesc)

		// define a ctx with timeout
		ctx := context.Background()

		// define the vps
		vpsNameId := "vps-9c33782a.vps.ovh.net"
		jsonx.PrettyPrintColor(vpsNameId)

		// // readyness check - define a specific timeout for waiting VPS
		// waitTimeout := 3 * time.Second
		// waitCtx, cancel := context.WithTimeout(ctx, waitTimeout)
		// defer cancel() // always defer cancel to release resources
		// start := time.Now()
		// healthFunc := ovh.CheckVpsIsReady(start)

		// fmt.Println("Waiting for resource to be ready...")
		// err := syncx.WaitForReady(waitCtx, logger, 2*time.Second, healthFunc)
		// if err != nil {
		// 	fmt.Printf("Resource not ready: %v\n", err)
		// 	return
		// }

		// fmt.Println("Resource is ready! ✅")

		// os.Exit(0)

		// 1 - api display state before reinstall
		vpsDetail, err := ovh.VpsGetDetail(ctx, logger, vpsNameId)
		if err != nil {
			logger.Errorf("get detail failed: %v", err)
			os.Exit(1)
		}
		jsonx.PrettyPrintColor(vpsDetail["state"])

		// 3 - api reinstall vps
		vpsInfo, err := ovh.VpsReinstallHandler(ctx, logger, vpsNameId)
		if err != nil {
			logger.Errorf("reinstall failed: %v", err)
			os.Exit(1)
		}
		// this display is sent to the user by the api before the reinstall - and ensure the reinstall will be done
		jsonx.PrettyPrintColor(vpsInfo)

		// 4 - wait for VPS readyness
		// 41 - time after which the readyness check will fail - nothing to do with the reinstall
		waitTimeout := 180 * time.Second
		waitCtx, cancel := context.WithTimeout(ctx, waitTimeout)
		defer cancel() // always defer cancel to release resources
		// 42 - thefunction thet test if the VPS is ready
		healthFunc := func() (bool, error) {
			return ovh.CheckVpsIsReady(waitCtx, logger, vpsNameId)
		}
		logger.Infof("Waiting for VPS %s to be ready...", vpsNameId)
		// 43 - launch the check at this frequency
		err = syncx.WaitForReady(waitCtx, logger, 10*time.Second, 5*time.Second, healthFunc)
		if err != nil {
			fmt.Printf("Resource not ready: %v\n", err)
		}

		// 5 - api display state after reinstall
		vpsDetail, err = ovh.VpsGetDetail(ctx, logger, vpsNameId)
		if err != nil {
			logger.Errorf("get detail failed: %v", err)
			os.Exit(1)
		}
		jsonx.PrettyPrintColor(vpsDetail["state"])
	},
}

func init() {
	installCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Bypass confirmation")
	installCmd.Flags().BoolP("list", "l", false, "List all available phases")
	installCmd.Flags().BoolP("runall", "r", false, "Run all phases in batch mode")
	// Make them mutually exclusive
	installCmd.MarkFlagsMutuallyExclusive("list", "runall")
}

// test04(ctx, logger)
// test03(ctx, logger)
// test01(logger)
// sshKeyList(ctx, logger)
// DetailMe(ctx, logger)
// listVps(ctx, logger)
// installVps(ctx, logger)
// vpsGetList(ctx, logger)
// vpsGetImageIdHandler(ctx, logger, "vps-9c33782a.vps.ovh.net")
