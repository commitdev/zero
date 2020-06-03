package generate

import (
	"log"
	"os/exec"
	"path/filepath"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/credentials"
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/flog"
)

// @TODO: These are specific to a k8s version. If we make the version a config option we will need to change this
var amiLookup = map[string]string{
	"us-east-1":    "ami-0392bafc801b7520f",
	"us-east-2":    "ami-082bb518441d3954c",
	"us-west-2":    "ami-05d586e6f773f6abf",
	"eu-west-1":    "ami-059c6874350e63ca9",
	"eu-central-1": "ami-0e21bc066a9dbabfa",
}

// GetOutputs captures the terraform output for the specific variables
func GetOutputs(cfg *projectconfig.ZeroProjectConfig, pathPrefix string, outputs []string) map[string]string {
	outputsMap := make(map[string]string)
	envars := credentials.MakeAwsEnvars(cfg, project.GetSecrets(util.GetCwd()))
	pathPrefix = filepath.Join(pathPrefix, "environments/staging")

	for _, output := range outputs {
		outputValue := util.ExecuteCommandOutput(exec.Command("terraform", "output", output), pathPrefix, envars)
		outputsMap[output] = outputValue
	}

	return outputsMap
}

// Init sets up anything required by Execute
func Init(cfg *projectconfig.ZeroProjectConfig, pathPrefix string) {
	if cfg.Infrastructure.AWS.AccountID != "" {
		flog.Infof("Preparing aws environment...")

		envars := project.MakeAwsEnvars(cfg, project.GetSecrets(util.GetCwd()))

		pathPrefix = filepath.Join(pathPrefix, "terraform")

		// @TODO : A check here would be nice to see if this stuff exists first, mostly for testing
		flog.Infof(":alarm_clock: Initializing remote backend...")
		util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "bootstrap/remote-state"), envars)
		util.ExecuteCommand(exec.Command("terraform", "apply", "-auto-approve"), filepath.Join(pathPrefix, "bootstrap/remote-state"), envars)

		// flog.Infof("Creating users...")
		// util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "bootstrap/create-users"), envars)
		// util.ExecuteCommand(exec.Command("terraform", "apply", "-auto-approve"), filepath.Join(pathPrefix, "bootstrap/create-users"), envars)
	}
}

// Execute terrafrom init & plan. May modify the config passed in
func Execute(cfg *projectconfig.ZeroProjectConfig, pathPrefix string) {
	if cfg.Infrastructure.AWS.AccountID != "" {
		log.Println("Preparing aws environment...")

		envars := project.MakeAwsEnvars(cfg, project.GetSecrets(util.GetCwd()))

		pathPrefix = filepath.Join(pathPrefix, "terraform")

		flog.Infof(":alarm_clock: Applying infrastructure configuration...")
		util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "environments/staging"), envars)
		util.ExecuteCommand(exec.Command("terraform", "apply", "-auto-approve"), filepath.Join(pathPrefix, "environments/staging"), envars)

		// @TODO get output fields from `mapOutputs` param in configs, can't be hardcoded
		outputs := []string{
			"cognito_pool_id",
			"cognito_client_id",
		}
		outputValues := GetOutputs(cfg, pathPrefix, outputs)
		cfg.Context["cognito_pool_id"] = outputValues["cognito_pool_id"]
		cfg.Context["cognito_client_id"] = outputValues["cognito_client_id"]
	}
}
