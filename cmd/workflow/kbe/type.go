/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package kbe

import (
	"os"

	"gopkg.in/yaml.v3"

	linuxdnfapt "github.com/abtransitionit/golinux/dnfapt"
	linuxk8s "github.com/abtransitionit/golinux/k8s"
)

type Config struct {
	K8sVersionShort  string `yaml:"k8sVersionShort"`
	CustomRcFileName string `yaml:"customRcFileName"`
	BinFolderPath    string `yaml:"binFolderPath"`

	Node struct {
		ControlPlane []string `yaml:"controlPlane"`
		Worker       []string `yaml:"worker"`
		All          []string `yaml:"all"`
	} `yaml:"node"`

	GoCli struct {
		Cluster []struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"cluster"`
		ControlPlane []struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"controlPlane"`
	} `yaml:"goCli"`

	Da struct {
		Repo struct {
			Node     linuxdnfapt.SliceDaRepo `yaml:"node"`
			Name     string                  `yaml:"name"`
			FileName string                  `yaml:"fileName"`
			Version  string                  `yaml:"version"`
		} `yaml:"repo"`
		Pkg struct {
			ControlPlane linuxdnfapt.SliceDaPack `yaml:"controlPlane"`
			Node         linuxdnfapt.SliceDaPack `yaml:"node"`
			Required     linuxdnfapt.SliceDaPack `yaml:"required"`
		} `yaml:"package"`
	} `yaml:"da"`

	Helm struct {
		Repo []struct {
			Name string `yaml:"name"`
		} `yaml:"repo"`
		Release []struct {
			Name      string `yaml:"name"`
			Chart     string `yaml:"chart"`
			Namespace string `yaml:"namespace"`
		} `yaml:"release"`
	} `yaml:"helm"`

	Cluster linuxk8s.K8sConf `yaml:"cluster"`
}

// Description: load yaml file. and return a struct
func loadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = yaml.Unmarshal(data, &conf)
	return &conf, err
}
