package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var projectName string

func init() {

	createCmd.PersistentFlags().StringVarP(&projectName, "project-name", "p", "", "project name")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project.",
	Run: func(cmd *cobra.Command, args []string) {

		rootDir := fmt.Sprintf("./%v", projectName)

		err := os.Mkdir(rootDir, os.ModePerm)
		if os.IsExist(err) {
			log.Fatalf("Directory %v already exists! Error: %v", projectName, err)
		} else if err != nil {
			log.Fatalf("Error creating root: %v ", err)
		}

		sproutConfigPath := fmt.Sprintf("%v/sprout.yml", projectName)

		f, err := os.Create(sproutConfigPath)
		if err != nil {
			log.Printf("Error creating sprout config: %v", err)
		}

		Templator.Sprout.Execute(f, projectName)
	},
}
