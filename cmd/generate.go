package cmd

import (
	"log"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/generate/docker"
	"github.com/commitdev/commit0/internal/generate/golang"
	"github.com/commitdev/commit0/internal/generate/http"
	"github.com/commitdev/commit0/internal/generate/kubernetes"
	"github.com/commitdev/commit0/internal/generate/proto"
	"github.com/commitdev/commit0/internal/generate/react"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var configPath string
var language string

const (
	Go         = "go"
	React      = "react"
	Kubernetes = "kubernetes"
)

var supportedLanguages = [...]string{Go, React, Kubernetes}

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

		var wg sync.WaitGroup
		switch language {
		case Go:
			proto.Generate(t, cfg, &wg)
			golang.Generate(t, cfg, &wg)

			docker.GenerateGoAppDockerFile(t, cfg, &wg)
			docker.GenerateGoDockerCompose(t, cfg, &wg)
		case React:
			react.Generate(t, cfg, &wg)
		case Kubernetes:
			kubernetes.Generate(t, cfg, &wg)
		}

		util.TemplateFileIfDoesNotExist("", "README.md", t.Readme, &wg, cfg)

		if cfg.CI.System != "" {
			ci.Generate(t.CI, cfg, ".", &wg)
		}

		if cfg.Network.Http.Enabled {
			http.GenerateHTTPGW(t, cfg, &wg)
			docker.GenerateGoHTTPGWDockerFile(t, cfg, &wg)
		}

		// Wait for all the templates to be generated
		wg.Wait()

		switch language {
		case Kubernetes:
			kubernetes.Execute(cfg)
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
