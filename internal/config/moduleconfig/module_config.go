package moduleconfig

import (
	"io/ioutil"

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
	Credentials []string `yaml:"requiredCredentials"`
	Prompts     []Prompt `yaml:"parameters"`
}

type Prompt struct {
	Field   string `yaml:"field,omitempty"`
	Label   string
	Options []string `yaml:"options,omitempty"`
	Execute string   `yaml:"execute,omitempty"`
}

type TemplateConfig struct {
	StrictMode bool
	Delimiters []string
	Output     string
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
