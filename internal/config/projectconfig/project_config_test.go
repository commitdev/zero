package projectconfig_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestLoadConfig(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write([]byte(validConfigContent()))
	filePath := file.Name()

	want := &projectconfig.ZeroProjectConfig{
		Name:    "abc",
		Modules: eksGoReactSampleModules(),
	}

	t.Run("Should load and unmarshal config correctly", func(t *testing.T) {
		got := projectconfig.LoadConfig(filePath)
		if !cmp.Equal(want, got, cmpopts.EquateEmpty()) {
			t.Errorf("projectconfig.ZeroProjectConfig.Unmarshal mismatch (-want +got):\n%s", cmp.Diff(want, got))
		}
	})

}

func eksGoReactSampleModules() projectconfig.Modules {
	parameters := projectconfig.Parameters{"a": "b"}
	return projectconfig.Modules{
		"aws-eks-stack":             projectconfig.NewModule(parameters, "zero-aws-eks-stack", "github.com/something/repo1", "github.com/commitdev/zero-aws-eks-stack"),
		"deployable-backend":        projectconfig.NewModule(parameters, "zero-deployable-backend", "github.com/something/repo2", "github.com/commitdev/zero-deployable-backend"),
		"deployable-react-frontend": projectconfig.NewModule(parameters, "zero-deployable-react-frontend", "github.com/something/repo3", "github.com/commitdev/zero-deployable-react-frontend"),
	}
}

// TODO: Combinde TestMoudle and TestGetProjectFIleContent sample:global_config_test.go ln:42
// TODO: remove test when you complete TestGetProjectFIleContent
func TestProject(t *testing.T) {

	module := projectconfig.NewModule(projectconfig.Parameters{
		"repoName": "infrastructure",
		"region":   "us-east-1",
	}, "/infrastructure", "https://github.com/myorg/infrastructure")

	modules := projectconfig.Modules{
		"awk-eks-stack": module,
	}

	modulesContent, _ := yaml.Marshal(&modules)
	expectedContents := `awk-eks-stack:
  parameters:
    region: us-east-1
		repoName: infrastructure
		accountId: 12345
  files:
    dir: /infrastructure
    repo: https://github.com/myorg/infrastructure
`
	assert.Equal(t, strings.Trim(expectedContents, " "), strings.Trim(string(modulesContent), " "))
}

func TestGetProjectFileContent(t *testing.T) {

	module := projectconfig.NewModule(projectconfig.Parameters{
		"repoName": "infrastructure",
		"region":   "us-east-1",
	}, "/infrastructure", "https://github.com/myorg/infrastructure")

	modules := projectconfig.Modules{
		"awk-eks-stack": module,
	}

	config := projectconfig.ZeroProjectConfig{
		Name:                   "abc",
		ShouldPushRepositories: true,
		Infrastructure: projectconfig.Infrastructure{
			AWS: nil,
		},

		Parameters: map[string]string{},
		Modules:    modules,
	}
	content := projectconfig.GetProjectFileContent(config)

	// TODO: assert file output to make sure this works.
	fmt.Println(content)
}

func validConfigContent() string {
	return `
name: abc

shouldPushRepositories: true

modules:
    aws-eks-stack:
        parameters:
            a: b
        files:
            dir: zero-aws-eks-stack
            repo: github.com/something/repo1
            source: github.com/commitdev/zero-aws-eks-stack
    deployable-backend:
        parameters:
            a: b
        files:
            dir: zero-deployable-backend
            repo: github.com/something/repo2
            source: github.com/commitdev/zero-deployable-backend
    deployable-react-frontend:
        parameters:
            a: b
        files:
            dir: zero-deployable-react-frontend
            repo: github.com/something/repo3
            source: github.com/commitdev/zero-deployable-react-frontend
`
}
