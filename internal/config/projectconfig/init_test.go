package projectconfig_test

import (
	"os"
	"path"
	"testing"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
)

func TestCreateProjectConfigFile(t *testing.T) {
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

	expectedConfig := &projectconfig.ZeroProjectConfig{
		Name:                   projectName,
		ShouldPushRepositories: false,
		Modules:                eksGoReactSampleModules(),
	}
	assert.NoError(t, projectconfig.CreateProjectConfigFile(projectconfig.RootDir, projectName, expectedConfig))

	// make sure the file exists
	if _, err := os.Stat(path.Join(testDirPath, constants.ZeroProjectYml)); err != nil {
		t.Fatal(err)
	}

	t.Run("Should return a valid project config", func(t *testing.T) {
		resultConfig := projectconfig.LoadConfig(path.Join(testDirPath, constants.ZeroProjectYml))

		if !cmp.Equal(expectedConfig, resultConfig, cmpopts.EquateEmpty()) {
			t.Errorf("projectconfig.ZeroProjectConfig.Unmarshal mismatch (-expected +result):\n%s", cmp.Diff(expectedConfig, resultConfig))
		}
	})

	t.Run("Should fail if modules are missing from project config", func(t *testing.T) {
		expectedConfig.Modules = nil
		assert.Error(t, projectconfig.CreateProjectConfigFile(projectconfig.RootDir, projectName, expectedConfig))
	})

}
