package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadFile[T any](filePath string) (*T, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("reading file %q: %w", filePath, err)
	}
	var cfg T
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshalling YAML %q: %w", filePath, err)
	}
	return &cfg, nil
}
