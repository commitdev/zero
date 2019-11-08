package cmd

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func Create(projectName string, outDir string, t *templator.Templator) string {
	rootDir := path.Join(outDir, projectName)
	log.Println(aurora.Cyan(emoji.Sprintf(":tada: Creating project %s.", projectName)))
	err := os.MkdirAll(rootDir, os.ModePerm)

	if os.IsExist(err) {
		log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Directory %v already exists! Error: %v", projectName, err)))
	} else if err != nil {
		log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Error creating root: %v ", err)))
	}
	var wg sync.WaitGroup

	defaultProjConfig := defaultProjConfig(projectName)

	util.TemplateFileIfDoesNotExist(rootDir, util.CommitYml, t.Commit0, &wg, defaultProjConfig)
	util.TemplateFileIfDoesNotExist(rootDir, ".gitignore", t.GitIgnore, &wg, projectName)

	wg.Wait()
	return rootDir
}

func defaultProjConfig(projectName string) util.ProjectConfiguration {
	return util.ProjectConfiguration{
		ProjectName:       projectName,
		FrontendFramework: "react",
		Organization:      "mycompany",
		Description:       "",
		Maintainers: []util.Maintainer{{
			Name:  "bob",
			Email: "bob@test.com",
		}},
		Services: []util.Service{{
			Name:  "User",
			Description: "User Service",
			Language: "go",
			GitRepo: "github.com/test/repo",
		}},
	}
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project with provided name.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Project name cannot be empty!")))
		}

		templates := packr.New("templates", "../templates")
		t := templator.NewTemplator(templates)

		projectName := args[0]

		Create(projectName, "./", t)
	},
}
