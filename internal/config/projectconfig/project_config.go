package projectconfig

import (
	"errors"
	"io/ioutil"

	"github.com/hashicorp/terraform/dag"
	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"

	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
)

// GraphRootName represents the root of the graph of modules in a project
const GraphRootName = "graphRoot"

type ZeroProjectConfig struct {
	Name                   string `yaml:"name"`
	ShouldPushRepositories bool   `yaml:"shouldPushRepositories"`
	Parameters             map[string]string
	Modules                Modules `yaml:"modules"`
}

type Modules map[string]Module

type Module struct {
	DependsOn  []string   `yaml:"dependsOn,omitempty"`
	Parameters Parameters `yaml:"parameters,omitempty"`
	Files      Files
	Conditions []Condition `yaml:"conditions,omitempty"`
}

// ReadVendorCredentialsFromModule uses parsed project-config's module
// based on vendor parameter, retrieve the vendor's credential
// for pre-defined functionalities (eg: Github api key for pushing repos to github)
func ReadVendorCredentialsFromModule(m Module, vendor string) (error, string) {
	// this mapping could be useful for module config as well
	vendorToParamMap := map[string]string{
		"github":   "githubAccessToken",
		"circleci": "circleciApiKey",
	}
	if parameterKey, ok := vendorToParamMap[vendor]; ok {
		if val, ok := m.Parameters[parameterKey]; ok {
			return nil, val
		}
		return errors.New("Parameter not found in module."), ""
	}
	return errors.New("Unsupported vendor provided."), ""
}

type Parameters map[string]string

type Condition struct {
	Action     string   `yaml:"action"`
	MatchField string   `yaml:"matchField"`
	WhenValue  string   `yaml:"whenValue"`
	Data       []string `yaml:"data,omitempty"`
}

type Files struct {
	Directory  string `yaml:"dir,omitempty"`
	Repository string `yaml:"repo,omitempty"`
	Source     string
}

func LoadConfig(filePath string) *ZeroProjectConfig {
	config := &ZeroProjectConfig{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		exit.Fatal("failed to read config: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		exit.Fatal("failed to parse config: %v", err)
	}
	flog.Debugf("Loaded project config: %s from %s", config.Name, filePath)
	return config
}

func (c *ZeroProjectConfig) Print() {
	pp.Println(c)
}

// GetDAG returns a graph of the module names used in this project config
func (c *ZeroProjectConfig) GetDAG() dag.AcyclicGraph {
	var g dag.AcyclicGraph

	// Add vertices to graph
	g.Add(GraphRootName)
	for name := range c.Modules {
		g.Add(name)
	}

	// Connect modules in graph
	for name, m := range c.Modules {
		if len(m.DependsOn) == 0 {
			g.Connect(dag.BasicEdge(GraphRootName, name))
		} else {
			for _, dependencyName := range m.DependsOn {
				g.Connect(dag.BasicEdge(dependencyName, name))
			}
		}
	}
	return g
}

func NewModule(parameters Parameters, directory string, repository string, source string, dependsOn []string, conditions []Condition) Module {
	return Module{
		Parameters: parameters,
		DependsOn:  dependsOn,
		Files: Files{
			Directory:  directory,
			Repository: repository,
			Source:     source,
		},
		Conditions: conditions,
	}
}
