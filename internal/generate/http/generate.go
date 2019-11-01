package http

import (
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func GenerateGoHTTPGW(templator *templator.Templator, data templator.GolangTemplateData, basePath string, wg *sync.WaitGroup) {
	path := filepath.Join(basePath, "http")
	util.TemplateFileAndOverwrite(path, "main.go", templator.Go.GoHTTPGW, wg, data)
}
