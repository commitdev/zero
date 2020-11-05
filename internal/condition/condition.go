package condition

/*
NOTE: To avoid cyclic dependencies, the actual struct datatype is not defined here.
It is doubly defined - once in project_config.go and once in module_config.go.
*/

import (
	"os"
	"path"

	"github.com/commitdev/zero/internal/config/projectconfig"
)

func Perform(cond projectconfig.Condition, mod projectconfig.Module) {
	value, found := mod.Parameters[cond.MatchField]

	// Exit if the condition isn't met.
	if !found || value != cond.WhenValue {
		return
	}

	// Okay, the condition was met, let's execute it.
	switch cond.Action {
	case "ignoreFile":
		ignoreFile(cond.Data, mod)
	}
}

func ignoreFile(data []string, mod projectconfig.Module) {
	for _, file := range data {
		filepath := path.Join(mod.Files.Directory, file)
		os.Remove(filepath)
	}
}
