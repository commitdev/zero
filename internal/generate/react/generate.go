package react

import (
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	templator.React.TemplateFiles(config, false, wg)
}
