package module

import (
	"testing"
)

func TestIsLocal(t *testing.T) {
	source := "./tests/test_data/modules/ci"
	res := IsLocal(source)
	if !res {
		t.Errorf("Error, source %s SHOULD BE determined as local", source)
	}

	source = "https://github.com/zthomas/react-mui-kit"
	res = IsLocal(source)
	if res {
		t.Errorf("Error, source %s SHOULD NOT BE determined as local", source)
	}
}

func TestGetSourceDir(t *testing.T) {
	source := "tests/test_data/modules/ci"
	relativeSource := source
	dir := GetSourceDir(source)

	t.Log("dir", dir)
	if dir != relativeSource {
		t.Errorf("Error, local sources should not be changed: %s", source)
	}

	source = "github.com/zthomas/react-mui-kit"
	dir = GetSourceDir(source)
	if dir == relativeSource {
		t.Errorf("Error, remote sources should be converted to a local dir: %s", source)
	}
}

// var testData = "../../tests/test_data/ci/"

// func setupTeardown(t *testing.T) func(t *testing.T) {
// 	outputPath := "../../tmp/tests"
// 	os.RemoveAll(outputPath)
// 	return func(t *testing.T) {
// 		os.RemoveAll(outputPath)
// 	}
// }

// func TestDirectoryTemplates(t *testing.T) {
// 	teardown := setupTeardown(t)
// 	defer teardown(t)

// 	// TODO organize test utils and write assertions
// 	templator := NewDirTemplator("../../tests/test_data/modules/ci", []string{"{{", "}}"})
// 	var data = map[string]string{
// 		"ci": "github",
// 	}

// 	templator.ExecuteTemplates(data, false, "tmp", "")
// }
