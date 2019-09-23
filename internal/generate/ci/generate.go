package ci

import (
	"fmt"
	"sync"

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
	config *config.Commit0Config
}

func (e *CIGenerationError) Error() string {
	return fmt.Sprintf("Error: %s. Unable to Generate CI/CD Pipeline with config:\n%v\n", e.err, e.config)
}

// Generate a CI configuration file based on your language and CI system
func Generate(templator *templator.CITemplator, config *config.Commit0Config, basePath string, wg *sync.WaitGroup) error {
	switch config.Language {
	case "go":
		if config.CI.LanguageVersion == "" {
			config.CI.LanguageVersion = defaultGoVersion
		}
		if config.CI.BuildImage == "" {
			config.CI.BuildImage = fmt.Sprintf("%s:%s", defaultGoDockerImage, config.CI.LanguageVersion)
		}
		if config.CI.BuildCommand == "" {
			config.CI.BuildCommand = defaultBuildCommand
		}
		if config.CI.TestCommand == "" {
			config.CI.TestCommand = defaultTestCommand
		}
	default:
		return &CIGenerationError{"Unsupported Language", config}
	}

	var ciConfigPath string
	var ciFilename string

	switch config.CI.System {
	case "jenkins":
		ciConfigPath = basePath
		ciFilename = "Jenkinsfile"
	case "circleci":
		ciConfigPath = fmt.Sprintf("%s/%s", basePath, ".circleci/")
		ciFilename = "config.yml"
	case "travisci":
		ciConfigPath = basePath
		ciFilename = ".travis.yml"
	default:
		return &CIGenerationError{"Unsupported CI System", config}
	}

	util.TemplateFileIfDoesNotExist(ciConfigPath, ciFilename, templator.TravisCI, wg, config)

	return nil
}
