package config

import (
	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Maintainers struct {
	Name  string
	Email string
}

type Network struct {
	HTTPPort    int `yaml:"http-port,omitempty"`
	GrpcPort    int `yaml:"grpc-port,omitempty"`
	GraphqlPort int `yaml:"graphql-port,omitempty"`
}

type Service struct {
	Name        string
	Description string
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
	pp.Print(c)

}
