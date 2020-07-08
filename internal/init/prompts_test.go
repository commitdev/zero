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

		result := prompt.GetParam(projectParams)
		assert.Equal(t, "my-acconut-id", result)
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

		result := prompt.GetParam(map[string]string{
			"INJECTEDENV": "SOME_ENV_VAR_VALUE",
		})
		assert.Equal(t, "SOME_ENV_VAR_VALUE", result)
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

		result := prompt.GetParam(projectParams)
		assert.Equal(t, "lorem-ipsum", result)
	})

}
