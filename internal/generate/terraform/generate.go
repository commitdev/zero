package terraform

import (
	"log"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"

	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

// @TODO : These are specific to a k8s version. If we make the version a config option we will need to change this
var amiLookup = map[string]string{
	"us-east-1":    "ami-0392bafc801b7520f",
	"us-east-2":    "ami-082bb518441d3954c",
	"us-west-2":    "ami-05d586e6f773f6abf",
	"eu-west-1":    "ami-059c6874350e63ca9",
	"eu-central-1": "ami-0e21bc066a9dbabfa",
}

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	if cfg.Infrastructure.AWS.EKS.WorkerAMI == "" {
		ami, found := amiLookup[cfg.Infrastructure.AWS.Region]
		if !found {
			log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Unable to look up an AMI for the chosen region")))
		}

		cfg.Infrastructure.AWS.EKS.WorkerAMI = ami
	}
	data := templator.GenericTemplateData{Config: *cfg}

	t.Terraform.TemplateFiles(data, false, wg, pathPrefix)
}

func GetOutputs(config *config.Commit0Config, pathPrefix string, outputs []string) map[string]string {
	outputsMap := make(map[string]string)

	log.Println("Preparing aws environment...")

	envars := util.MakeAwsEnvars(util.GetSecrets())

	path := filepath.Join(pathPrefix, "terraform")

	for _, output := range outputs {
		outputValue := util.ExecuteCommandOutput(exec.Command("terraform", "output", output), path, envars)
		outputsMap[output] = outputValue
	}

	return outputsMap
}

// Execute terrafrom init & plan
func Execute(config *config.Commit0Config, pathPrefix string) {
	// @TODO : Change this check. Most likely we should discover the accountid
	if config.Infrastructure.AWS.AccountId != "" {
		log.Println("Preparing aws environment...")

		envars := util.MakeAwsEnvars(util.GetSecrets())

		pathPrefix = filepath.Join(pathPrefix, "terraform")

		// @TODO : A check here would be nice to see if this stuff exists first, mostly for testing
		log.Println(aurora.Cyan(emoji.Sprintf(":alarm_clock: Initializing remote backend...")))
		util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "bootstrap/remote-state"), envars)
		util.ExecuteCommand(exec.Command("terraform", "apply", "-auto-approve"), filepath.Join(pathPrefix, "bootstrap/remote-state"), envars)

		log.Println(aurora.Cyan(":alarm_clock: Planning infrastructure..."))
		util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "environments/staging"), envars)
		util.ExecuteCommand(exec.Command("terraform", "plan"), filepath.Join(pathPrefix, "environments/staging"), envars)

		log.Println(aurora.Cyan(":alarm_clock: Applying infrastructure configuration..."))
		util.ExecuteCommand(exec.Command("terraform", "apply"), filepath.Join(pathPrefix, "environments/staging"), envars)

		log.Println(aurora.Cyan(":alarm_clock: Applying kubernetes configuration..."))
		util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "environments/staging/kubernetes"), envars)
		util.ExecuteCommand(exec.Command("terraform", "plan"), filepath.Join(pathPrefix, "environments/staging/kubernetes"), envars)
	}
}
