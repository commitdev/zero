package config

import (
	"fmt"
	"github.com/commitdev/commit0/pkg/util/exit"
	"io/ioutil"
	"path"
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
  - source: "github.com/zthomas/commit0-terraform-basic"`

func CreateExample(projectName string) {
	content := []byte(fmt.Sprintf(exampleConfig, projectName))

	err := ioutil.WriteFile(path.Join(projectName, "commit0.yml"), content, 0644)
	if err != nil {
		exit.Fatal("Failed to create example commit.yml")
	}
}
