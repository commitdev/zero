// This module is invoked when we do template rendering during "zero create."
//
// Each module can have a "conditions" section in their zero-module.yml that
// specifies a condition in the form:
//
//   conditions:
//     - action: ignoreFile
//       matchField: <the name of a parameter in zero-module.yml>
//       whenValue: <value for the matchField that triggers this condition>
//       data:
//       - <arbitrary string>
//
// The structure for this is defined in:
// internal/config/projectconfig/project_config.go.
// The definition is in that file simply to avoid cyclic dependencies; but
// The logic for each type of condition exists here.
//
// See: internal/generate/generate_modules.go
// See: internal/config/projectconfig/project_config.go
//
package condition

import (
	"os"
	"path"

	"github.com/commitdev/zero/internal/config/projectconfig"
)

// Function dispatch for any kind of condition.
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

// Excludes paths from template rendering.
// This occurs after-the-fact. That is, we render all templates to disk, then
// use 'paths' to determine which files and directories to remove from disk.
//
func ignoreFile(paths []string, mod projectconfig.Module) {
	for _, file := range paths {
		filepath := path.Join(mod.Files.Directory, file)
		os.RemoveAll(filepath)
	}
}
