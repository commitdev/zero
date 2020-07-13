package projectconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"text/template"

	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/flog"
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

// CreateProjectConfigFile extracts the required content for zero project config file then write to disk.
func CreateProjectConfigFile(dir string, projectName string, projectContext *ZeroProjectConfig) error {
	content, err := getProjectFileContent(*projectContext)
	if err != nil {
		return err
	}

	filePath := path.Join(dir, projectName, constants.ZeroProjectYml)
	flog.Debugf("Project file path: %s", filePath)
	writeErr := ioutil.WriteFile(filePath, []byte(content), 0644)
	if writeErr != nil {
		return err
	}

	return nil
}

func getProjectFileContent(projectConfig ZeroProjectConfig) (string, error) {
	var tplBuffer bytes.Buffer
	tmpl, err := template.New("projectConfig").Parse(zeroProjectConfigTemplate)
	if err != nil {
		return "", err
	}

	if len(projectConfig.Modules) == 0 {
		return "", fmt.Errorf("Invalid project config, expected config modules to be non-empty")
	}

	pConfigModules, err := yaml.Marshal(projectConfig.Modules)
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
		Modules:                util.IndentString(string(pConfigModules), 2),
	}

	if err := tmpl.Execute(&tplBuffer, t); err != nil {
		return "", err
	}
	return tplBuffer.String(), nil
}
