package module_test

import (
	"testing"

	"github.com/commitdev/zero/internal/module"
)

func TestGetSourceDir(t *testing.T) {
	source := "tests/test_data/modules"
	relativeSource := source
	dir := module.GetSourceDir(source)

	t.Log("dir", dir)
	if dir != relativeSource {
		t.Errorf("Error, local sources should not be changed: %s", source)
	}

	source = "github.com/commitdev/my-repo"
	dir = module.GetSourceDir(source)
	if dir == relativeSource {
		t.Errorf("Error, remote sources should be converted to a local dir: %s", source)
	}
}
