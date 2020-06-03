package projectconfig

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/pkg/util/exit"
)

const exampleConfig = `name: %s

# Context is normally populated automatically but could be used to inject global params
context:

# module can be in any format the go-getter supports (path, github, url, etc.)
# supports https://github.com/hashicorp/go-getter#url-format
# Example:
# - source: "../development/modules/ci"
# - output: "github-actions"

modules:
	- source: "github.com/commitdev/zero-aws-eks-stack"
	- source: "github.com/commitdev/zero-deployable-backend"
	- source: "github.com/commitdev/zero-deployable-react-frontend"
`

var RootDir = "./"

func SetRootDir(dir string) {
	RootDir = dir
}

func Init(dir string, projectName string, projectContext *ZeroProjectConfig) {
	// TODO: template the zero-project.yml with projectContext
	content := []byte(fmt.Sprintf(exampleConfig, projectName))

	err := ioutil.WriteFile(path.Join(dir, projectName, constants.ZeroProjectYml), content, 0644)
	if err != nil {
		exit.Fatal(fmt.Sprintf("Failed to create example %s", constants.ZeroProjectYml))
	}
}
