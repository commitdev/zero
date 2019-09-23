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

type Web struct {
	Enabled bool
	Port    int
}

type Network struct {
	Grpc    Grpc
	Http    Http
	Web     Web
	Graphql Graphql
}

type Service struct {
	Name        string
	Description string
}

type CI struct {
	System          string `yaml:"system"`
	BuildImage      string `yaml:"build-image"`
	BuildCommand    string `yaml:"build-command"`
	TestCommand     string `yaml:"test-command"`
	LanguageVersion string `yaml:"language-version"`
}

type Commit0Config struct {
	Language     string        `yaml:"string"`
	Organization string        `yaml:"organization"`
	Name         string        `yaml:"name"`
	Description  string        `yaml:"description"`
	GitRepo      string        `yaml:"git-repo"`
	DockerRepo   string        `yaml:"docker-repo"`
	Maintainers  []Maintainers `yaml:"maintainers"`
	Network      Network       `yaml:"network"`
	Services     []Service     `yaml:"services"`
	React        React         `yaml:"react"`
	Kubernetes   Kubernetes    `yaml:"kubernetes"`
	CI           CI            `yaml:"ci"`
}

type Kubernetes struct {
	ClusterName  string `yaml:"clusterName"`
	Deploy       bool   `yaml:"deploy"`
	AWSAccountId string `yaml:"awsAccountId"`
	AWSRegion    string `yaml:"awsRegion"`
}

func LoadConfig(filePath string) *Commit0Config {
	config := &Commit0Config{}
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

func (c *Commit0Config) Print() {
	pp.Println(c)

}
