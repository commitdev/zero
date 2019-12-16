package config

type ModuleConfig struct {
	Name        string
	Description string
	Author      string
	Icon        string
	Thumbnail   string
	Template    TemplateConfig
	Prompts     []Prompt
}

type Prompt struct {
	Field   string
	Label   string
	Options map[string]string
}

type TemplateConfig struct {
	StrictMode  bool
	Delimiters  []string
	Destination string
}

func LoadModuleConfig(filePath string, out interface{}) {
	config := &ModuleConfig{}
	LoadYamlConfig(filePath, config)
}
