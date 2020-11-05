package moduleconfig

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type ModuleConfig struct {
	Name                string
	Description         string
	Author              string
	DependsOn           []string `yaml:"dependsOn,omitempty"`
	TemplateConfig      `yaml:"template"`
	RequiredCredentials []string `yaml:"requiredCredentials"`
	Parameters          []Parameter
	Conditions          []Condition `yaml:"conditions,omitempty"`
}

type Parameter struct {
	Field           string
	Label           string   `yaml:"label,omitempty"`
	Options         []string `yaml:"options,omitempty"`
	Execute         string   `yaml:"execute,omitempty"`
	Value           string   `yaml:"value,omitempty"`
	Default         string   `yaml:"default,omitempty"`
	Info            string   `yaml:"info,omitempty"`
	FieldValidation Validate `yaml:"fieldValidation,omitempty"`
}

type Condition struct {
	Action     string   `yaml:"action"`
	MatchField string   `yaml:"matchField"`
	WhenValue  string   `yaml:"whenValue"`
	Data       []string `yaml:"data,omitempty"`
}

type Validate struct {
	Type         string `yaml:"type,omitempty"`
	Value        string `yaml:"value,omitempty"`
	ErrorMessage string `yaml:"errorMessage,omitempty"`
}

type TemplateConfig struct {
	StrictMode bool
	Delimiters []string
	InputDir   string `yaml:"inputDir"`
	OutputDir  string `yaml:"outputDir"`
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

	return config, nil
}
