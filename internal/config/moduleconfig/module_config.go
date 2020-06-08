package moduleconfig

import (
	"io/ioutil"

	"github.com/k0kubun/pp"
	yaml "gopkg.in/yaml.v2"
)

type ModuleConfig struct {
	Name                string
	Description         string
	Author              string
	TemplateConfig      `yaml:"template"`
	RequiredCredentials []string
	Parameters          []Parameter
}

type Parameter struct {
	Field   string
	Label   string   `yaml:"label,omitempty"`
	Options []string `yaml:"options,omitempty"`
	Execute string   `yaml:"execute,omitempty"`
	Value   string   `yaml:"value,omitempty"`
	Default string   `yaml:"default,omitempty"`
}

type TemplateConfig struct {
	StrictMode bool
	Delimiters []string
	InputDir   string
	OutputDir  string
}

func LoadModuleConfig(filePath string) (ModuleConfig, error) {
	config := ModuleConfig{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	pp.Println("Module Config:", config)
	return config, nil
}
