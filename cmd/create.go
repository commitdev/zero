package cmd

import (
	"log"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(createCmd)
}

func Create(projectName string, outDir string) string {
	rootDir := path.Join(outDir, projectName)
	log.Printf("Creating project %s.", projectName)
	err := os.MkdirAll(rootDir, os.ModePerm)

	if os.IsExist(err) {
		log.Fatalf("Directory %v already exists! Error: %v", projectName, err)
	} else if err != nil {
		log.Fatalf("Error creating root: %v ", err)
	}

	commit0ConfigPath := path.Join(rootDir, "commit0.yml")
	log.Printf("%s", commit0ConfigPath)

	f, err := os.Create(commit0ConfigPath)
	if err != nil {
		log.Printf("Error creating commit0 config: %v", err)
	}
	Templator.Commit0.Execute(f, projectName)

	gitIgnorePath := path.Join(rootDir, ".gitignore")
	f, err = os.Create(gitIgnorePath)
	if err != nil {
		log.Printf("Error creating commit0 config: %v", err)
	}
	Templator.GitIgnore.Execute(f, projectName)

	return rootDir
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project with provided name.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Project name cannot be empty!")
		}

		projectName := args[0]

		Create(projectName, "./")
	},
}
