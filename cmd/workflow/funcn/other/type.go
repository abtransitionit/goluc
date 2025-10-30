package main

type Workflow struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Phases      map[string]struct {
		Description  string            `yaml:"description"`
		Node         string            `yaml:"node"`
		Fn           string            `yaml:"fn"`
		Params       map[string]string `yaml:"params"`
		Dependencies []string          `yaml:"dependencies"`
	} `yaml:"phases"`
}

type YamlFunc func()
