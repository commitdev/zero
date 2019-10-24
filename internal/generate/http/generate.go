package http

import (
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func GenerateHTTPGW(templator *templator.Templator, config *config.Commit0Config, wg sync.WaitGroup) {
	util.TemplateFileAndOverwrite("http", "main.go", templator.Go.GoHTTPGW, wg, config)
}
