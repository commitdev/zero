package cmd

import (
	"github.com/commitdev/zero/internal/config"
	"github.com/commitdev/zero/internal/context"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with provided name and initialize configuration based on user input.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			exit.Fatal("Project name cannot be empty!")
		}

		projectName := args[0]
		projectContext := context.Init(projectName, config.RootDir)
		config.Init(config.RootDir, projectName, projectContext)
	},
}
