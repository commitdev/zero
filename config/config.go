package config

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
)

type Maintainers struct {
	Name  string
	Email string
}

type Grpc struct {
	Host string
	Port int
}

type Graphql struct {
	Enabled bool
	Port    int
}

type Http struct {
	Enabled bool
	Port    int
}

type Network struct {
	Grpc    Grpc
	Http    Http
	Graphql Graphql
}

type Service struct {
	Name        string
	Description string
}

type CI struct {
	System       string `yaml:"system"`
	BuildImage   string `yaml:"build-image"`
	BuildCommand string `yaml:"build-command"`
	TestCommand  string `yaml:"test-command"`
}

type SproutConfig struct {
	Organization string        `yaml:"organization"`
	Name         string        `yaml:"name"`
	Description  string        `yaml:"description"`
	GitRepo      string        `yaml:"git-repo"`
	DockerRepo   string        `yaml:"docker-repo"`
	Maintainers  []Maintainers `yaml:"maintainers"`
	Network      Network       `yaml:"network"`
	Services     []Service     `yaml:"services"`
	CI           CI            `yaml:"ci"`
}

func LoadConfig(filePath string) *SproutConfig {
	config := &SproutConfig{}
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Panicf("failed to read config: %v", err)
	}
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Panicf("failed to unmarshall config: %v", err)
	}

	return config
}

func (c *SproutConfig) Print() {
	pp.Println(c)

}
