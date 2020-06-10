package projectconfig

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/pkg/util/exit"
	"gopkg.in/yaml.v2"
)

const exampleConfig = `name: %s

# Context is normally populated automatically but could be used to inject global params
context:

# module can be in any format the go-getter supports (path, github, url, etc.)
# supports https://github.com/hashicorp/go-getter#url-format
# Example:
# - repo: "../development/modules/ci"
# - dir: "github-actions"
modules:
	aws-eks-stack:
		parameters:
			repoName: infrastructure
			region: us-east-1
			accountId: 12345
			productionHost: something.com
		files:
			dir: infrastructure
			repo: https://github.com/myorg/infrastructure
	some-other-module:
		parameters:
			repoName: api
		files:
			dir: api
			repo: https://github.com/myorg/api


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
