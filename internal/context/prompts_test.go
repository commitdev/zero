package context_test

import (
	"testing"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/commitdev/zero/internal/context"

	"github.com/stretchr/testify/assert"
)

func TestGetParam(t *testing.T) {
	projectParams := map[string]string{}
	t.Run("Should execute params without prompt", func(t *testing.T) {
		param := moduleconfig.Parameter{
			Field:   "account-id",
			Execute: "echo \"my-acconut-id\"",
		}

		prompt := context.PromptHandler{
			param,
			context.NoCondition,
		}

		result := prompt.GetParam(projectParams)
		assert.Equal(t, "my-acconut-id", result)
	})

	t.Run("executes with project context", func(t *testing.T) {
		param := moduleconfig.Parameter{
			Field:   "myEnv",
			Execute: "echo $INJECTEDENV",
		}

		prompt := context.PromptHandler{
			param,
			context.NoCondition,
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

		prompt := context.PromptHandler{
			param,
			context.NoCondition,
		}

		result := prompt.GetParam(projectParams)
		assert.Equal(t, "lorem-ipsum", result)
	})

}
