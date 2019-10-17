package kubernetes

import (
	//"github.com/commitdev/commit0/util"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/templator"
)

func Generate(templator *templator.Templator, config *config.Commit0Config) {
	templator.Kubernetes.TemplateFiles(config, false)
}

