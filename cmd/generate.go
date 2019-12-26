package cmd

import (
	"github.com/commitdev/commit0/configs"
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate"
	"github.com/spf13/cobra"
)

var configPath string

// var runInit bool
// var runApply bool

func init() {
	generateCmd.PersistentFlags().StringVarP(&configPath, "config", "c", configs.CommitYml, "config path")
	// generateCmd.PersistentFlags().BoolVarP(&runInit, "init", "", false, "Initialize config after generating")
	// generateCmd.PersistentFlags().BoolVarP(&runApply, "apply", "", false, "Apply config after generating")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate idl & application folders",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadGeneratorConfig(configPath)
		// cfg.Print()

		generate.GenerateModules(cfg)
	},
}
