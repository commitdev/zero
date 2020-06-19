package projectconfig_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
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

func TestGetProjectFileContent(t *testing.T) {
	projectConfig := projectconfig.ZeroProjectConfig{
		Name:                   "abc",
		ShouldPushRepositories: false,
	}

	t.Run("Should fail if modules are missing from project config", func(t *testing.T) {
		// Remove the modules
		projectConfig.Modules = nil

		result, err := projectconfig.GetProjectFileContent(projectConfig)

		assert.Error(t, err)
		assert.Equal(t, result, "")
	})

	t.Run("Should return a valid project config", func(t *testing.T) {
		projectConfig.Modules = eksGoReactSampleModules()
		result, err := projectconfig.GetProjectFileContent(projectConfig)

		assert.NoError(t, err)
		assert.NotEmpty(t, result)
	})
}

// TODO: Write test init function to check file written to disk, use test_data/

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
