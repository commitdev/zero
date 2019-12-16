package config

type GeneratorConfig struct {
	Name    string
	Context map[string]string
	Modules []Module
}

type Module struct {
	Source string
	Params map[string]string
}

func LoadGeneratorConfig(filePath string, out interface{}) {
	config := &GeneratorConfig{}
	LoadYamlConfig(filePath, config)
}
