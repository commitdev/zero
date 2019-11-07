package cmd

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var configPath string

func init() {

	generateCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "commit0.yml", "config path")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate idl & application folders",
	Run: func(cmd *cobra.Command, args []string) {

		templates := packr.New("templates", "../templates")
		t := templator.NewTemplator(templates)

		cfg := config.LoadConfig(configPath)
		cfg.Print()

		generate.GenerateArtifactsHelper(t, cfg)

	},
}
