package ci

import (
	"fmt"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/templator"
	"github.com/commitdev/commit0/util"
)

const (
	defaultGoDockerImage = "golang/golang"
	defaultGoVersion     = "1.12"
	defaultBuildCommand  = "make build"
	defaultTestCommand   = "make test"
)

// Generate a CI configuration file based on your language and CI system
func Generate(templator *templator.CITemplator, config *config.Commit0Config) {
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
	}

	switch config.CI.System {
	case "jenkins":
		util.TemplateFileIfDoesNotExist(".", "Jenkinsfile", templator.Jenkins, config)
	case "circleci":
		util.TemplateFileIfDoesNotExist(".circleci", "config.yml", templator.CircleCI, config)
	case "travisci":
		util.TemplateFileIfDoesNotExist(".", ".travis.yml", templator.TravisCI, config)
	}
}
