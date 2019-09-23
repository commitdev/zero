package ci_test

import (
	"sync"
	"testing"
	"text/template"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/templator"
)

func TestGenerateJenkins(t *testing.T) {
	testConf := &config.Commit0Config{
		Language: "go",
		CI: config.CI{
			System: "jenkins",
		},
	}
	testTemp := &templator.CITemplator{
		Jenkins:  &template.Template{},
		CircleCI: &template.Template{},
		TravisCI: &template.Template{},
	}
	var wg sync.WaitGroup
	err := ci.Generate(testTemp, testConf, "/dev/null", &wg)
	if err != nil {
		t.Errorf("Error when executing test. %s", err)
	}

	expectedBuildImage := "golang/golang:1.12"
	actualBuildImage := testConf.CI.BuildImage
	if actualBuildImage != expectedBuildImage {
		t.Errorf("want: %s, got: %s", expectedBuildImage, actualBuildImage)
	}

	expectedBuildCommand := "make build"
	actualBuildCommand := testConf.CI.BuildCommand
	if actualBuildCommand != expectedBuildCommand {
		t.Errorf("want: %s, got: %s", expectedBuildCommand, actualBuildCommand)
	}

	expectedTestCommand := "make test"
	actualTestCommand := testConf.CI.TestCommand
	if actualTestCommand != expectedTestCommand {
		t.Errorf("want: %s, got: %s", expectedTestCommand, actualTestCommand)
	}
}

func TestGenerateInvalidLanguage(t *testing.T) {
	testConf := &config.Commit0Config{
		Language: "invalidLanguage",
	}
	testTemp := &templator.CITemplator{
		Jenkins:  &template.Template{},
		CircleCI: &template.Template{},
		TravisCI: &template.Template{},
	}
	var wg sync.WaitGroup
	err := ci.Generate(testTemp, testConf, "/dev/null", &wg)
	if err == nil {
		t.Errorf("Error should be thrown with invalid language specified. %s", err.Error())
	}
}

func TestGenerateInvalidCISystem(t *testing.T) {
	testConf := &config.Commit0Config{
		Language: "go",
		CI: config.CI{
			System: "invalidCISystem",
		},
	}
	testTemp := &templator.CITemplator{
		Jenkins:  &template.Template{},
		CircleCI: &template.Template{},
		TravisCI: &template.Template{},
	}
	var wg sync.WaitGroup
	err := ci.Generate(testTemp, testConf, "/dev/null", &wg)
	if err == nil {
		t.Errorf("Error should be thrown with invalid ci system specified. %s", err.Error())
	}
}
