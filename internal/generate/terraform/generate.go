package terraform

import (
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"sync"

	"log"
	"os/exec"
	"path/filepath"

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

func ExecuteWithOuput(config *config.Commit0Config, pathPrefix string, outputs []string) map[string]string {
	outputsMap := make(map[string]string)

	if config.Infrastructure.AWS.Cognito.Deploy {
		log.Println("Preparing aws environment...")

		envars := util.MakeAwsEnvars(util.GetSecrets())

		pathPrefix = filepath.Join(pathPrefix, "terraform")

		log.Println(aurora.Cyan(":alarm_clock: Applying infrastructure configuration..."))
		util.ExecuteCommand(exec.Command("terraform", "init"), pathPrefix, envars)
		util.ExecuteCommand(exec.Command("terraform", "apply"), pathPrefix, envars)

		for _, output := range outputs {
			outputValue := ExecuteOutput(exec.Command("terraform", "output", output), pathPrefix, envars)
			outputsMap[output] = outputValue
		}

	}

	return outputsMap
}

func ExecuteOutput(cmd *exec.Cmd, pathPrefix string, envars []string) string {
	dir := util.GetCwd()

	cmd.Dir = path.Join(dir, pathPrefix)

	if envars != nil {
		cmd.Env = envars
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Executing terraform command failed: %v\n", err)
	}
	return string(out)
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
