package globalconfig_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/commitdev/zero/internal/config/globalconfig"
	"github.com/stretchr/testify/assert"
)

const baseTestFixturesDir = "../../../tests/test_data/configs/"

var testCredentialFile = func() (func() string, func()) {
	tmpConfigPath := getTmpConfig()
	mockFunc := func() string { return tmpConfigPath }
	teardownFunc := func() { os.RemoveAll(tmpConfigPath) }
	return mockFunc, teardownFunc
}

func getTmpConfig() string {
	pathFrom := path.Join(baseTestFixturesDir, fmt.Sprintf("credentials%s.yml", ""))
	pathTo := path.Join(baseTestFixturesDir, fmt.Sprintf("credentials%s.yml", "-tmp"))
	copyFile(pathFrom, pathTo)
	return pathTo
}

func copyFile(from string, to string) {
	bytesRead, err := ioutil.ReadFile(from)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(to, bytesRead, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
func TestReadOrCreateUserCredentialsFile(t *testing.T) {
	globalconfig.GetCredentialsPath = func() string {
		return path.Join(baseTestFixturesDir, "does-not-exist.yml")
	}
	credPath := globalconfig.GetCredentialsPath()

	defer os.RemoveAll(credPath)
	_, fileStateErr := os.Stat(credPath)
	assert.True(t, os.IsNotExist(fileStateErr), "File should not exist")
	// attempting to read the file should create the file
	globalconfig.GetUserCredentials("any-project")

	stats, err := os.Stat(credPath)
	assert.False(t, os.IsNotExist(err), "File should be created")
	assert.Equal(t, "does-not-exist.yml", stats.Name(), "Should create yml automatically")
}

func TestGetUserCredentials(t *testing.T) {
	var teardownFn func()
	globalconfig.GetCredentialsPath, teardownFn = testCredentialFile()
	defer teardownFn()

	t.Run("Fixture file should have existing project with creds", func(t *testing.T) {
		projectName := "my-project"
		project := globalconfig.GetUserCredentials(projectName)

		// Reading from fixtures: tests/test_data/configs/credentials.yml
		assert.Equal(t, "AKIAABCD", project.AWSResourceConfig.AccessKeyId)
		assert.Equal(t, "ZXCV", project.AWSResourceConfig.SecretAccessKey)
		assert.Equal(t, "0987", project.GithubResourceConfig.AccessToken)
		assert.Equal(t, "SOME_API_KEY", project.CircleCiResourceConfig.ApiKey)
	})

	t.Run("Fixture file should support multiple projects", func(t *testing.T) {
		projectName := "another-project"
		project := globalconfig.GetUserCredentials(projectName)
		assert.Equal(t, "654", project.GithubResourceConfig.AccessToken)
	})
}

func TestEditUserCredentials(t *testing.T) {
	var teardownFn func()
	globalconfig.GetCredentialsPath, teardownFn = testCredentialFile()
	defer teardownFn()

	t.Run("Should create new project if not exist", func(t *testing.T) {
		projectName := "test-project3"
		project := globalconfig.GetUserCredentials(projectName)
		project.AWSResourceConfig.AccessKeyId = "TEST_KEY_ID_1"
		globalconfig.Save(project)
		newKeyID := globalconfig.GetUserCredentials(projectName).AWSResourceConfig.AccessKeyId
		assert.Equal(t, "TEST_KEY_ID_1", newKeyID)
	})
	t.Run("Should edit old project if already exist", func(t *testing.T) {
		projectName := "my-project"
		project := globalconfig.GetUserCredentials(projectName)
		project.AWSResourceConfig.AccessKeyId = "EDITED_ACCESS_KEY_ID"
		globalconfig.Save(project)
		newKeyID := globalconfig.GetUserCredentials(projectName).AWSResourceConfig.AccessKeyId
		assert.Equal(t, "EDITED_ACCESS_KEY_ID", newKeyID)
	})
}
