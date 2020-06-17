package projectconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/exit"
	"gopkg.in/yaml.v2"
)

const zeroProjectConfigTemplate = `
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
	var tplBuffer bytes.Buffer
	tmpl, err := template.New("projectConfig").Parse(zeroProjectConfigTemplate)
	if err != nil {
		exit.Fatal(fmt.Sprintf("Failed to parse the config template %s", constants.ZeroProjectYml))
	}

	pConfig, err := yaml.Marshal(projectConfig.Modules)
	if err != nil {
		exit.Fatal(fmt.Sprintf("Failed while serializing the modules %s", constants.ZeroProjectYml))
	}

	t := struct {
		Name                   string
		ShouldPushRepositories bool
		Modules                string
	}{
		Name:                   projectConfig.Name,
		ShouldPushRepositories: projectConfig.ShouldPushRepositories,
		Modules:                util.IndentString(string(pConfig), 2),
	}

	if err := tmpl.Execute(&tplBuffer, t); err != nil {
		exit.Fatal(fmt.Sprintf("Failed while executing the template %s", constants.ZeroProjectYml))
	}
	result := tplBuffer.String()
	return result
}
