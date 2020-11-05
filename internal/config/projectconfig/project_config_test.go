package projectconfig_test

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
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
		Name:                   "abc",
		ShouldPushRepositories: true,
		Modules:                eksGoReactSampleModules(),
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
		"aws-eks-stack":             projectconfig.NewModule(parameters, "zero-aws-eks-stack", "github.com/something/repo1", "github.com/commitdev/zero-aws-eks-stack", []string{}, []projectconfig.Condition{}),
		"deployable-backend":        projectconfig.NewModule(parameters, "zero-deployable-backend", "github.com/something/repo2", "github.com/commitdev/zero-deployable-backend", []string{}, []projectconfig.Condition{}),
		"deployable-react-frontend": projectconfig.NewModule(parameters, "zero-deployable-react-frontend", "github.com/something/repo3", "github.com/commitdev/zero-deployable-react-frontend", []string{}, []projectconfig.Condition{}),
	}
}

func validConfigContent() string {
	return `
# Templated zero-project.yml file
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

func TestProjectConfigModuleGraph(t *testing.T) {
	configPath := filepath.Join("../../../tests/test_data/projectconfig/", constants.ZeroProjectYml)

	t.Run("Should generate a valid, correct graph based on the project config", func(t *testing.T) {
		pc := projectconfig.LoadConfig(configPath)
		graph := pc.GetDAG()

		// Validate the graph
		assert.NoError(t, graph.Validate())

		// Check the structure of the graph
		root, err := graph.Root()
		assert.NoError(t, err)
		assert.Equal(t, "graphRoot", root)

		want := `graphRoot
  project1
project1
  project2
  project3
project2
  project4
project3
  project4
  project5
project4
project5
`
		assert.Equal(t, want, graph.String())

	})

}
