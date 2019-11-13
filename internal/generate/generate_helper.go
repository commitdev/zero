package generate

import (
	"log"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/golang"
	"github.com/commitdev/commit0/internal/generate/kubernetes"
	"github.com/commitdev/commit0/internal/generate/proto"
	"github.com/commitdev/commit0/internal/generate/react"
	"github.com/commitdev/commit0/internal/generate/terraform"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

func GenerateArtifactsHelper(t *templator.Templator, cfg *config.Commit0Config, pathPrefix string, runInit bool, runApply bool) {
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

	log.Println(aurora.Cyan(emoji.Sprintf("Generating Terraform")))
	terraform.Generate(t, cfg, &wg, pathPrefix)

	if cfg.Infrastructure.AWS.EKS.ClusterName != "" {
		log.Println(aurora.Cyan(emoji.Sprintf("Generating Kubernetes Configuration")))
		kubernetes.Generate(t, cfg, &wg, pathPrefix)
	}

	util.TemplateFileIfDoesNotExist(pathPrefix, "README.md", t.Readme, &wg, templator.GenericTemplateData{*cfg})

	// Wait for all the templates to be generated
	wg.Wait()

	log.Println(aurora.Cyan(emoji.Sprintf("Initializing Infrastructure")))
	if cfg.Infrastructure.AWS.EKS.ClusterName != "" && runInit {
		terraform.Init(cfg, pathPrefix)
	}

	log.Println(aurora.Cyan(emoji.Sprintf("Creating Infrastructure")))
	if cfg.Infrastructure.AWS.EKS.ClusterName != "" && runApply {
		terraform.Execute(cfg, pathPrefix)
		kubernetes.Execute(cfg, pathPrefix)
	}

	// @TODO : This strucuture probably needs to be adjusted. Probably too generic.
	switch cfg.Frontend.Framework {
	case util.React:
		if cfg.Infrastructure.AWS.Cognito.Enabled && cfg.Frontend.Env.CognitoPoolID != "" {
			log.Println(aurora.Cyan(emoji.Sprintf("Creating React frontend")))
			react.Generate(t, cfg, &wg, pathPrefix)
		} else {
			log.Println(aurora.Yellow(emoji.Sprintf(":warning: Missing React environment variables, skipping generation")))
		}
	}

}
