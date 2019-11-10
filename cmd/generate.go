package cmd

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/golang"
	"github.com/commitdev/commit0/internal/generate/kubernetes"
	"github.com/commitdev/commit0/internal/generate/proto"
	"github.com/commitdev/commit0/internal/generate/react"
	"github.com/commitdev/commit0/internal/generate/terraform"
	"github.com/commitdev/commit0/internal/generate"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
)

var configPath string

func init() {

	generateCmd.PersistentFlags().StringVarP(&configPath, "config", "c", util.CommitYml, "config path")

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

<<<<<<< HEAD
		var wg sync.WaitGroup
		if !ValidLanguage(cfg.Frontend.Framework) {
			log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: '%s' is not a supported framework.", cfg.Frontend.Framework)))
		}

		for _, s := range cfg.Services {
			if !ValidLanguage(cfg.Frontend.Framework) {
				log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: '%s' in service '%s' is not a supported language.", s.Name, s.Language)))
			}
		}

		for _, s := range cfg.Services {
			switch s.Language {
			case Go:
				log.Println(aurora.Cyan(emoji.Sprintf("Creating Go service")))
				proto.Generate(t, cfg, s, &wg)
				golang.Generate(t, cfg, s, &wg)
			}
		}

		if cfg.Infrastructure.AWS.EKS.ClusterName != "" {
			log.Println(aurora.Cyan(emoji.Sprintf("Generating Terraform")))
			kubernetes.Generate(t, cfg, &wg)
		}

		log.Println(aurora.Cyan(emoji.Sprintf("Generating Terraform")))
		terraform.Generate(t, cfg, &wg)

		// @TODO : This strucuture probably needs to be adjusted. Probably too generic.
		switch cfg.Frontend.Framework {
		case React:
			log.Println(aurora.Cyan(emoji.Sprintf("Creating React frontend")))
			react.Generate(t, cfg, &wg)
		}

		util.TemplateFileIfDoesNotExist("", "README.md", t.Readme, &wg, templator.GenericTemplateData{*cfg})

		// Wait for all the templates to be generated
		wg.Wait()

		log.Println("Executing commands")
		// @TODO : Move this stuff to another command? Or genericize it a bit.
		if cfg.Infrastructure.AWS.EKS.Deploy {
			kubernetes.Execute(cfg)
		}
=======
		generate.GenerateArtifactsHelper(t, cfg, "")
>>>>>>> master

	},
}
