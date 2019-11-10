package cmd_test

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/commitdev/commit0/cmd"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
)

func TestCreateWorks(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "commit0-")
	if err != nil {
		t.Fatal(err)
	}

	projectName := "test-project"

	templates := packr.New("templates", "../templates")
	templator := templator.NewTemplator(templates)

	root := cmd.Create(projectName, tmpdir, templator)
	defer os.RemoveAll(tmpdir)

	st, err := os.Stat(path.Join(root, util.CommitYml))
	if err != nil {
		t.Fatal(err)
	}

	if st.Size() == 0 {
		t.Fatalf("commit0.yml is empty")
	}
}
