package cmd

import (
	"github.com/commitdev/zero/internal/config/projectconfig"
	initPrompts "github.com/commitdev/zero/internal/init"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with provided name and initialize configuration based on user input.",
	Run: func(cmd *cobra.Command, args []string) {
		projectContext := initPrompts.Init(projectconfig.RootDir)
		projectconfig.Init(projectconfig.RootDir, projectContext.Name, projectContext)
	},
}
