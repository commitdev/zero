package main

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/generate/docker"
	"github.com/commitdev/commit0/internal/generate/golang"
	"github.com/commitdev/commit0/internal/generate/http"
	"github.com/commitdev/commit0/internal/generate/kubernetes"
	"github.com/commitdev/commit0/internal/generate/proto"
	"github.com/commitdev/commit0/internal/generate/react"
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

	GenerateArtifacts(projectConfig.ProjectName, projectConfig.Language)

	return rootDir
}

func GenerateArtifacts(projectName, language string) {
	if !util.ValidateLanguage(language) {
		log.Fatalf("'%s' is not a supported language.", language)
	}

	templates := packr.New("templates", "../templates")
	t := templator.NewTemplator(templates)

	err := os.Chdir(projectName) //cd into project
	if (err != nil) {
		panic(err)
	}

	cfg := config.LoadConfig(util.ApiGeneratedYamlName)
	cfg.Language = language
	cfg.Print()

	var wg sync.WaitGroup
	switch language {
	case util.Go:
		proto.Generate(t, cfg, &wg)
		golang.Generate(t, cfg, &wg)

		docker.GenerateGoAppDockerFile(t, cfg, &wg)
		docker.GenerateGoDockerCompose(t, cfg, &wg)
	case util.React:
		react.Generate(t, cfg, &wg)
	case util.Kubernetes:
		kubernetes.Generate(t, cfg, &wg)
	}

	util.TemplateFileIfDoesNotExist("", "README.md", t.Readme, &wg, cfg)

	if cfg.CI.System != "" {
		ci.Generate(t.CI, cfg, ".", &wg)
	}

	if cfg.Network.Http.Enabled {
		http.GenerateHTTPGW(t, cfg, &wg)
		docker.GenerateGoHTTPGWDockerFile(t, cfg, &wg)
	}

	// Wait for all the templates to be generated
	wg.Wait()

	switch language {
	case util.Kubernetes:
		kubernetes.Execute(cfg)
	}
}
