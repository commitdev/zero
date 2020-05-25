package cmd

import (
	"github.com/commitdev/zero/configs"
	"github.com/commitdev/zero/internal/config"
	"github.com/commitdev/zero/internal/generate"
	"github.com/spf13/cobra"
)

var createConfigPath string

func init() {
	createCmd.PersistentFlags().StringVarP(&createConfigPath, "config", "c", configs.CommitYml, "config path")

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create directories and render templates based on included modules.",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadGeneratorConfig(createConfigPath)

		generate.GenerateModules(cfg)
	},
}
