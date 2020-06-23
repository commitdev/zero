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
	dir := "../../tests/test_data/apply/"
	applyConfigPath := constants.ZeroProjectYml
	applyEnvironments := []string{"staging", "production"}

	tmpDir := filepath.Join(os.TempDir(), "apply")

	err := os.RemoveAll(tmpDir)
	assert.NoError(t, err)

	err = shutil.CopyTree(dir, tmpDir, nil)
	assert.NoError(t, err)

	t.Run("Should run apply and execute make on each folder module", func(t *testing.T) {
		apply.Apply(tmpDir, applyConfigPath, applyEnvironments)
		assert.FileExists(t, filepath.Join(tmpDir, "project1/project.out"))
		assert.FileExists(t, filepath.Join(tmpDir, "project2/project.out"))

		content, err := ioutil.ReadFile(filepath.Join(tmpDir, "project1/project.out"))
		assert.NoError(t, err)
		assert.Equal(t, "foo: bar\nrepo: github.com/commitdev/project1\n", string(content))

		content, err = ioutil.ReadFile(filepath.Join(tmpDir, "project2/project.out"))
		assert.NoError(t, err)
		assert.Equal(t, "baz: qux\n", string(content))
	})
}
