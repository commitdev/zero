package terraform

import (
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup) {
	data := templator.GenericTemplateData{*cfg}

	t.Terraform.TemplateFiles(data, false, wg)
}
