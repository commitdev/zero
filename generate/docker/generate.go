package docker

import (
	"github.com/commitdev/commit0/util"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/templator"
)

func GenerateGoAppDockerFile(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileIfDoesNotExist("docker/app", "Dockerfile", templator.Docker.ApplicationDocker, config)
}

func GenerateGoHttpGWDockerFile(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileIfDoesNotExist("docker/http", "Dockerfile", templator.Docker.HttpGatewayDocker, config)
}
