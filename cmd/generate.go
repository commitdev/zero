package cmd

import (
	"log"
	"os"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/docker"
	"github.com/commitdev/commit0/internal/generate/golang"
	"github.com/commitdev/commit0/internal/generate/http"
	"github.com/commitdev/commit0/internal/generate/proto"
	"github.com/commitdev/commit0/internal/generate/react"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var configPath string
var language string

const (
	Go    = "go"
	React = "react"
)

var supportedLanguages = [...]string{Go, React}

func init() {

	generateCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "commit0.yml", "config path")
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

		templates := packr.New("templates", "../templates")
		t := templator.NewTemplator(templates)

		cfg := config.LoadConfig(configPath)
		cfg.Language = language
		cfg.Print()

		switch language {
		case Go:
			proto.Generate(t, cfg)
			golang.Generate(t, cfg)

			docker.GenerateGoAppDockerFile(t, cfg)
			docker.GenerateGoDockerCompose(t, cfg)
		case React:
			react.Generate(t, cfg)
		}

		f, err := os.Create("README.md")
		if err != nil {
			log.Printf("Error creating commit0 config: %v", err)
		}
		go t.Readme.Execute(f, cfg)

		if cfg.Network.Http.Enabled {
			http.GenerateHTTPGW(t, cfg)
			docker.GenerateGoHTTPGWDockerFile(t, cfg)
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
