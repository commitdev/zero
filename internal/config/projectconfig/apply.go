package projectconfig

import (
	"fmt"
	"os/exec"
	"path"
	"strings"

	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/flog"
)

// Apply will bootstrap the runtime environment for the project
func Apply(dir string, projectContext *ZeroProjectConfig, applyEnvironments []string) []string {
	flog.Infof(":tada: Bootstrapping project %s. Please use the zero.[hcl, yaml] file to modify the project as needed. %s.", projectContext.Name)

	flog.Infof("Cloud provider: %s", "AWS") // will this come from the config?

	flog.Infof("Runtime platform: %s", "Kubernetes")

	flog.Infof("Infrastructure executor: %s", "Terraform")

	// other details...

	return makeAll(dir, projectContext, applyEnvironments)
}

func makeAll(dir string, projectContext *ZeroProjectConfig, applyEnvironments []string) []string {
	environmentArg := fmt.Sprintf("ENVIRONMENT=%s", strings.Join(applyEnvironments, ","))
	envars := []string{environmentArg}

	outputs := []string{}

	for _, module := range projectContext.Modules {
		// @TODO what's the root dir for these modules?
		modulePath := path.Join(dir, projectContext.Name, module.Files.Directory)
		output := util.ExecuteCommandOutput(exec.Command("make", environmentArg), modulePath, envars)
		flog.Infof("%s", output)

		outputs = append(outputs, output)
	}
	return outputs
}
