package ci_test

import (
	"sync"
	"testing"
	"text/template"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/templator"
)

func TestGenerateInvalidCISystem(t *testing.T) {
	testConf := &config.Commit0Config{}
	testCI := config.CI{
		System: "invalidCISystem",
	}

	testTemp := &templator.CITemplator{
		Jenkins:  &template.Template{},
		CircleCI: &template.Template{},
		TravisCI: &template.Template{},
	}

	var wg sync.WaitGroup
	err := ci.Generate(testTemp, testConf, testCI, "/dev/null", &wg)
	if err == nil {
		t.Errorf("Error should be thrown with invalid ci system specified. %s", err.Error())
	}
}
