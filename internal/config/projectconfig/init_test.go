package projectconfig_test

import (
	"os"
	"path"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
)

func TestInit(t *testing.T) {
	const testDir = "../../test-sandbox"
	projectName := "test-project"

	projectconfig.SetRootDir(testDir)
	defer os.RemoveAll(testDir)

	testDirPath := path.Join(projectconfig.RootDir, projectName)
	// create sandbox dir
	err := os.MkdirAll(testDirPath, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	config := projectconfig.ZeroProjectConfig{}
	projectconfig.Init(projectconfig.RootDir, projectName, &config)

	if _, err := os.Stat(path.Join(testDirPath, constants.ZeroProjectYml)); err != nil {
		t.Fatal(err)
	}
}
