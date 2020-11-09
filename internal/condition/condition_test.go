package condition_test

import (
	"encoding/base64"
	"math/rand"
	"os"
	"testing"

	"github.com/commitdev/zero/internal/condition"
	"github.com/commitdev/zero/internal/config/projectconfig"
)

func testSetup(paramKey, paramValue string) (string, projectconfig.Module) {
	bytes := make([]byte, 15)
	_, _ = rand.Read(bytes)
	name := string(base64.StdEncoding.EncodeToString(bytes[:]))

	_, _ = os.Create(name)

	params := make(projectconfig.Parameters)
	params[paramKey] = paramValue

	mod := projectconfig.Module{
		Parameters: params,
		Files: projectconfig.Files{
			Directory: ".",
			Source:    ".",
		},
	}

	return name, mod
}

func TestPerformIgnoreFileConditionNotMet(t *testing.T) {
	field := "testField"
	value := "trigger"

	filename, mod := testSetup(field, "other value")
	defer os.Remove(filename)

	cond := projectconfig.Condition{
		Action:     "ignoreFile",
		MatchField: field,
		WhenValue:  value,
		Data:       []string{filename},
	}
	condition.Perform(cond, mod)

	_, err := os.Stat(filename)
	if err != nil && !os.IsExist(err) {
		t.Errorf("Expected %v not to be removed\n", filename)
	}
}

func TestPerformIgnoreFileConditionMet(t *testing.T) {
	field := "testField"
	value := "trigger"

	filename, mod := testSetup(field, value)
	defer os.Remove(filename)

	cond := projectconfig.Condition{
		Action:     "ignoreFile",
		MatchField: field,
		WhenValue:  value,
		Data:       []string{filename},
	}
	condition.Perform(cond, mod)

	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		t.Errorf("Expected %v to be removed\n", filename)
	}
}
