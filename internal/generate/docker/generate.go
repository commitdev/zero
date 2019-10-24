package docker

import (
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func GenerateGoAppDockerFile(templator *templator.Templator, config *config.Commit0Config, wg sync.WaitGroup) {
	util.TemplateFileIfDoesNotExist("docker/app", "Dockerfile", templator.Docker.ApplicationDocker, wg, config)
}

func GenerateGoHTTPGWDockerFile(templator *templator.Templator, config *config.Commit0Config, wg sync.WaitGroup) {
	util.TemplateFileIfDoesNotExist("docker/http", "Dockerfile", templator.Docker.HttpGatewayDocker, wg, config)
}

func GenerateGoDockerCompose(templator *templator.Templator, config *config.Commit0Config, wg sync.WaitGroup) {
	util.TemplateFileIfDoesNotExist("", "docker-compose.yml", templator.Docker.DockerCompose, wg, config)
}
