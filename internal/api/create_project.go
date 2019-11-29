package api

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

func createProject(projectConfig util.ProjectConfiguration) string {
	templates := packr.New("templates", "../../templates")
	t := templator.NewTemplator(templates)
	outDir := "./"
	rootDir := path.Join(outDir, projectConfig.ProjectName)
	log.Printf("Creating project %s.", projectConfig)
	err := os.MkdirAll(rootDir, os.ModePerm)

	if os.IsExist(err) {
		log.Fatalf("Directory %v already exists! Error: %v", projectConfig.ProjectName, err)
	} else if err != nil {
		log.Fatalf("Error creating root: %v ", err)
	}
	var wg sync.WaitGroup

	util.TemplateFileIfDoesNotExist(rootDir, util.CommitYml, t.Commit0, &wg, projectConfig)

	util.TemplateFileIfDoesNotExist(rootDir, ".gitignore", t.GitIgnore, &wg, projectConfig.ProjectName)

	wg.Wait()

	GenerateArtifacts(projectConfig)

	return rootDir
}

func GenerateArtifacts(projectConfig util.ProjectConfiguration) {

	templates := packr.New("templates", "../templates")
	t := templator.NewTemplator(templates)

	generatedYml := path.Join(projectConfig.ProjectName, util.CommitYml)

	cfg := config.LoadConfig(generatedYml)
	cfg.Print()

	generate.GenerateArtifactsHelper(t, cfg, projectConfig.ProjectName, false, false)
}
