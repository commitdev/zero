package config

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v2"
)

type maintainer struct {
	Name  string
	Email string
}

type grpc struct {
	Host string
	Port int
}

type graphql struct {
	Enabled bool
	Port    int
}

type http struct {
	Enabled bool
	Port    int
}

type web struct {
	Enabled bool
	Port    int
}

type network struct {
	Grpc    grpc
	Http    http
	Web     web
	Graphql graphql
}

type Service struct {
	Name        string
	Description string
	Language    string
	GitRepo     string `yaml:"gitRepo"`
	DockerRepo  string `yaml:"dockerRepo"`
	Network     network
	CI          CI
}

type CI struct {
	System       string
	Language     string
	BuildImage   string `yaml:"buildImage"`
	BuildTag     string `yaml:"buildTag"`
	BuildCommand string `yaml:"buildCommand"`
	TestCommand  string `yaml:"testCommand"`
}

type Commit0Config struct {
	Organization   string
	Name           string
	Description    string
	Maintainers    []maintainer
	Services       []Service
	Frontend       frontend
	Infrastructure infrastructure
}

type infrastructure struct {
	AWS aws
}

type aws struct {
	AccountId string `yaml:"accountId"`
	Region    string
	EKS       eks
	Cognito   cognito
	Terraform terraform
}

type terraform struct {
	RemoteState bool
}

type cognito struct {
	Deploy   bool
}

type eks struct {
	ClusterName string `yaml:"clusterName"`
	WorkerAMI   string `yaml:"workerAMI"`
	Deploy      bool
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
