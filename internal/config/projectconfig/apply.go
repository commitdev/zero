package projectconfig

import (
	"fmt"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/flog"
)

// Apply will bootstrap the runtime environment for the project
func Apply(dir string, projectContext *ZeroProjectConfig, applyEnvironments []string) {
	flog.Infof(":tada: Bootstrapping project %s. Please use the zero.[hcl, yaml] file to modify the project as needed. %s.", projectContext.Name)

	flog.Infof("Cloud provider: %s", "AWS") // will this come from the config?

	flog.Infof("Runtime platform: %s", "Kubernetes")

	flog.Infof("Infrastructure executor: %s", "Terraform")

	// other details...

	makeAll(dir, projectContext, applyEnvironments)
}

func makeAll(dir string, projectContext *ZeroProjectConfig, applyEnvironments []string) error {
	environmentArg := fmt.Sprintf("ENVIRONMENT=%s", strings.Join(applyEnvironments, ","))
	envars := []string{environmentArg}

	for _, module := range projectContext.Modules {
		// TODO what's the root dir for these modules?
		// what's the real path to these modules? It's probably not the module name...
		modulePath, err := filepath.Abs(path.Join(dir, projectContext.Name, module))
		if err != nil {
			return err
		}

		// @TODO mock exec?
		output := util.ExecuteCommandOutput(exec.Command("make", environmentArg), modulePath, envars)
		fmt.Println(output)
	}
	return nil
}
