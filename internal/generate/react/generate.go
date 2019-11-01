package react

import (
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	templator.React.TemplateFiles(config, false, wg)
	if config.Frontend.CI.System != "" {
		ci.Generate(templator.CI, config, config.Frontend.CI, "react/", wg)
	}
}
