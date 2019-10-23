package cmd_test

import (
	"github.com/commitdev/commit0/cmd"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestInitWorks(t *testing.T) {
	cmd.Init()
}

func TestCreateWorks(t *testing.T) {
	tmpdir, err := ioutil.TempDir("", "commit0-")
	if err != nil {
		t.Fatal(err)
	}

	projectName := "test-project"

	root := cmd.Create(projectName, tmpdir)
	defer os.RemoveAll(tmpdir)

	st, err := os.Stat(path.Join(root, "commit0.yml"))
	if err != nil {
		t.Fatal(err)
	}

	if st.Size() == 0 {
		t.Fatalf("commit0.yml is empty")
	}
}
