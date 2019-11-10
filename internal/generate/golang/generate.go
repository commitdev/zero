package golang

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/generate/docker"
	"github.com/commitdev/commit0/internal/generate/http"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, service config.Service, wg *sync.WaitGroup, pathPrefix string) {
	basePath := filepath.Join(pathPrefix, "service", service.Name)
	healthPath := filepath.Join(basePath, "health")

	data := templator.GolangTemplateData{
		*cfg,
		service,
	}

	util.TemplateFileIfDoesNotExist(basePath, "main.go", t.Go.GoMain, wg, data)
	util.TemplateFileIfDoesNotExist(basePath, "go.mod", t.Go.GoMod, wg, data)
	util.TemplateFileIfDoesNotExist(basePath, "server.go", t.Go.GoServer, wg, data)
	util.TemplateFileIfDoesNotExist(healthPath, "health.go", t.Go.GoHealthServer, wg, data)

	file := fmt.Sprintf("%s.go", service.Name)

	util.TemplateFileIfDoesNotExist(basePath, file, t.Go.GoServer, wg, data)

	if service.CI.System != "" {
		ci.Generate(t.CI, cfg, service.CI, basePath, wg)
	}

	docker.GenerateGoAppDockerFile(t, data, basePath, wg)
	docker.GenerateGoDockerCompose(t, data, basePath, wg)

	if service.Network.Http.Enabled {
		http.GenerateGoHTTPGW(t, data, basePath, wg)
		docker.GenerateGoHTTPGWDockerFile(t, data, basePath, wg)
	}

}
