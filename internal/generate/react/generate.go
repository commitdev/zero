package react

import (
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	data := templator.GenericTemplateData{*cfg}

	t.React.TemplateFiles(data, false, wg)
	basePath := path.Join(pathPrefix, "react/")
	if cfg.Frontend.CI.System != "" {
		ci.Generate(t.CI, cfg, cfg.Frontend.CI, basePath, wg)
	}
}
