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

func TestLoadConfig(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())
	file.Write([]byte(validConfigContent()))
	filePath := file.Name()

	modules := projectconfig.InfrastructureSampleModules()
	sampleModules := projectconfig.EKSGoReactSampleModules()

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
