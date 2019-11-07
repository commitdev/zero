package main

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
)

func CreateProject(projectConfig ProjectConfiguration) string {
	templates := packr.New("templates", "../../templates")
	t := templator.NewTemplator(templates)
	outDir := "./"
	rootDir := path.Join(outDir, projectConfig.ProjectName)
	log.Printf("Creating project %s.", projectConfig.ProjectName)
	err := os.MkdirAll(rootDir, os.ModePerm)

	if os.IsExist(err) {
		log.Fatalf("Directory %v already exists! Error: %v", projectConfig.ProjectName, err)
	} else if err != nil {
		log.Fatalf("Error creating root: %v ", err)
	}
	var wg sync.WaitGroup

	util.TemplateFileIfDoesNotExist(rootDir, util.ApiGeneratedYamlName, t.ApiCommit0, &wg, projectConfig)

	util.TemplateFileIfDoesNotExist(rootDir, ".gitignore", t.GitIgnore, &wg, projectConfig.ProjectName)

	wg.Wait()

	GenerateArtifacts(projectConfig)

	return rootDir
}

func GenerateArtifacts(projectConfig ProjectConfiguration) {

	templates := packr.New("templates", "../templates")
	t := templator.NewTemplator(templates)

	err := os.Chdir(projectConfig.ProjectName) //cd into project
	if (err != nil) {
		panic(err)
	}
	cfg := config.LoadConfig(util.ApiGeneratedYamlName)
	cfg.Print()

	generate.GenerateArtifactsHelper(t, cfg)
}
