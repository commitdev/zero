package generate

import (
	"os"
	"testing"
)

func setupTeardown(t *testing.T) func(t *testing.T) {
	outputPath := "tmp/test"
	os.RemoveAll(outputPath)
	return func(t *testing.T) {
		os.RemoveAll(outputPath)
	}
}

func TestGenerateModules(t *testing.T) {
	teardown := setupTeardown(t)
	defer teardown(t)

	// TODO organize test utils and write assertions
	// generatorConfig := config.LoadGeneratorConfig("../../tests/test_data/configs/zero-basic.yml")

	// GenerateModules(generatorConfig)
}
