package templator

import (
	"os"
	"testing"
)

var testData = "../../tests/test_data/ci/"

func setupTeardown(t *testing.T) func(t *testing.T) {
	outputPath := "../../tmp/tests"
	os.RemoveAll(outputPath)
	return func(t *testing.T) {
		os.RemoveAll(outputPath)
	}
}

func TestDirectoryTemplates(t *testing.T) {
	teardown := setupTeardown(t)
	defer teardown(t)

	// TODO organize test utils and write assertions
	templator := NewDirTemplator("../../tests/test_data/modules/ci", []string{"{{", "}}"})
	var data = map[string]string{
		"ci": "github",
	}

	templator.ExecuteTemplates(data, false, "tmp", "")
}

// func TestExecuteTemplate(t *testing.T) {
// 	filePath := "test_file.yml"
// 	outputPath := "test_file_output.yml"
// 	data := map[string]string{
// 		"foobar": "hello",
// 	}
// 	ExecuteTemplate(filePath, outputPath, data)
// }
