package generate

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/golang"
	"github.com/commitdev/commit0/internal/generate/kubernetes"
	"github.com/commitdev/commit0/internal/generate/proto"
	"github.com/commitdev/commit0/internal/generate/react"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
	"log"
	"sync"
)

func GenerateArtifactsHelper(t *templator.Templator, cfg *config.Commit0Config, pathPrefix string) {
	var wg sync.WaitGroup
	if !util.ValidateLanguage(cfg.Frontend.Framework) {
		log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: '%s' is not a supported framework.", cfg.Frontend.Framework)))
	}

	for _, s := range cfg.Services {
		if !util.ValidateLanguage(cfg.Frontend.Framework) {
			log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: '%s' in service '%s' is not a supported language.", s.Name, s.Language)))
		}
	}

	for _, s := range cfg.Services {
		switch s.Language {
		case util.Go:
			log.Println(aurora.Cyan(emoji.Sprintf("Creating Go service")))
			proto.Generate(t, cfg, s, &wg, pathPrefix)
			golang.Generate(t, cfg, s, &wg, pathPrefix)
		}
	}

	if cfg.Infrastructure.AWS.EKS.ClusterName != "" {
		log.Println(aurora.Cyan(emoji.Sprintf("Generating Terraform")))
		kubernetes.Generate(t, cfg, &wg, pathPrefix)
	}

	// @TODO : This strucuture probably needs to be adjusted. Probably too generic.
	switch cfg.Frontend.Framework {
	case util.React:
		log.Println(aurora.Cyan(emoji.Sprintf("Creating React frontend")))
		react.Generate(t, cfg, &wg, pathPrefix)
	}

	util.TemplateFileIfDoesNotExist(pathPrefix, "README.md", t.Readme, &wg, templator.GenericTemplateData{*cfg})

	// Wait for all the templates to be generated
	wg.Wait()

	log.Println("Executing commands")
	// @TODO : Move this stuff to another command? Or genericize it a bit.
	if cfg.Infrastructure.AWS.EKS.Deploy {
		kubernetes.Execute(cfg, pathPrefix)
	}
}
