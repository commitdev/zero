package registry

import (
	"io/ioutil"

	"github.com/commitdev/zero/internal/constants"
	"github.com/hashicorp/go-getter"

	yaml "gopkg.in/yaml.v2"
)

type Registry []Stack

type Stack struct {
	Name          string   `yaml:"name"`
	ModuleSources []string `yaml:"moduleSources"`
}

func GetRegistry(localModulePath, registryFilePath string) (Registry, error) {
	registry := Registry{}

	err := getter.GetFile(constants.TmpRegistryYml, registryFilePath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadFile(constants.TmpRegistryYml)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, &registry)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(registry); i++ {
		for j := 0; j < len(registry[i].ModuleSources); j++ {
			registry[i].ModuleSources[j] = localModulePath + registry[i].ModuleSources[j]
		}
	}

	return registry, nil
}

func GetModulesByName(registry Registry, name string) []string {
	for _, v := range registry {
		if v.Name == name {
			return v.ModuleSources
		}
	}
	return []string{}
}

func AvailableLabels(registry Registry) []string {
	labels := make([]string, len(registry))
	i := 0
	for _, stack := range registry {
		labels[i] = stack.Name
		i++
	}
	return labels
}
