package generate

import (
	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/module"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/pkg/util/exit"
	"github.com/commitdev/commit0/pkg/util/flog"
)

func GenerateModules(cfg *config.GeneratorConfig) {
	var templateModules []*module.TemplateModule

	// Initiate all the modules defined in the config
	for _, moduleConfig := range cfg.Modules {
		mod, err := module.NewTemplateModule(moduleConfig)

		if err != nil {
			exit.Error("module failed to load: %s", err)
		}
		templateModules = append(templateModules, mod)
	}

	// Prompt for module params and execute each of the generator modules
	for _, mod := range templateModules {
		err := mod.PromptParams()
		if err != nil {
			flog.Warnf("module %s: params prompt failed", mod.Source)
		}

		err = Generate(mod)
		if err != nil {
			exit.Error("module %s: %s", mod.Source, err)
		}
	}
}

func Generate(mod *module.TemplateModule) error {
	t := templator.NewDirTemplator(module.GetSourceDir(mod.Source), mod.Config.Template.Delimiters)
	t.ExecuteTemplates(mod.Params, false, mod.Config.Template.Output, module.GetSourceDir(mod.Source))
	return nil
}
