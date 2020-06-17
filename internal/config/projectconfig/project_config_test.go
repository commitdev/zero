package projectconfig_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

// @TODO only an example, needs refactoring
func eksGoReactSampleModules() Modules {
	parameters := Parameters{}
	return Modules{
		"aws-eks-stack":             projectconfig.NewModule(parameters, "zero-aws-eks-stack", "github.com/something/repo1", "github.com/commitdev/zero-aws-eks-stack"),
		"deployable-backend":        projectconfig.NewModule(parameters, "zero-deployable-backend", "github.com/something/repo2", "github.com/commitdev/zero-deployable-backend"),
		"deployable-react-frontend": projectconfig.NewModule(parameters, "zero-deployable-react-frontend", "github.com/something/repo3", "github.com/commitdev/zero-deployable-react-frontend"),
	}
}

// @TODO only an example, needs refactoring
func sampleModules() Modules {
	parameters := Parameters{
		"repoName":       "infrastructure",
		"region":         "us-east-1",
		"accountId":      "12345",
		"productionHost": "something.com",
	}
	return Modules{
		"infrastructure": projectconfig.NewModule(parameters, "infrastructure", "https://github.com/myorg/infrastructure", "github.com/commitdev/zero-aws-eks-stack"),
	}
}
func TestLoadConfig(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write([]byte(validConfigContent()))
	filePath := file.Name()

	modules := sampleModules()
	sampleModules := eksGoReactSampleModules()

	for k, v := range sampleModules {
		modules[k] = v
	}

	want := &projectconfig.ZeroProjectConfig{
		Name:    "abc",
		Modules: modules,
	}

	t.Run("Should load and unmarshall config correctly", func(t *testing.T) {
		got := projectconfig.LoadConfig(filePath)
		if !cmp.Equal(want, got, cmpopts.EquateEmpty()) {
			t.Errorf("projectconfig.ZeroProjectConfig.Unmarshal mismatch (-want +got):\n%s", cmp.Diff(want, got))
		}
	})

}

func validConfigContent() string {
	return `
name: abc

context:

modules:
    infrastructure:
        parameters:
            repoName: infrastructure
            region: us-east-1
            accountId: 12345
            productionHost: something.com
        files:
            dir: infrastructure
            repo: https://github.com/myorg/infrastructure
    aws-eks-stack:
        files:
            dir: zero-aws-eks-stack
            repo: github.com/commitdev/zero-aws-eks-stack
    deployable-backend:
        files:
            dir: zero-deployable-backend
            repo: github.com/commitdev/zero-deployable-backend
    deployable-react-frontend:
        files:
            dir: zero-deployable-react-frontend
            repo: github.com/commitdev/zero-deployable-react-frontend
`
}
