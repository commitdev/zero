package golang

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func Generate(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	util.TemplateFileIfDoesNotExist("", "main.go", templator.Go.GoMain, wg, config)
	util.TemplateFileIfDoesNotExist("", "go.mod", templator.Go.GoMod, wg, config)
	util.TemplateFileIfDoesNotExist("server/health", "health.go", templator.Go.GoHealthServer, wg, config)
	GenerateServers(templator, config, wg)
}

func GenerateServers(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	serverDirPath := "server"
	err := util.CreateDirIfDoesNotExist(serverDirPath)
	if err != nil {
		log.Printf("Error creating server path: %v", err)
	}

	for _, s := range config.Services {
		path := filepath.Join("server", s.Name)
		file := fmt.Sprintf("%s.go", s.Name)

		data := map[string]string{
			"ProjectName": config.Name,
			"ServiceName": s.Name,
			"GitRepo":     config.GitRepo,
		}

		util.TemplateFileIfDoesNotExist(path, file, templator.Go.GoServer, wg, data)
	}

}
