package projectconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/exit"
	"gopkg.in/yaml.v2"
)

// {{  .ShouldPushRepositories | printf "%q"}}
const exampleConfig = `
# Templated zero-project.yml file
name: {{.Name}}

shouldPushRepositories: {{.ShouldPushRepositories | printf "%v"}}

modules:
{{.Modules}}
`

var RootDir = "./"

func SetRootDir(dir string) {
	RootDir = dir
}

func Init(dir string, projectName string, projectContext *ZeroProjectConfig) {
	// TODO: template the zero-project.yml with projectContext
	// content := []byte(fmt.Sprintf(exampleConfig, projectName))
	content, err := yaml.Marshal(projectContext)
	if err != nil {
		exit.Fatal(fmt.Sprintf("Failed to serialize configuration file %s", constants.ZeroProjectYml))
	}

	writeErr := ioutil.WriteFile(path.Join(dir, projectName, constants.ZeroProjectYml), content, 0644)
	if writeErr != nil {
		exit.Fatal(fmt.Sprintf("Failed to create config file %s", constants.ZeroProjectYml))
	}
}

func GetProjectFileContent(projectConfig ZeroProjectConfig) string {
	var tpl bytes.Buffer
	tmpl := template.New("sa")
	tmpl, err := tmpl.Parse(exampleConfig)
	if err != nil {
		exit.Fatal(fmt.Sprintf("Failed to parse the sample Zero module config file %s", constants.ZeroProjectYml))
	}

	type tempProjectConfig struct {
		Name                   string
		ShouldPushRepositories bool
		Modules                string
	}

	foo, err := yaml.Marshal(projectConfig.Modules)
	if err != nil {
		fmt.Println(err)
	}

	t := tempProjectConfig{
		Name:                   projectConfig.Name,
		ShouldPushRepositories: projectConfig.ShouldPushRepositories,
		Modules:                util.IndentString(string(foo), 2),
	}

	tmpl.Execute(os.Stdout, t)
	result := tpl.String()
	return result
}
