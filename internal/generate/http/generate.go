package http

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func GenerateHttpGW(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileAndOverwrite("http", "main.go", templator.Go.GoHttpGW, config)
}
