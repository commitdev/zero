package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project with provided name.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Project name cannot be empty!")
		}

		projectName := args[0]

		rootDir := fmt.Sprintf("./%v", projectName)

		log.Printf("Creating project %s.", projectName)

		err := os.Mkdir(rootDir, os.ModePerm)
		if os.IsExist(err) {
			log.Fatalf("Directory %v already exists! Error: %v", projectName, err)
		} else if err != nil {
			log.Fatalf("Error creating root: %v ", err)
		}

		sproutConfigPath := fmt.Sprintf("%v/sprout.yml", rootDir)

		f, err := os.Create(sproutConfigPath)
		if err != nil {
			log.Printf("Error creating sprout config: %v", err)
		}
		Templator.Sprout.Execute(f, projectName)

		gitIgnorePath := fmt.Sprintf("%v/.gitignore", rootDir)
		f, err = os.Create(gitIgnorePath)
		if err != nil {
			log.Printf("Error creating sprout config: %v", err)
		}
		Templator.GitIgnore.Execute(f, projectName)
	},
}
