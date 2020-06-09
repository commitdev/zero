package cmd

import (
	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/context"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with provided name and initialize configuration based on user input.",
	Run: func(cmd *cobra.Command, args []string) {
		projectContext := context.Init(projectconfig.RootDir)
		projectconfig.Init(projectconfig.RootDir, projectContext.Name, projectContext)
	},
}
