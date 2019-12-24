package templator

import (
	"testing"
)

func TestDirectoryTemplates(t *testing.T) {
	templator := NewDirTemplator("../../tests/test_data/modules/ci", []string{"{{", "}}"})
	var data = map[string]string{
		"ci": "github",
	}

	templator.ExecuteTemplates(data, false, "tmp")
}

func TestExecuteTemplate(t *testing.T) {
	filePath := "test_file.yml"
	outputPath := "test_file_output.yml"
	data := map[string]string{
		"foobar": "hello",
	}
	ExecuteTemplate(filePath, outputPath, data)
}
