package config

import (
	"io/ioutil"
	"log"

	"github.com/k0kubun/pp"
	yaml "gopkg.in/yaml.v2"
)

type Commit0Config struct {
	Name           string
	Infrastructure Infrastructure // TODO simplify and flatten / rename?
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
