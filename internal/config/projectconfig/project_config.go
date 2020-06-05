package projectconfig

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	yaml "gopkg.in/yaml.v2"
)

type ZeroProjectConfig struct {
	Name                   string `yaml:"name"`
	ShouldPushRepositories bool
	Infrastructure         Infrastructure // TODO simplify and flatten / rename?
	Parameters             map[string]string
	Modules                Modules `yaml:"modules"`
}

type Infrastructure struct {
	AWS *AWS
}

type AWS struct {
	AccountID string `yaml:"accountId"`
	Region    string
	Terraform terraform // TODO simplify and flatten?
}

type terraform struct {
	RemoteState bool
}

type Modules map[string]Module

type Module struct {
	Parameters Parameters `yaml:"parameters"`
	Files      Files      `yaml:"files"`
}

type Parameters map[string]string

type Files struct {
	Directory  string `yaml:"dir,omitempty"`
	Repository string `yaml:"repo,omitempty"`
}

func LoadConfig(filePath string) *ZeroProjectConfig {
	config := &ZeroProjectConfig{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("failed to read config: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("failed to parse config: %v", err)
	}

	return config
}

func (c *ZeroProjectConfig) Print() {
	pp.Println(c)
}

// @TODO only an example, needs refactoring
func EKSGoReactSampleModules() Modules {
	parameters := Parameters{}
	return Modules{
		"zero-aws-eks-stack":             NewModule(parameters, "zero-aws-eks-stack", "github.com/commitdev/zero-aws-eks-stack"),
		"zero-deployable-backend":        NewModule(parameters, "zero-deployable-backend", "github.com/commitdev/zero-deployable-backend"),
		"zero-deployable-react-frontend": NewModule(parameters, "zero-deployable-react-frontend", "github.com/commitdev/zero-deployable-react-frontend"),
	}
}

// @TODO only an example, needs refactoring
func InfrastructureSampleModules() Modules {
	parameters := Parameters{
		"repoName":       "infrastructure",
		"region":         "us-east-1",
		"accountId":      "12345",
		"productionHost": "something.com",
	}
	return Modules{
		"infrastructure": NewModule(parameters, "infrastructure", "https://github.com/myorg/infrastructure"),
	}
}

func NewModule(parameters Parameters, directory string, repository string) Module {
	return Module{
		Parameters: parameters,
		Files: Files{
			Directory:  directory,
			Repository: repository,
		},
	}
}
