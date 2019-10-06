package cmd

import (
	"github.com/commitdev/sprout/config"
	"github.com/commitdev/sprout/generate/golang"
	"github.com/commitdev/sprout/generate/proto"
	"log"

	"github.com/spf13/cobra"
)

var configPath string
var language string

const (
	Go = "go"
)

var supportedLanguages = [...]string{Go}

func init() {

	generateCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "sprout.yml", "config path")
	generateCmd.PersistentFlags().StringVarP(&language, "language", "l", "", "language to generate project in")

	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate idl & application folders",
	Run: func(cmd *cobra.Command, args []string) {
		if !ValidLanguage() {
			log.Fatalf("'%s' is not a supported language.", language)
		}

		cfg := config.LoadConfig(configPath)
		cfg.Language = language
		cfg.Print()

		proto.Generate(Templator, cfg)
		switch language {
		case Go:
			golang.Generate(Templator, cfg)

		}
	},
}

func ValidLanguage() bool {
	for _, l := range supportedLanguages {
		if l == language {
			return true
		}
	}

	return false
}
