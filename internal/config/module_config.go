package config

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	yaml "gopkg.in/yaml.v2"
)

type ModuleConfig struct {
	Name        string
	Description string
	Author      string
	Icon        string
	Thumbnail   string
	Template    TemplateConfig
	Prompts     []Prompt
}

type Prompt struct {
	Field   string
	Label   string
	Options []string
}

type TemplateConfig struct {
	StrictMode  bool
	Delimiters  []string
	Destination string
}

func LoadModuleConfig(filePath string) *ModuleConfig {
	config := &ModuleConfig{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("failed to read config: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}
	pp.Println(config)
	return config
}
