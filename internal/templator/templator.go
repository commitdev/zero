package templator

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"text/template"

	"github.com/commitdev/commit0/configs"
	"github.com/commitdev/commit0/internal/util"
	"github.com/commitdev/commit0/pkg/util/exit"
	"github.com/commitdev/commit0/pkg/util/flog"
	"github.com/commitdev/commit0/pkg/util/fs"
)

type DirectoryTemplator struct {
	Templates []*template.Template
}

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
			exit.Fatal("Failed to initialize template %s", path)
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

func (d *DirectoryTemplator) ExecuteTemplates(data interface{}, overwrite bool, pathPrefix string, sourcePath string) {
	var wg sync.WaitGroup

	for _, template := range d.Templates {
		templatePath := template.Name()
		_, file := filepath.Split(templatePath)
		if strings.HasSuffix(file, configs.TemplateExtn) {
			file = strings.Replace(file, configs.TemplateExtn, "", -1)
		}
		outputPath := fs.ReplacePath(templatePath, sourcePath, pathPrefix)

		if !overwrite {
			if exists, _ := fs.FileExists(outputPath); exists {
				flog.Warnf("%v already exists. skipping.", outputPath)
				continue
			}
		}

		outputDir, _ := path.Split(outputPath)
		err := fs.CreateDirs(outputDir)
		if err != nil {
			flog.Errorf("Error creating directory '%s': %v", templatePath, err)
		}
		f, err := os.Create(outputPath)
		if err != nil {
			flog.Errorf("Error initializing file '%s'", err)
		}
		err = template.Execute(f, data)

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
