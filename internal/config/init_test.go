package config_test

import (
	"os"
	"path"
	"testing"

	"github.com/commitdev/zero/configs"
	"github.com/commitdev/zero/internal/config"
)

func TestInit(t *testing.T) {
	const testDir = "../../test-sandbox"
	projectName := "test-project"

	config.SetRootDir(testDir)
	defer os.RemoveAll(testDir)

	testDirPath := path.Join(config.RootDir, projectName)
	// create sandbox dir
	err := os.MkdirAll(testDirPath, os.ModePerm)
	if err != nil {
		t.Fatal(err)
	}

	config.Init(config.RootDir, projectName, nil)

	if _, err := os.Stat(path.Join(testDirPath, configs.ZeroProjectYml)); err != nil {
		t.Fatal(err)
	}
}
