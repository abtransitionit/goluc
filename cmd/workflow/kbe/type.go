/*
Copyright © 2025 AB TRANSITION IT abtransitionit@hotmail.com
*/

package kbe

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	CmdName          string `yaml:"cmdName"`
	Description      string `yaml:"description"`
	K8sVersion       string `yaml:"k8sVersion"`
	K8sVersionShort  string `yaml:"k8sVersionShort"`
	CustomRcFileName string `yaml:"customRcFileName"`
	BinFolderPath    string `yaml:"binFolderPath"`

	Nodes struct {
		ControlPlane []string `yaml:"controlPlane"`
		Worker       []string `yaml:"worker"`
	} `yaml:"node"`

	Go struct {
		Cli []struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		} `yaml:"cli"`
	} `yaml:"go"`

	Da struct {
		Repo []struct {
			Name     string `yaml:"name"`
			FileName string `yaml:"fileName"`
			Version  string `yaml:"version"`
		} `yaml:"repo"`
		Package struct {
			ControlPlane []struct {
				Name string `yaml:"name"`
			} `yaml:"controlPlane"`
			Node []struct {
				Name string `yaml:"name"`
			} `yaml:"node"`
			Required []struct {
				Name string `yaml:"name"`
			} `yaml:"required"`
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

	Cluster struct {
		PodCidr      string `yaml:"PodCidr"`
		ServiceCidr  string `yaml:"ServiceCidr"`
		CrSocketName string `yaml:"crSocketName"`
	} `yaml:"cluster"`
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
