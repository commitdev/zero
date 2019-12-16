package generate

import (
	"log"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/module"
	"github.com/commitdev/commit0/internal/templator"
)

func GenerateModules(cfg *config.GeneratorConfig) {
	// TODO how do you execute commands?
	// TODO unit test: repeatability, GetIdentifier
	// TODO swap default yaml parser with a dedicated configurator loader https://github.com/jinzhu/configor
	// TODO display go-getter progress
	var templateModules []*module.TemplateModule
	for _, moduleConfig := range cfg.Modules {
		module, err := module.NewTemplateModule(moduleConfig)

		if err != nil {
			log.Panicf("module %s: failed to load %s", module.Source, err)
		}
		templateModules = append(templateModules, module)
	}

	for _, module := range templateModules {
		err := module.PromptParams()
		if err != nil {
			log.Panicf("module %s: prompt failed %s", module.Source, err)
		}

		err = Generate(module)
		if err != nil {
			log.Panicf("module %s: %s", module.Source, err)
		}
	}
}

func Generate(module *module.TemplateModule) error {
	var wg sync.WaitGroup
	t := templator.NewDirTemplator(module.GetSourceDir(), module.Config.Template.Delimiters)
	t.TemplateFiles(module.Params, false, &wg, "pathPrefix")
	wg.Wait()
	return nil
}
