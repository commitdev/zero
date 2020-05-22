package cmd

import (
	"github.com/commitdev/zero/internal/config"
	"github.com/commitdev/zero/internal/context"
	"github.com/commitdev/zero/pkg/util/exit"
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
			exit.Fatal("Project name cannot be empty!")
		}

		projectName := args[0]
		context.Create(projectName, "./")

		config.CreateExample(projectName)
	},
}
