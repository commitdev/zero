package kubernetes

import (
	"log"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/logrusorgru/aurora"
)

// Generate templates
func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	data := templator.GenericTemplateData{*cfg}
	t.Kubernetes.TemplateFiles(data, false, wg, pathPrefix)
}

// Execute terrafrom init & plan
func Execute(config *config.Commit0Config, pathPrefix string) {
	if config.Infrastructure.AWS.EKS.Deploy {
		envars := util.MakeAwsEnvars(util.GetSecrets())

		pathPrefix = filepath.Join(pathPrefix, "kubernetes/terraform")

		log.Println(aurora.Cyan(":alarm_clock: Applying kubernetes configuration..."))
		util.ExecuteCommand(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "environments/staging/kubernetes"), envars)
		util.ExecuteCommand(exec.Command("terraform", "plan"), filepath.Join(pathPrefix, "environments/staging/kubernetes"), envars)
	}
}
