package apply_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/commitdev/zero/internal/apply"
	"github.com/commitdev/zero/internal/constants"
	"github.com/stretchr/testify/assert"
	"github.com/termie/go-shutil"
)

func TestApply(t *testing.T) {
	applyConfigPath := constants.ZeroProjectYml
	applyEnvironments := []string{"staging", "production"}
	var tmpDir string

	t.Run("Should run apply and execute make on each folder module", func(t *testing.T) {
		tmpDir = setupTmpDir(t, "../../tests/test_data/apply/")
		err := apply.Apply(tmpDir, applyConfigPath, applyEnvironments)
		assert.FileExists(t, filepath.Join(tmpDir, "project1/project.out"))
		assert.FileExists(t, filepath.Join(tmpDir, "project2/project.out"))

		content, err := ioutil.ReadFile(filepath.Join(tmpDir, "project1/project.out"))
		assert.NoError(t, err)
		assert.Equal(t, "foo: bar\nrepo: github.com/commitdev/project1\n", string(content))

		content, err = ioutil.ReadFile(filepath.Join(tmpDir, "project2/project.out"))
		assert.NoError(t, err)
		assert.Equal(t, "baz: qux\n", string(content))

	})

	t.Run("Moudles runs command overides", func(t *testing.T) {
		content, err := ioutil.ReadFile(filepath.Join(tmpDir, "project2/check.out"))
		assert.NoError(t, err)
		assert.Equal(t, "custom check\n", string(content))
	})

	t.Run("Zero apply honors the envVarName overwrite from module definition", func(t *testing.T) {
		content, err := ioutil.ReadFile(filepath.Join(tmpDir, "project1/feature.out"))
		assert.NoError(t, err)
		assert.Equal(t, "envVarName of viaEnvVarName: baz\n", string(content))
	})

	t.Run("Moudles with failing checks should return error", func(t *testing.T) {
		tmpDir = setupTmpDir(t, "../../tests/test_data/apply-failing/")

		err := apply.Apply(tmpDir, applyConfigPath, applyEnvironments)
		assert.Regexp(t, "^Module checks failed:", err.Error())
		assert.Regexp(t, "Module \\(project1\\)", err.Error())
		assert.Regexp(t, "Module \\(project2\\)", err.Error())
		assert.Regexp(t, "Module \\(project3\\)", err.Error())
	})

}

func setupTmpDir(t *testing.T, exampleDirPath string) string {
	var err error
	tmpDir := filepath.Join(os.TempDir(), "apply")

	err = os.RemoveAll(tmpDir)
	assert.NoError(t, err)

	err = shutil.CopyTree(exampleDirPath, tmpDir, nil)
	assert.NoError(t, err)
	return tmpDir
}
