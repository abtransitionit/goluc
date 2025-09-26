package ovh

import (
	"context"
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

func vpsGetList(ctx context.Context, logger logx.Logger) {
	// get list of vps
	listVps, err := ovh.GetListVpsFromFile()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	// get list of os image
	jsonx.PrettyPrintColor(listVps)
}

func vpsGetImageIdHandler(ctx context.Context, logger logx.Logger, vpsNameId string) {
	id, err := ovh.GetVpsImageId(vpsNameId)
	if err != nil {
		logger.Errorf("%v", err)
		return
	}
	logger.Infof("Image ID: %s", id)
}
