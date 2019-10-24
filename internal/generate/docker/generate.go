package docker

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func GenerateGoAppDockerFile(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileIfDoesNotExist("docker/app", "Dockerfile", templator.Docker.ApplicationDocker, config)
}

func GenerateGoHttpGWDockerFile(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileIfDoesNotExist("docker/http", "Dockerfile", templator.Docker.HttpGatewayDocker, config)
}

func GenerateGoDockerCompose(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileIfDoesNotExist("", "docker-compose.yml", templator.Docker.DockerCompose, config)
}
