package test

import (
	"context"
	"os"

	"github.com/abtransitionit/gocore/jsonx"
	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
)

func DetailMe(ctx context.Context, logger logx.Logger) {
	MeInfo, err := ovh.MeGetInfo(ctx, logger)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	jsonx.PrettyPrintColor(MeInfo)
}
