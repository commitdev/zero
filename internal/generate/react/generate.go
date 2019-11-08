package react

import (
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup) {
	data := templator.GenericTemplateData{*cfg}

	t.React.TemplateFiles(data, false, wg)
	if cfg.Frontend.CI.System != "" {
		ci.Generate(t.CI, cfg, cfg.Frontend.CI, "react/", wg)
	}
}
