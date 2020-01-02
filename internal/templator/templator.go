package templator

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/commitdev/commit0/configs"
	"github.com/commitdev/commit0/internal/util"
	"github.com/commitdev/commit0/pkg/util/flog"
	"github.com/commitdev/commit0/pkg/util/fs"
)

type DirectoryTemplator struct {
	Templates []*template.Template
}

// // @TODO deprecate
// func (d *DirectoryTemplator) TemplateFiles(data interface{}, overwrite bool, wg *sync.WaitGroup, pathPrefix string) {
// 	for _, template := range d.Templates {
// 		templatePath := path.Join(pathPrefix, template.Name())
// 		dir, file := filepath.Split(templatePath)
// 		if strings.HasSuffix(file, ".tmpl") {
// 			file = strings.Replace(file, ".tmpl", "", -1)
// 		}
// 		if overwrite {
// 			util.TemplateFileAndOverwrite(dir, file, template, wg, data)
// 		} else {
// 			util.TemplateFileIfDoesNotExist(dir, file, template, wg, data)
// 		}
// 	}
// }

func NewDirTemplator(moduleDir string, delimiters []string) *DirectoryTemplator {
	templates := []*template.Template{}
	leftDelim := delimiters[0]
	rightDelim := delimiters[1]
	if leftDelim == "" {
		leftDelim = "{{"
	}
	if rightDelim == "" {
		rightDelim = "}}"
	}

	paths, err := GetAllFilePathsInDirectory(moduleDir)
	if err != nil {
		panic(err)
	}

	for _, path := range paths {
		ignoredPaths, _ := regexp.Compile(configs.IgnoredPaths)
		if ignoredPaths.MatchString(path) {
			continue
		}
		template, err := template.New(path).Delims(leftDelim, rightDelim).Funcs(util.FuncMap).ParseFiles(path)
		if err != nil {
			flog.Errorf("Failed to initialize template %s", path)
			panic(err)
		}
		templates = append(templates, template)
	}

	return &DirectoryTemplator{
		Templates: templates,
	}
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

func ExecuteTemplate(templatePath string, outputPath string, data interface{}) error {
	tmplt, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	return tmplt.Execute(f, data)
}

func (d *DirectoryTemplator) ExecuteTemplates(data interface{}, overwrite bool, pathPrefix string) {
	var wg sync.WaitGroup

	for _, template := range d.Templates {
		templatePath := template.Name()
		_, file := filepath.Split(templatePath)
		if strings.HasSuffix(file, ".tmpl") {
			file = strings.Replace(file, ".tmpl", "", -1)
		}
		outputPath := fs.PrependPath(templatePath, pathPrefix)

		if !overwrite {
			if exists, _ := fs.FileExists(outputPath); exists {
				flog.Warnf("%v already exists. skipping.", outputPath)
				continue
			}
		}

		err := fs.CreateDirs(outputPath)
		if err != nil {
			err = ExecuteTemplate(templatePath, outputPath, data)
		}

		if err != nil {
			flog.Errorf("Error templating '%s': %v", templatePath, err)
		} else {
			flog.Successf("Finished templating : %s", outputPath)
		}
	}

	wg.Wait()
}

// func removeTmplDuplicates(keys []string) []string {
// 	filteredKeys := []string{}
// 	for _, key := range keys {
// 		if !containsStr(keys, key+".tmpl") {
// 			filteredKeys = append(filteredKeys, key)
// 		}
// 	}
// 	return filteredKeys
// }

// func containsStr(arr []string, key string) bool {
// 	for _, val := range arr {
// 		if val == key {
// 			return true
// 		}
// 	}
// 	return false
// }
