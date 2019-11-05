package ci

import (
	"fmt"
	"sync"
	"text/template"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

const (
	defaultGoDockerImage = "golang/golang"
	defaultGoVersion     = "1.12"
	defaultBuildCommand  = "make build"
	defaultTestCommand   = "make test"
)

type CIGenerationError struct {
	err    string
	config config.CI
}

func (e *CIGenerationError) Error() string {
	return fmt.Sprintf("Error: %s. Unable to Generate CI/CD Pipeline with config:\n%v\n", e.err, e.config)
}

// Generate a CI configuration file based on your language and CI system
func Generate(t *templator.CITemplator, cfg *config.Commit0Config, ciConfig config.CI, basePath string, wg *sync.WaitGroup) error {

	var ciConfigPath string
	var ciFilename string
	var ciTemp *template.Template

	switch ciConfig.System {
	case "jenkins":
		ciConfigPath = basePath
		ciFilename = "Jenkinsfile"
		ciTemp = t.Jenkins
	case "circleci":
		ciConfigPath = fmt.Sprintf("%s/%s", basePath, ".circleci/")
		ciFilename = "config.yml"
		ciTemp = t.CircleCI
	case "travisci":
		ciConfigPath = basePath
		ciFilename = ".travis.yml"
		ciTemp = t.TravisCI
	default:
		return &CIGenerationError{"Unsupported CI System", ciConfig}
	}

	data := templator.CITemplateData{
		*cfg,
		ciConfig,
	}
	util.TemplateFileIfDoesNotExist(ciConfigPath, ciFilename, ciTemp, wg, data)

	return nil
}
