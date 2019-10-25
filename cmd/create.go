package cmd

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func Create(projectName string, outDir string, t *templator.Templator) string {
	rootDir := path.Join(outDir, projectName)
	log.Printf("Creating project %s.", projectName)
	err := os.MkdirAll(rootDir, os.ModePerm)

	if os.IsExist(err) {
		log.Fatalf("Directory %v already exists! Error: %v", projectName, err)
	} else if err != nil {
		log.Fatalf("Error creating root: %v ", err)
	}
	var wg sync.WaitGroup

	util.TemplateFileIfDoesNotExist(rootDir, "commit0.yml", t.Commit0, wg, projectName)
	util.TemplateFileIfDoesNotExist(rootDir, ".gitignore", t.GitIgnore, wg, projectName)

	wg.Wait()
	return rootDir
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project with provided name.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Project name cannot be empty!")
		}

		templates := packr.New("templates", "../templates")
		t := templator.NewTemplator(templates)

		projectName := args[0]

		Create(projectName, "./", t)
	},
}
