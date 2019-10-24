package golang

import (
	"fmt"
	"log"
	"os"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func Generate(templator *templator.Templator, config *config.Commit0Config) {
	GenerateGoMain(templator, config)
	GenerateGoMod(templator, config)
	GenerateHealthServer(templator, config)
	GenerateServers(templator, config)
}

func GenerateGoMain(templator *templator.Templator, config *config.Commit0Config) {
	if _, err := os.Stat("main.go"); os.IsNotExist(err) {

		f, err := os.Create("main.go")

		if err != nil {
			log.Printf("Error: %v", err)
		}

		templator.Go.GoMain.Execute(f, config)
	} else {
		log.Printf("main.go already exists. skipping.")
	}
}

func GenerateGoMod(templator *templator.Templator, config *config.Commit0Config) {
	f, err := os.Create("go.mod")

	if err != nil {
		log.Printf("Error: %v", err)
	}

	templator.Go.GoMod.Execute(f, config)
}

func GenerateServers(templator *templator.Templator, config *config.Commit0Config) {
	serverDirPath := "server"
	err := util.CreateDirIfDoesNotExist(serverDirPath)
	if err != nil {
		log.Printf("Error creating server path: %v", err)
	}

	for _, s := range config.Services {
		serverLibPath := fmt.Sprintf("%s/%s", serverDirPath, s.Name)
		err := os.Mkdir(serverLibPath, os.ModePerm)
		if os.IsExist(err) {
			log.Printf("%s server exists skipping.", s.Name)
			continue
		}
		log.Printf("generating %s", s.Name)
		if err != nil {
			log.Printf("Error generating server: %v", err)
		}

		serverFilePath := fmt.Sprintf("%s/%s.go", serverLibPath, s.Name)
		f, err := os.Create(serverFilePath)

		if err != nil {
			log.Printf("Error: %v", err)
		}

		data := map[string]string{
			"ProjectName": config.Name,
			"ServiceName": s.Name,
			"GitRepo":     config.GitRepo,
		}

		templator.Go.GoServer.Execute(f, data)
	}

}

func GenerateHealthServer(templator *templator.Templator, config *config.Commit0Config) {
	serverDirPath := "server"
	err := util.CreateDirIfDoesNotExist(serverDirPath)
	if err != nil {
		log.Printf("Error creating server path: %v", err)
	}

	serverLibPath := fmt.Sprintf("%s/%s", serverDirPath, "health")
	err = util.CreateDirIfDoesNotExist(serverLibPath)
	if err != nil {
		log.Printf("Error generating server: %v", err)
	}

	serverFilePath := fmt.Sprintf("%s/%s.go", serverLibPath, "health")
	f, err := os.Create(serverFilePath)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	templator.Go.GoHealthServer.Execute(f, config)
}
