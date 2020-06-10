package apply_test

import (
	"testing"

	"github.com/commitdev/zero/internal/apply"
	"github.com/commitdev/zero/internal/constants"
	"github.com/stretchr/testify/assert"
)

func TestApply(t *testing.T) {
	// @TODO is there a way to do this without relative paths?
	dir := "../../tests/test_data/sample_project/"
	applyConfigPath := constants.ZeroProjectYml
	applyEnvironments := []string{"staging", "production"}

	want := []string{
		"make module zero-aws-eks-stack\n",
		"make module zero-deployable-backend\n",
		"make module zero-deployable-react-frontend\n",
	}

	t.Run("Should run apply and execute make on each folder module", func(t *testing.T) {
		got := apply.Apply(dir, applyConfigPath, applyEnvironments)
		assert.ElementsMatch(t, want, got)
	})
}
