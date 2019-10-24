package react

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(templator *templator.Templator, config *config.Commit0Config) {
	templator.React.TemplateFiles(config, false)
}
