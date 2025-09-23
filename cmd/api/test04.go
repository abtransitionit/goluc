/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package api

import (
	"context"
	"os"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
)

func test04(ctx context.Context, logger logx.Logger) {
	token, err := ovh.GetAccessTokenFromServiceAccount(ctx, logger)
	if err != nil {
		logger.Errorf("error getting access token: %v", err)
		os.Exit(1)
	}

	logger.Infof("Access Token: %s", token)
	logger.Info("updated Access token successfully in credential file")
}
