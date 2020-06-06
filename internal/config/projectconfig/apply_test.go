package projectconfig_test

import (
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	// @TODO is there a way to do this without relative paths?
	// execCMD will use the current folder as target...
	dir := "../../../tests/test_data/"
	projectName := "sample_project"
	projectContext := &projectconfig.ZeroProjectConfig{
		Name:    projectName,
		Modules: projectconfig.EKSGoReactSampleModules(),
	}
	applyEnvironments := []string{"staging", "production"}

	want := []string{
		"make module zero-aws-eks-stack\n",
		"make module zero-deployable-backend\n",
		"make module zero-deployable-react-frontend\n",
	}

	t.Run("Should run apply and execute make on each folder module", func(t *testing.T) {
		got := projectconfig.Apply(dir, projectContext, applyEnvironments)
		assert.Equal(t, want, got)
	})
}
