package docker

import (
	"github.com/commitdev/sprout/util"

	"github.com/commitdev/sprout/config"
	"github.com/commitdev/sprout/templator"
)

func GenerateGoAppDockerFile(templator *templator.Templator, config *config.SproutConfig) {
	util.TemplateFileIfDoesNotExist("docker/app", "Dockerfile", templator.Docker.ApplicationDocker, config)
	util.TemplateFileIfDoesNotExist("./", ".dockerignore", templator.Docker.DockerIgnore, config)

}

func GenerateGoHttpGWDockerFile(templator *templator.Templator, config *config.SproutConfig) {
	util.TemplateFileIfDoesNotExist("docker/http", "Dockerfile", templator.Docker.HttpGatewayDocker, config)
}
