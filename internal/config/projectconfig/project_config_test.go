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

	type args struct {
		filePath string
	}

	modules := projectconfig.EKSGoReactSampleModules()
	infrastructureModules := projectconfig.InfrastructureSampleModules()

	for k, v := range infrastructureModules {
		modules[k] = v
	}

	expected := &projectconfig.ZeroProjectConfig{
		Name:    "abc",
		Modules: modules,
	}

	tests := []struct {
		name string
		args args
		want *projectconfig.ZeroProjectConfig
	}{
		{
			"Working config",
			args{filePath: file.Name()},
			expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// @TODO handle nil/empty map unmarshall case?
			if got := projectconfig.LoadConfig(tt.args.filePath); !cmp.Equal(got, tt.want, cmpopts.EquateEmpty()) {
				t.Errorf(cmp.Diff(got, tt.want))
			}
		})
	}
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
    zero-aws-eks-stack:
        files:
            dir: zero-aws-eks-stack
            repo: github.com/commitdev/zero-aws-eks-stack
    zero-deployable-backend:
        files:
            dir: zero-deployable-backend
            repo: github.com/commitdev/zero-deployable-backend
    zero-deployable-react-frontend:
        files:
            dir: zero-deployable-react-frontend
            repo: github.com/commitdev/zero-deployable-react-frontend
`
}
