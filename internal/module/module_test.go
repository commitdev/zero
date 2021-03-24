package module_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/stretchr/testify/assert"

	"github.com/commitdev/zero/internal/module"
	"github.com/commitdev/zero/version"
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
		moduleconfig.ValidateZeroVersion(mod)

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

	t.Run("OmitFromProjectFile default", func(t *testing.T) {
		param, err := findParameter(mod.Parameters, "platform")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, false, param.OmitFromProjectFile, "OmitFromProjectFile should default to false")
		useCredsParam, useCredsErr := findParameter(mod.Parameters, "useExistingAwsProfile")
		if useCredsErr != nil {
			panic(useCredsErr)
		}
		assert.Equal(t, true, useCredsParam.OmitFromProjectFile, "OmitFromProjectFile should be read from file")
	})

	t.Run("Parsing Conditions and Typed prompts from config", func(t *testing.T) {
		param, err := findParameter(mod.Parameters, "profilePicker")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "AWSProfilePicker", param.Type)
		assert.Equal(t, "KeyMatchCondition", param.Conditions[0].Action)
		assert.Equal(t, "useExistingAwsProfile", param.Conditions[0].MatchField)
		assert.Equal(t, "yes", param.Conditions[0].WhenValue)
	})

	t.Run("parsing envVarName from module config", func(t *testing.T) {
		param, err := findParameter(mod.Parameters, "accessKeyId")
		if err != nil {
			panic(err)
		}
		assert.Equal(t, "AWS_ACCESS_KEY_ID", param.EnvVarName)
	})

	t.Run("TemplateConfig is unmarshaled", func(t *testing.T) {
		mod, _ = module.ParseModuleConfig(testModuleSource)
		assert.Equal(t, ".circleci", mod.TemplateConfig.OutputDir)
		assert.Equal(t, "templates", mod.TemplateConfig.InputDir)
		assert.Equal(t, []string{"<%", "%>"}, mod.TemplateConfig.Delimiters)
	})

	t.Run("Parsing zero version constraints", func(t *testing.T) {
		moduleConstraints := mod.ZeroVersion.Constraints.String()
		assert.Equal(t, ">= 3.0.0, < 4.0.0", moduleConstraints)
	})

	t.Run("Should Fail against old zero version", func(t *testing.T) {
		moduleConstraints := mod.ZeroVersion.Constraints.String()

		// Mocking zero's version, testing against ">= 3.0.0, <= 4.0.0"
		originalVersion := version.AppVersion
		version.AppVersion = "2.0.0"
		defer func() { version.AppVersion = originalVersion }()
		// end of mock

		isValid := moduleconfig.ValidateZeroVersion(mod)
		assert.Equal(t, false, isValid, fmt.Sprintf("Version should satisfy %s", moduleConstraints))
	})

	t.Run("Should Fail against too new zero version", func(t *testing.T) {
		moduleConstraints := mod.ZeroVersion.Constraints.String()

		// Mocking zero's version, testing against ">= 3.0.0, <= 4.0.0"
		originalVersion := version.AppVersion
		version.AppVersion = "4.0.0"
		defer func() { version.AppVersion = originalVersion }()
		// end of mock

		isValid := moduleconfig.ValidateZeroVersion(mod)
		assert.Equal(t, false, isValid, fmt.Sprintf("Version should satisfy %s", moduleConstraints))
	})

	t.Run("Should validate against valid versions", func(t *testing.T) {
		moduleConstraints := mod.ZeroVersion.Constraints.String()

		// Mocking zero's version, testing against ">= 3.0.0, <= 4.0.0"
		const newZeroVersion = "3.0.5"
		originalVersion := version.AppVersion
		version.AppVersion = newZeroVersion
		defer func() { version.AppVersion = originalVersion }()
		// end of mock

		isValid := moduleconfig.ValidateZeroVersion(mod)
		assert.Equal(t, true, isValid, fmt.Sprintf("Version should satisfy %s", moduleConstraints))
	})

	t.Run("default to SNAPSHOT version passes tests", func(t *testing.T) {
		assert.Equal(t, "SNAPSHOT", version.AppVersion)
		isValid := moduleconfig.ValidateZeroVersion(mod)
		assert.Equal(t, true, isValid, "default test run should pass version constraint")
	})

}

func TestModuleWithNoVersionConstraint(t *testing.T) {
	testModuleSource := "../../tests/test_data/modules/no-version-constraint"
	var mod moduleconfig.ModuleConfig
	var err error

	t.Run("Parsing Module with no version constraint", func(t *testing.T) {
		mod, err = module.ParseModuleConfig(testModuleSource)
		assert.Equal(t, "", mod.ZeroVersion.String())
		assert.Nil(t, err)
	})

	t.Run("Should pass Validation if constraint not specified", func(t *testing.T) {
		isValid := moduleconfig.ValidateZeroVersion(mod)
		assert.Equal(t, true, isValid, "Module with no constraint should pass version validation")
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
