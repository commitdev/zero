package init_test

import (
	"testing"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	// init is a reserved word
	initPrompts "github.com/commitdev/zero/internal/init"

	"github.com/stretchr/testify/assert"
)

func TestGetParam(t *testing.T) {

	projectParams := map[string]string{}
	t.Run("Should execute params without prompt", func(t *testing.T) {
		param := moduleconfig.Parameter{
			Field:   "account-id",
			Execute: "echo \"my-acconut-id\"",
		}

		prompt := initPrompts.PromptHandler{
			param,
			initPrompts.NoCondition,
			initPrompts.NoValidation,
		}

		prompt.RunPrompt(projectParams)
		assert.Equal(t, "my-acconut-id", projectParams[param.Field])
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
		prompt.RunPrompt(projectParams)
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

		prompt.RunPrompt(projectParams)
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

		prompt.RunPrompt(projectParams)
		assert.Equal(t, "foo", projectParams["existing_value"])
		assert.Equal(t, "bar", projectParams[param.Field])
	})
}
