package ovh

import (
	"context"
	"fmt"
	"os"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
)

func listVps(ctx context.Context, logger logx.Logger) {
	// get list of vps id
	listVps, err := ovh.VpsGetList(ctx, logger)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	// Loop over slice
	for _, vpsId := range listVps {
		// request vps:id
		vpsInfo, err := ovh.VpsGetDetail(ctx, logger, vpsId)
		if err != nil {
			logger.Errorf("%v", err)
			os.Exit(1)
		}
		// request vps:os
		vpsOs, err := ovh.VpsGetOs(ctx, logger, vpsId)
		if err != nil {
			logger.Errorf("%v", err)
			os.Exit(1)
		}
		// filter the json
		vpsInfoSummary := jsonx.Json{
			"DisplayName": vpsInfo["displayName"],
			"IamId":       vpsInfo["iam"].(map[string]interface{})["id"],
			"MemoryLimit": vpsInfo["memoryLimit"],
			"Name":        vpsInfo["name"],
			"State":       vpsInfo["state"],
			"Vcore":       vpsInfo["vcore"],
			"OsId":        vpsOs["id"],
			"OsDistro":    vpsOs["name"],
		}
		jsonx.PrettyPrintColor(vpsInfoSummary)
	}
}

func installVps(ctx context.Context, logger logx.Logger) {
	// get ssh key id
	sshKeyId, err := ovh.SshKeyGetIdFromFile()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	// api get ssh detail
	sshKeyDetail, err := ovh.SshKeyGetDetail(ctx, logger, sshKeyId)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	// api get ssh public key
	sshPubKey, err := ovh.SshKeyGetPublic(ctx, logger, sshKeyDetail)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	// get OS image id
	imageId, err := ovh.GetOsImageId("vps-9c33782a.vps.ovh.net")
	if err != nil {
		logger.Errorf("Error:", err)
		os.Exit(1)
	}
	fmt.Println("OS Image ID:", imageId)
	os.Exit(0)

	// define the reinstall parameter
	reinstallParam := ovh.VpsReinstallParam{
		DoNotSendPassword: true,
		ImageId:           imageId,
		PublicSshKey:      sshPubKey, // example
	}

	vpsId := "vps-9c33782a.vps.ovh.net"
	jsonx.PrettyPrintColor(reinstallParam)
	fmt.Println(vpsId)

	// reinstall the vps via api
	vpsInfo, err := ovh.VpsReinstall(ctx, logger, vpsId, reinstallParam)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	jsonx.PrettyPrintColor(vpsInfo)
}

func vpsGetList(ctx context.Context, logger logx.Logger) {
	// get list of vps id
	listVps, err := ovh.GetListVpsFromFile()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	jsonx.PrettyPrintColor(listVps)
}
