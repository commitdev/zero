package generate

import (
	"log"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/module"
	"github.com/commitdev/commit0/internal/templator"
)

func GenerateModules(cfg *config.GeneratorConfig) {
	// TODO unit tests: repeatability, GetIdentifier
	// TODO swap default yaml parser with a dedicated configurator loader https://github.com/jinzhu/configor
	// TODO display go-getter progress
	var templateModules []*module.TemplateModule
	for _, moduleConfig := range cfg.Modules {
		mod, err := module.NewTemplateModule(moduleConfig)

		if err != nil {
			log.Panicf("module failed to load: %s", err)
		}
		templateModules = append(templateModules, mod)
	}

	for _, mod := range templateModules {
		err := mod.PromptParams()
		if err != nil {
			log.Panicf("module %s: prompt failed %s", mod.Source, err)
		}

		err = Generate(mod)
		if err != nil {
			log.Panicf("module %s: %s", mod.Source, err)
		}
	}
}

func Generate(mod *module.TemplateModule) error {
	var wg sync.WaitGroup
	t := templator.NewDirTemplator(module.GetSourceDir(mod.Source), mod.Config.Template.Delimiters)
	t.TemplateFiles(mod.Params, false, &wg, "pathPrefix")
	wg.Wait()
	return nil
}
