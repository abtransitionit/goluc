/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"github.com/abtransitionit/gocore/filex"
	"github.com/abtransitionit/gocore/logx"
	linuxfile "github.com/abtransitionit/golinux/filex"
)

func testScpAsSudo() {
	logger := logx.GetLogger()
	logger.Infof("testScp")
	source := "/tmp/goluc-linux"
	destination := "o1u:/usr/local/bin/goluc"

	copyok, err := linuxfile.ScpAsSudo(logger, source, destination)
	if err != nil {
		logger.Errorf("%v", err)
	}
	if copyok {
		logger.Info("remote file copied successfully")
	} else {
		logger.Error("remote file copy failed")
	}

}
func testScp() {
	logger := logx.GetLogger()
	logger.Infof("testScp")
	source := "/tmp/goluc-linux"
	destination := "o1u:/tmp/goluc"

	err := filex.Scp(logger, source, destination)
	if err != nil {
		logger.Errorf("%v", err)
	}
	logger.Info("remote file copied successfully")
}
