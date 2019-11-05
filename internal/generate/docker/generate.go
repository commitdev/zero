package docker

import (
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func GenerateGoAppDockerFile(templator *templator.Templator, data templator.GolangTemplateData, basePath string, wg *sync.WaitGroup) {
	path := filepath.Join(basePath, "docker/app")
	util.TemplateFileIfDoesNotExist(path, "Dockerfile", templator.Docker.ApplicationDocker, wg, data)
}

func GenerateGoHTTPGWDockerFile(templator *templator.Templator, data templator.GolangTemplateData, basePath string, wg *sync.WaitGroup) {
	path := filepath.Join(basePath, "docker/http")
	util.TemplateFileIfDoesNotExist(path, "Dockerfile", templator.Docker.HttpGatewayDocker, wg, data)
}

func GenerateGoDockerCompose(templator *templator.Templator, data templator.GolangTemplateData, basePath string, wg *sync.WaitGroup) {
	util.TemplateFileIfDoesNotExist(basePath, "docker-compose.yml", templator.Docker.DockerCompose, wg, data)
}
