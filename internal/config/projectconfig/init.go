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
	content, err := GetProjectFileContent(*projectContext)
	if err != nil {
		exit.Fatal(fmt.Sprintf("Failed extracting the file config content %s", constants.ZeroProjectYml))
	}

	writeErr := ioutil.WriteFile(path.Join(dir, projectName, constants.ZeroProjectYml), []byte(content), 0644)
	if writeErr != nil {
		exit.Fatal(fmt.Sprintf("Failed to create config file %s", constants.ZeroProjectYml))
	}
}

func GetProjectFileContent(projectConfig ZeroProjectConfig) (string, error) {
	var tplBuffer bytes.Buffer
	tmpl, err := template.New("projectConfig").Parse(zeroProjectConfigTemplate)
	if err != nil {
		return "", err
	}

	if len(projectConfig.Modules) == 0 {
		return "", fmt.Errorf("Invalid project config, expected config modules to be non-empty")
	}

	pConfigModule, err := yaml.Marshal(projectConfig.Modules)
	if err != nil {
		return "", err
	}

	t := struct {
		Name                   string
		ShouldPushRepositories bool
		Modules                string
	}{
		Name:                   projectConfig.Name,
		ShouldPushRepositories: projectConfig.ShouldPushRepositories,
		Modules:                util.IndentString(string(pConfigModule), 2),
	}

	if err := tmpl.Execute(&tplBuffer, t); err != nil {
		return "", err
	}
	return tplBuffer.String(), nil
}
