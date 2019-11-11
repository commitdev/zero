package terraform

import (
	"path"
	"sync"

	"log"
	"os/exec"
	"path/filepath"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/kubernetes"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"

	"github.com/logrusorgru/aurora"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	data := templator.GenericTemplateData{*cfg}

	t.Terraform.TemplateFiles(data, false, wg, pathPrefix)
}

func Execute(config *config.Commit0Config, pathPrefix string, outputs []string) map[string]string {
	outputsMap := make(map[string]string)

	if config.Infrastructure.AWS.Cognito.Deploy {
		log.Println("Preparing aws environment...")

		awsSecrets := kubernetes.PromptCredentials()
		envars := kubernetes.GetAwsEnvars(awsSecrets)

		pathPrefix = filepath.Join(pathPrefix, "terraform")

		log.Println(aurora.Cyan(":alarm_clock: Applying infrastructure configuration..."))
		kubernetes.ExecuteCmd(exec.Command("terraform", "init"), pathPrefix, envars)
		kubernetes.ExecuteCmd(exec.Command("terraform", "apply"), pathPrefix, envars)

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
	// fmt.Printf("combined out:\n%s\n", string(out))
	return string(out)
}
