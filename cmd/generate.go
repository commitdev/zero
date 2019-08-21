package cmd

import (
	"github.com/commitdev/sprout/config"
	"github.com/commitdev/sprout/generate/golang"
	"github.com/spf13/cobra"
)

var configPath string
var outputPath string

func init() {

	generateCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "sprout.yml", "config path")
	generateCmd.PersistentFlags().StringVarP(&outputPath, "output-path", "o", "./", "config path")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate project from config.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(configPath)
		cfg.Print()
		golang.Generate(Templator, cfg, outputPath)
	},
}
