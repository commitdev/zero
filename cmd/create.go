package cmd

import (
	"fmt"
	"path"
	"strings"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/generate"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/spf13/cobra"
)

var createConfigPath string

func init() {
	createCmd.PersistentFlags().StringVarP(&createConfigPath, "config", "c", constants.ZeroProjectYml, "config path")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: fmt.Sprintf("Create projects for modules and configuration specified in %s", constants.ZeroProjectYml),
	Run: func(cmd *cobra.Command, args []string) {
		Create(projectconfig.RootDir, createConfigPath)
	},
}

func Create(dir string, createConfigPath string) {
	if strings.Trim(createConfigPath, " ") == "" {
		exit.Fatal("config path cannot be empty!")
	}
	configFilePath := path.Join(dir, createConfigPath)
	projectConfig := projectconfig.LoadConfig(configFilePath)

	generate.Generate(*projectConfig)
}
