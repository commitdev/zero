package init_test

import (
	"testing"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	// init is a reserved word
	initPrompts "github.com/commitdev/zero/internal/init"

	"github.com/stretchr/testify/assert"
)

func TestGetParam(t *testing.T) {

	envVarTranslationMap := map[string]string{}
	projectParams := map[string]string{}
	t.Run("Should execute params without prompt", func(t *testing.T) {
		param := moduleconfig.Parameter{
			Field:   "account-id",
			Execute: "echo \"my-account-id\"",
		}

		prompt := initPrompts.PromptHandler{
			param,
			initPrompts.NoCondition,
			initPrompts.NoValidation,
		}

		prompt.RunPrompt(projectParams, envVarTranslationMap)
		assert.Equal(t, "my-account-id", projectParams[param.Field])
	})

	t.Run("executes with project context", func(t *testing.T) {
		param := moduleconfig.Parameter{
			Field:   "myEnv",
			Execute: "echo $INJECTEDENV",
		}

		prompt := initPrompts.PromptHandler{
			param,
			initPrompts.NoCondition,
			initPrompts.NoValidation,
		}

		projectParams := map[string]string{"INJECTEDENV": "SOME_ENV_VAR_VALUE"}
		prompt.RunPrompt(projectParams, envVarTranslationMap)
		assert.Equal(t, "SOME_ENV_VAR_VALUE", projectParams[param.Field])
	})

	t.Run("Should return static value", func(t *testing.T) {
		param := moduleconfig.Parameter{
			Field: "placeholder",
			Value: "lorem-ipsum",
		}

		prompt := initPrompts.PromptHandler{
			param,
			initPrompts.NoCondition,
			initPrompts.NoValidation,
		}

		prompt.RunPrompt(projectParams, envVarTranslationMap)
		assert.Equal(t, "lorem-ipsum", projectParams[param.Field])
	})

	t.Run("Prompt value to retain existing params", func(t *testing.T) {
		projectParams = map[string]string{
			"existing_value": "foo",
		}
		param := moduleconfig.Parameter{
			Field: "new_value",
			Value: "bar",
		}

		prompt := initPrompts.PromptHandler{
			param,
			initPrompts.NoCondition,
			initPrompts.NoValidation,
		}

		prompt.RunPrompt(projectParams, envVarTranslationMap)
		assert.Equal(t, "foo", projectParams["existing_value"])
		assert.Equal(t, "bar", projectParams[param.Field])
	})

	t.Run("Prompt to apply in order and allow EnvVarMapping", func(t *testing.T) {

		projectParams = map[string]string{}
		params := []moduleconfig.Parameter{
			{
				Field:      "param1",
				Value:      "foo",
				EnvVarName: "envvar1",
			},
			{
				Field:      "param2",
				Execute:    "echo $envvar1 bar",
				EnvVarName: "envvar2",
			},
			{
				Field:   "param3",
				Execute: "echo $envvar2 baz",
			},
		}
		module := moduleconfig.ModuleConfig{Parameters: params}

		projectParams, _ = initPrompts.PromptModuleParams(module, projectParams)

		assert.Equal(t, "foo", projectParams["param1"])
		assert.Equal(t, "foo bar", projectParams["param2"], "should reference param1 via env-var")
		assert.Equal(t, "foo bar baz", projectParams["param3"], "should reference param2 via env-var")
	})

	t.Run("Prompt conditions", func(t *testing.T) {

		projectParams = map[string]string{}
		params := []moduleconfig.Parameter{
			{
				Field: "param1",
				Value: "pass",
			},
			{
				Field: "passing_condition",
				Value: "pass",
				Conditions: []moduleconfig.Condition{
					{
						Action:     "KeyMatchCondition",
						MatchField: "param1",
						WhenValue:  "pass",
					},
				},
			},
			{
				Field: "failing_condition",
				Value: "pass",
				Conditions: []moduleconfig.Condition{
					{
						Action:     "KeyMatchCondition",
						MatchField: "param1",
						WhenValue:  "not pass",
					},
				},
			},
			{
				Field: "multiple_condition",
				Value: "pass",
				Conditions: []moduleconfig.Condition{
					{
						Action:     "KeyMatchCondition",
						MatchField: "param1",
						WhenValue:  "pass",
					},
					{
						Action:     "KeyMatchCondition",
						MatchField: "passing_condition",
						WhenValue:  "pass",
					},
				},
			},
			{
				Field: "condition_with_default",
				Value: "pass",
				Conditions: []moduleconfig.Condition{
					{
						Action:     "KeyMatchCondition",
						MatchField: "param1",
						WhenValue:  "not pass",
						ElseValue:  "itsadefault",
					},
				},
			},
		}
		module := moduleconfig.ModuleConfig{Parameters: params}
		projectParams, _ = initPrompts.PromptModuleParams(module, projectParams)

		assert.Equal(t, "pass", projectParams["param1"], "Value just hardcoded")
		assert.Equal(t, "pass", projectParams["passing_condition"], "Expected to pass condition and set value")
		assert.NotContains(t, projectParams, "failing_condition", "Expected to fail condition and not set value")
		assert.Equal(t, "pass", projectParams["multiple_condition"], "Expected to pass multiple condition and set value")
		assert.Equal(t, "itsadefault", projectParams["condition_with_default"], "Expected to set a default value for a condition that failed")
	})

	t.Run("Should return error upon unsupported custom prompt type", func(t *testing.T) {

		projectParams = map[string]string{}
		params := []moduleconfig.Parameter{
			{
				Field: "param1",
				Type:  "random-type",
			},
		}
		module := moduleconfig.ModuleConfig{Parameters: params}
		_, err := initPrompts.PromptModuleParams(module, projectParams)
		assert.Equal(t, "Unsupported custom prompt type random-type.", err.Error())
	})
}

func TestValidateProjectNam(t *testing.T) {
	t.Run("Should return error upon invalid project name", func(t *testing.T) {
		err := initPrompts.ValidateProjectName("0invalid")
		assert.Error(t, err, "Project name should not start with a number")
	})

	t.Run("Should return error upon invalid project name length", func(t *testing.T) {
		err := initPrompts.ValidateProjectName("invalid name with more than 30 characters")
		assert.Error(t, err, "Project name should not be longer than 30 characters")
	})

	t.Run("Should return nil upon valid project name", func(t *testing.T) {
		err := initPrompts.ValidateProjectName("valid-name")
		assert.Nil(t, err)
	})
}
