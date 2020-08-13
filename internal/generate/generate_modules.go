package generate

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/internal/module"
	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/commitdev/zero/pkg/util/fs"

	"github.com/gabriel-vasile/mimetype"
)

// Generate accepts a projectconfig struct and renders the templates for all referenced modules
func Generate(projectConfig projectconfig.ZeroProjectConfig) error {
	flog.Infof(":clock: Fetching Modules")

	// Make sure module sources are on disk
	wg := sync.WaitGroup{}
	wg.Add(len(projectConfig.Modules))
	for _, mod := range projectConfig.Modules {
		go module.FetchModule(mod.Files.Source, &wg)
	}
	wg.Wait()

	flog.Infof(":memo: Rendering Modules")
	for _, mod := range projectConfig.Modules {
		// Load module configuration
		moduleConfig, err := module.ParseModuleConfig(mod.Files.Source)
		if err != nil {
			return fmt.Errorf("unable to load module:  %v", err)
		}

		moduleDir := path.Join(module.GetSourceDir(mod.Files.Source), moduleConfig.InputDir)
		delimiters := moduleConfig.Delimiters
		outputDir := mod.Files.Directory

		// Data that will be passed in to each template
		templateData := struct {
			Name   string
			Params projectconfig.Parameters
		}{
			projectConfig.Name,
			mod.Parameters,
		}

		fileTemplates, binFiles := newTemplates(moduleDir, outputDir, false)

		//FIXME: remove me; debug log;
		// for _, binFile := range binFiles {
		// 	fmt.Printf("%s\n", "Bin files to just copy over:")
		// 	fmt.Printf("%+v\n", binFile)
		// }

		// TODO: function to copy binFile src -> dest.

		executeTemplates(fileTemplates, templateData, delimiters)
	}
	return nil
}

type TemplateConfig struct {
	source      string
	destination string
	isTemplate  bool
}

type BinFileConfig struct {
	source      string
	destination string
	mineType    string
}

// FIXME: refactor function name
// newTemplates walks the module directory to find all to be templated
func newTemplates(moduleDir string, outputDir string, overwrite bool) ([]*TemplateConfig, []*BinFileConfig) {
	templates := []*TemplateConfig{}
	binFile := []*BinFileConfig{}

	paths, err := getAllFilePathsInDirectory(moduleDir)
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		ignoredPaths, _ := regexp.Compile(constants.IgnoredPaths)
		if ignoredPaths.MatchString(path) {
			continue
		}

		// detect the file type
		detectedMIME, err := mimetype.DetectFile(path)
		if err != nil {
			panic(err)
		}

		// detect if the file type is binary
		isBinary := true
		for mime := detectedMIME; mime != nil; mime = mime.Parent() {
			if mime.Is("text/plain") {
				isBinary = false
			}
		}

		if isBinary {
			binFile = append(binFile, &BinFileConfig{
				source:      path,
				destination: fs.ReplacePath(path, moduleDir, outputDir),
				mineType:    detectedMIME.String(),
			})
			continue
		}

		_, file := filepath.Split(path)
		hasTmpltSuffix := strings.HasSuffix(file, constants.TemplateExtn)
		if hasTmpltSuffix {
			file = strings.Replace(file, constants.TemplateExtn, "", -1)
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
	return templates, binFile
}

// getAllFilePathsInDirectory Recursively get all file paths in directory, including sub-directories.
func getAllFilePathsInDirectory(moduleDir string) ([]string, error) {
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

func executeTemplates(templates []*TemplateConfig, data interface{}, delimiters []string) {
	var wg sync.WaitGroup
	leftDelim := delimiters[0]
	rightDelim := delimiters[1]
	if leftDelim == "" {
		leftDelim = "{{"
	}
	if rightDelim == "" {
		rightDelim = "}}"
	}
	// flog.Infof("Templating params:")
	// pp.Println(data)

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
