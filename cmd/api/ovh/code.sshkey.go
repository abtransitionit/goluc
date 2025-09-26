package ovh

import (
	"context"
	"os"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/ovh"
)

func sshKeyList(ctx context.Context, logger logx.Logger) {
	// get ssh key id
	sshKeyId, err := ovh.SshKeyGetIdFromFileCached()
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	logger.Infof("SshKeyId: %s", sshKeyId)

	// get ssh detail
	sshKeyDetail, err := ovh.SshKeyGetDetail(ctx, logger, sshKeyId)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}

	// get ssh public key
	sshPubKey, err := ovh.SshKeyGetPublic(ctx, logger, sshKeyDetail)
	if err != nil {
		logger.Errorf("%v", err)
		os.Exit(1)
	}
	logger.Infof("SshPubKey: %s", sshPubKey)

}
