package config

import (
	"fmt"
	"github.com/commitdev/commit0/pkg/util/exit"
	"io/ioutil"
	"path"
)

const exampleConfig = `name: %s
	
# Context will populated automatically or could be added manually
context: 

modules: 
	# module can be in any format the go-getter supports (path, github, url, etc.)
	# supports https://github.com/hashicorp/go-getter#url-format
	# - source: "../tests/test_data/modules/ci"
	- source: "github.com/zthomas/react-mui-kit"`

func CreateExample(projectName string) {
	content := []byte(fmt.Sprintf(exampleConfig, projectName))

	err := ioutil.WriteFile(path.Join(projectName, "commit0.yml"), content, 0644)
	if err != nil {
		exit.Fatal("Failed to create example commit.yml")
	}
}
