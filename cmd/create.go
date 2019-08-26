package cmd

import (
	"log"
	"os"
	"fmt"
	

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

		err := os.Mkdir(projectName, os.ModePerm)
		if os.IsExist(err){
			log.Fatalf("Directory %v already exists!", projectName)
		}

		sproutConfigPath := fmt.Sprintf("%v/sprout.yml", projectName)

		f, err := os.Create(sproutConfigPath)
		if err != nil {
			log.Printf("Error creating sprout config: %v", err)
		}

		Templator.Sprout.Execute(f, projectName)
	},
}
