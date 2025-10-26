/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package kbe

import (
	"os"

	core_helm "github.com/abtransitionit/gocore/k8s-helm"
	"github.com/abtransitionit/golinux/da"
	"gopkg.in/yaml.v3"
)

type Config struct {
	K8sVersion       string
	CustomRcFileName string
	BinFolderPath    string
	Nodes            struct {
		ControlPlane []string
		Worker       []string
	}
	Dnfapt struct {
		Repos    []da.Repo
		Packages []da.Package
	}
	Helm struct {
		Repos    []core_helm.HelmRepo
		Releases []core_helm.HelmRelease
	}
}

// Description: load yaml file. and return a struct
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(data, &conf)
	return &conf, err
}
