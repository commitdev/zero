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
var registryFilePath string

func init() {
	initCmd.PersistentFlags().StringVarP(&localModulePath, "local-module-path", "m", "github.com/commitdev", "local module path - for using local modules instead of downloading from github")
	initCmd.PersistentFlags().StringVarP(&registryFilePath, "registry-file-path", "r", "https://raw.githubusercontent.com/commitdev/zero/main/registry.yaml", "registry file path - for using a custom list of stacks")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create new project with provided name and initialize configuration based on user input.",
	Run: func(cmd *cobra.Command, args []string) {
		flog.Debugf("Root directory is %s", projectconfig.RootDir)
		projectContext := initPrompts.Init(projectconfig.RootDir, localModulePath, registryFilePath)
		projectConfigErr := projectconfig.CreateProjectConfigFile(projectconfig.RootDir, projectContext.Name, projectContext)

		if projectConfigErr != nil {
			exit.Fatal(fmt.Sprintf(" Init failed while creating the zero project config file %s", projectConfigErr.Error()))
		} else {
			flog.Infof(`:tada: Done - Your project definition file has been initialized with your choices. Please review it, make any required changes and then create your project.
cd %s
cat zero-project.yml
zero create`, projectContext.Name)
		}
	},
}
