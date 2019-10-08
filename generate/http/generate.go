package http

import (
	"github.com/commitdev/sprout/util"

	"github.com/commitdev/sprout/config"
	"github.com/commitdev/sprout/templator"
)

func GenerateHttpGW(templator *templator.Templator, config *config.SproutConfig) {
	util.TemplateFileIfDoesNotExist("http", "main.go", templator.Go.GoHttpGW, config)
}
