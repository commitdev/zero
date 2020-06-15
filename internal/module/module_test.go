package module_test

import (
	"errors"
	"testing"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/stretchr/testify/assert"

	"github.com/commitdev/zero/internal/module"
)

func TestGetSourceDir(t *testing.T) {
	source := "tests/test_data/modules"
	relativeSource := source
	dir := module.GetSourceDir(source)

	t.Log("dir", dir)
	if dir != relativeSource {
		t.Errorf("Error, local sources should not be changed: %s", source)
	}

	source = "github.com/commitdev/my-repo"
	dir = module.GetSourceDir(source)
	if dir == relativeSource {
		t.Errorf("Error, remote sources should be converted to a local dir: %s", source)
	}
}

func TestParseModuleConfig(t *testing.T) {
	testModuleSource := "../../tests/test_data/modules/ci"
	var mod moduleconfig.ModuleConfig

	t.Run("Loading module from source", func(t *testing.T) {
		mod, _ = module.ParseModuleConfig(testModuleSource)

		assert.Equal(t, "CI templates", mod.Name)
	})

	t.Run("Parameters are loaded", func(t *testing.T) {
		param, err := findParameter(mod.Parameters, "platform")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "platform", param.Field)
		assert.Equal(t, "CI Platform", param.Label)
	})

	t.Run("requiredCredentials are loaded", func(t *testing.T) {
		assert.Equal(t, []string{"aws", "circleci", "github"}, mod.RequiredCredentials)
	})

	t.Run("TemplateConfig is unmarshaled", func(t *testing.T) {
		mod, _ = module.ParseModuleConfig(testModuleSource)
		assert.Equal(t, ".circleci", mod.TemplateConfig.OutputDir)
		assert.Equal(t, "templates", mod.TemplateConfig.InputDir)
		assert.Equal(t, []string{"<%", "%>"}, mod.TemplateConfig.Delimiters)
	})

}

func findParameter(params []moduleconfig.Parameter, field string) (moduleconfig.Parameter, error) {
	for _, v := range params {
		if v.Field == field {
			return v, nil
		}
	}
	return moduleconfig.Parameter{}, errors.New("parameter not found")
}
