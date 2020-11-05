package generate_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/generate"
	"github.com/stretchr/testify/assert"
)

const baseTestFixturesDir = "../../tests/test_data/generate/"

func setupTeardown(t *testing.T) (func(t *testing.T), string) {
	tmpDir := filepath.Join(os.TempDir(), "generate")
	os.MkdirAll(tmpDir, 0755)
	os.RemoveAll(tmpDir)
	return func(t *testing.T) {
		os.RemoveAll(tmpDir)
	}, tmpDir
}

func TestGenerateModules(t *testing.T) {
	teardown, tmpDir := setupTeardown(t)
	defer teardown(t)

	projectConfig := projectconfig.ZeroProjectConfig{
		Name: "foo",
		Modules: projectconfig.Modules{
			"mod1": projectconfig.NewModule(map[string]string{"test": "bar"}, tmpDir, "github.com/fake-org/repo-foo", baseTestFixturesDir, []string{}, []projectconfig.Condition{}),
		},
	}
	generate.Generate(projectConfig, true)

	content, err := ioutil.ReadFile(filepath.Join(tmpDir, "file_to_template.txt"))
	assert.NoError(t, err)

	expectedContent := `Name is foo
Params.test is bar
Files.Repository is github.com/fake-org/repo-foo
`
	assert.Equal(t, string(content), expectedContent)
}
