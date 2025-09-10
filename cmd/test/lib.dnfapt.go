/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/
package test

import (
	"fmt"

	"github.com/abtransitionit/gocore/logx"
	"github.com/abtransitionit/gocore/run"
	"github.com/abtransitionit/golinux/dnfapt"
	"github.com/abtransitionit/golinux/filex"
)

var sliceDaRepo dnfapt.SliceDaRepo

func getFile(logger logx.Logger) (string, error) {
	logger.Info("getFile")
	logger.Infof("%v", sliceDaRepo)

	// define test dataset
	vmName := "o2a"      // o1u
	vmOsFamily := "rhel" // debian
	repoFilePath := "/tmp/test"

	// get the repo file templated content
	repoFileContent, err := dnfapt.GetRepoFileContent(vmOsFamily, sliceDaRepo[0])
	if err != nil {
		return "", err
	}
	logger.Infof("repoFilePath: %s", repoFilePath)
	logger.Infof("vm is : %s", vmName)
	logger.Info("repoFileContent before saving it is")
	println(repoFileContent)

	// save the repo file
	cli := filex.CreateFileFromStringAsSudo(repoFilePath, repoFileContent)
	_, err = run.RunCliSsh(vmName, cli)
	if err != nil {
		return "", fmt.Errorf("failed to play cli %s on vm '%s': %w", cli, vmName, err)
	}

	// display the repo file content
	cli = fmt.Sprintf("ls -ial %s && echo && cat %s", repoFilePath, repoFilePath)
	output, err := run.RunCliSsh(vmName, cli)
	if err != nil {
		return "", fmt.Errorf("failed to play cli %s on vm '%s': %w", cli, vmName, err)
	}

	logger.Info("repoFileContent after saving it is")
	fmt.Println(output)
	// success

	return "", nil
}

func init() {
	sliceDaRepo = dnfapt.SliceDaRepo{
		{Name: "crio", FileName: "kbe-crio", Version: "1.32"},
		{Name: "k8s", FileName: "kbe-k8s", Version: "1.32"},
	}

}

// logger.Info("repoFileContent after saving is")
// println(output)
