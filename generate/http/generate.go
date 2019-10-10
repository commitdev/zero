package http

import (
	"github.com/commitdev/commit0/util"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/templator"
)

func GenerateHttpGW(templator *templator.Templator, config *config.Commit0Config) {
	util.TemplateFileAndOverwrite("http", "main.go", templator.Go.GoHttpGW, config)
}
