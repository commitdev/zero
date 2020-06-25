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
	globalconfig.GetProjectCredentials("any-project")

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
		project := globalconfig.GetProjectCredentials(projectName)

		// Reading from fixtures: tests/test_data/configs/credentials.yml
		assert.Equal(t, "AKIAABCD", project.AWSResourceConfig.AccessKeyID)
		assert.Equal(t, "ZXCV", project.AWSResourceConfig.SecretAccessKey)
		assert.Equal(t, "0987", project.GithubResourceConfig.AccessToken)
		assert.Equal(t, "SOME_API_KEY", project.CircleCiResourceConfig.ApiKey)
	})

	t.Run("Fixture file should support multiple projects", func(t *testing.T) {
		projectName := "another-project"
		project := globalconfig.GetProjectCredentials(projectName)
		assert.Equal(t, "654", project.GithubResourceConfig.AccessToken)
	})
}

func TestEditUserCredentials(t *testing.T) {
	var teardownFn func()
	globalconfig.GetCredentialsPath, teardownFn = testCredentialFile()
	defer teardownFn()

	t.Run("Should create new project if not exist", func(t *testing.T) {
		projectName := "test-project3"
		project := globalconfig.GetProjectCredentials(projectName)
		project.AWSResourceConfig.AccessKeyID = "TEST_KEY_ID_1"
		globalconfig.Save(project)
		newKeyID := globalconfig.GetProjectCredentials(projectName).AWSResourceConfig.AccessKeyID
		assert.Equal(t, "TEST_KEY_ID_1", newKeyID)
	})
	t.Run("Should edit old project if already exist", func(t *testing.T) {
		projectName := "my-project"
		project := globalconfig.GetProjectCredentials(projectName)
		project.AWSResourceConfig.AccessKeyID = "EDITED_ACCESS_KEY_ID"
		globalconfig.Save(project)
		newKeyID := globalconfig.GetProjectCredentials(projectName).AWSResourceConfig.AccessKeyID
		assert.Equal(t, "EDITED_ACCESS_KEY_ID", newKeyID)
	})
}

func TestMarshalProjectCredentialAsEnvVars(t *testing.T) {
	t.Run("Should be able to marshal a ProjectCredential into env vars", func(t *testing.T) {
		pc := globalconfig.ProjectCredential{
			AWSResourceConfig: globalconfig.AWSResourceConfig{
				AccessKeyID:     "AKID",
				SecretAccessKey: "SAK",
			},
			CircleCiResourceConfig: globalconfig.CircleCiResourceConfig{
				ApiKey: "APIKEY",
			},
		}

		envVars := pc.AsEnvVars()
		assert.Equal(t, "AKID", envVars["AWS_ACCESS_KEY_ID"])
		assert.Equal(t, "SAK", envVars["AWS_SECRET_ACCESS_KEY"])
		assert.Equal(t, "APIKEY", envVars["CIRCLECI_API_KEY"])
	})

	t.Run("should honor omitempty and left out empty values", func(t *testing.T) {
		pc := globalconfig.ProjectCredential{}

		envVars := pc.AsEnvVars()
		assert.Equal(t, 0, len(envVars))
	})
}

func TestMarshalSelectedVendorsCredentialsAsEnv(t *testing.T) {
	pc := globalconfig.ProjectCredential{
		AWSResourceConfig: globalconfig.AWSResourceConfig{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SAK",
		},
		GithubResourceConfig: globalconfig.GithubResourceConfig{
			AccessToken: "FOOBAR",
		},
		CircleCiResourceConfig: globalconfig.CircleCiResourceConfig{
			ApiKey: "APIKEY",
		},
	}

	t.Run("cherry pick credentials by vendor", func(t *testing.T) {
		envs := pc.SelectedVendorsCredentialsAsEnv([]string{"aws", "github"})
		assert.Equal(t, "AKID", envs["AWS_ACCESS_KEY_ID"])
		assert.Equal(t, "SAK", envs["AWS_SECRET_ACCESS_KEY"])
		assert.Equal(t, "FOOBAR", envs["GITHUB_ACCESS_TOKEN"])
	})

	t.Run("omits vendors not selected", func(t *testing.T) {
		envs := pc.SelectedVendorsCredentialsAsEnv([]string{"github"})
		assert.Equal(t, "FOOBAR", envs["GITHUB_ACCESS_TOKEN"])

		_, hasAWSKeyID := envs["AWS_ACCESS_KEY_ID"]
		assert.Equal(t, false, hasAWSKeyID)
		_, hasAWSSecretAccessKey := envs["AWS_SECRET_ACCESS_KEY"]
		assert.Equal(t, false, hasAWSSecretAccessKey)
		_, hasCircleCIKey := envs["CIRCLECI_API_KEY"]
		assert.Equal(t, false, hasCircleCIKey)
	})

	t.Run("omits vendors not selected", func(t *testing.T) {
		envs := pc.SelectedVendorsCredentialsAsEnv([]string{})
		assert.Equal(t, 0, len(envs))
	})

}
