package ci

import (
	"fmt"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/templator"
	"github.com/commitdev/commit0/util"
)

func Generate(templator *templator.Templator, config *config.Commit0Config) {
	fmt.Printf("%v\n\n", config.CI.System)
	fmt.Printf("%v\n\n", templator.CI)

	util.TemplateFileIfDoesNotExist(".circleci", "circleci.yml", templator.CI, config)
}
