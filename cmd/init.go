package cmd

import (
	"fmt"

	"github.com/commitdev/zero/internal/config/projectconfig"
	initPrompts "github.com/commitdev/zero/internal/init"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/spf13/cobra"
)

var localModulePath string

func init() {
	initCmd.PersistentFlags().StringVarP(&localModulePath, "local-module-path", "m", "github.com/commitdev", "local module path - for using local modules instead of downloading from github")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with provided name and initialize configuration based on user input.",
	Run: func(cmd *cobra.Command, args []string) {
		flog.Debugf("Root directory is %s", projectconfig.RootDir)
		projectContext := initPrompts.Init(projectconfig.RootDir, localModulePath)
		projectConfigErr := projectconfig.CreateProjectConfigFile(projectconfig.RootDir, projectContext.Name, projectContext)

		if projectConfigErr != nil {
			exit.Fatal(fmt.Sprintf(" Init failed while creating the zero project config file %s", projectConfigErr.Error()))
		}
	},
}
