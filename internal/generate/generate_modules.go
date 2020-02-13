package generate

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/module"
	"github.com/commitdev/commit0/internal/util"
	"github.com/commitdev/commit0/pkg/util/exit"
	"github.com/commitdev/commit0/pkg/util/flog"
	"github.com/k0kubun/pp"

	"github.com/commitdev/commit0/configs"
	"github.com/commitdev/commit0/pkg/util/fs"
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

		err = Generate(mod, cfg)
		if err != nil {
			exit.Error("module %s: %s", mod.Source, err)
		}
	}
}

type TemplateParams struct {
	Name    string
	Context map[string]string
	Params  map[string]string
}

func Generate(mod *module.TemplateModule, generatorCfg *config.GeneratorConfig) error {
	moduleDir := module.GetSourceDir(mod.Source)
	delimiters := mod.Config.Template.Delimiters
	overwrite := true // @TODO get from configs
	outputDir := mod.Config.Template.Output

	templateData := TemplateParams{}
	templateData.Name = generatorCfg.Name
	templateData.Context = generatorCfg.Context
	templateData.Params = mod.Params

	fileTmplts := NewTemplates(moduleDir, outputDir, overwrite)

	ExecuteTemplates(fileTmplts, templateData, delimiters)

	return nil
}

type TemplateConfig struct {
	source      string
	destination string
	isTemplate  bool
}

// NewTemplates walks the module directory to find all  to be templated
func NewTemplates(moduleDir string, outputDir string, overwrite bool) []*TemplateConfig {
	templates := []*TemplateConfig{}

	paths, err := GetAllFilePathsInDirectory(moduleDir)
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		ignoredPaths, _ := regexp.Compile(configs.IgnoredPaths)
		if ignoredPaths.MatchString(path) {
			continue
		}

		_, file := filepath.Split(path)
		hasTmpltSuffix := strings.HasSuffix(file, configs.TemplateExtn)
		if hasTmpltSuffix {
			file = strings.Replace(file, configs.TemplateExtn, "", -1)
		}
		outputPath := fs.ReplacePath(path, moduleDir, outputDir)

		if !overwrite {
			if exists, _ := fs.FileExists(outputPath); exists {
				flog.Warnf("%v already exists. skipping.", outputPath)
				continue
			}
		}

		templates = append(templates, &TemplateConfig{
			source:      path,
			destination: outputPath,
			isTemplate:  hasTmpltSuffix,
		})
	}
	return templates
}

// GetAllFilePathsInDirectory Recursively get all file paths in directory, including sub-directories.
func GetAllFilePathsInDirectory(moduleDir string) ([]string, error) {
	var paths []string
	err := filepath.Walk(moduleDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}

func ExecuteTemplates(templates []*TemplateConfig, data interface{}, delimiters []string) {
	var wg sync.WaitGroup
	leftDelim := delimiters[0]
	rightDelim := delimiters[1]
	if leftDelim == "" {
		leftDelim = "{{"
	}
	if rightDelim == "" {
		rightDelim = "}}"
	}
	flog.Infof("Templating params:")
	pp.Println(data)

	for _, tmpltConfig := range templates {
		source := tmpltConfig.source
		dest := tmpltConfig.destination

		outputDirPath, _ := path.Split(dest)
		err := fs.CreateDirs(outputDirPath)
		if err != nil {
			flog.Errorf("Error creating directory '%s': %v", source, err)
		}
		f, err := os.Create(dest)
		if err != nil {
			flog.Errorf("Error initializing file '%s'", err)
		}
		// @TODO if strict mode then only copy file
		name := path.Base(source)
		template, err := template.New(name).Delims(leftDelim, rightDelim).Funcs(util.FuncMap).ParseFiles(source)
		err = template.Execute(f, data)

		if err != nil {
			flog.Errorf("Error templating '%s': %v", source, err)
		} else {
			flog.Successf("Finished templating : %s", dest)
		}
	}

	wg.Wait()
}
