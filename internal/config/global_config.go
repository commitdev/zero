package config

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"

	"github.com/commitdev/zero/configs"
	"github.com/commitdev/zero/pkg/util/exit"
	yaml "gopkg.in/yaml.v2"
)

var GetCredentialsPath = getCredentialsPath

type ProjectCredentials map[string]ProjectCredential

type ProjectCredential struct {
	ProjectName            string `yaml:"-"`
	AWSResourceConfig      `yaml:"aws,omitempty"`
	GithubResourceConfig   `yaml:"github,omitempty"`
	CircleCiResourceConfig `yaml:"circleci,omitempty"`
}

type AWSResourceConfig struct {
	AccessKeyId     string `yaml:"accessKeyId,omitempty"`
	SecretAccessKey string `yaml:"secretAccessKey,omitempty"`
}
type GithubResourceConfig struct {
	AccessToken string `yaml:"accessToken,omitempty"`
}
type CircleCiResourceConfig struct {
	ApiKey string `yaml:"apiKey,omitempty"`
}

func (p ProjectCredentials) Unmarshal(data []byte) error {
	if len(data) == 0 {
		return nil
	}
	err := yaml.NewDecoder(bytes.NewReader(data)).Decode(p)
	if err != nil {
		return err
	}
	for k, v := range p {
		v.ProjectName = k
		p[k] = v
	}
	return nil
}

func LoadUserCredentials() ProjectCredentials {
	data := readOrCreateUserCredentialsFile()

	projects := ProjectCredentials{}

	err := projects.Unmarshal(data)

	if err != nil {
		exit.Fatal("Failed to parse configuration: %v", err)
	}
	return projects
}

func getCredentialsPath() string {
	usr, err := user.Current()
	if err != nil {
		exit.Fatal("Failed to get user directory path: %v", err)
	}

	rootDir := path.Join(usr.HomeDir, configs.ZeroHomeDirectory)
	os.MkdirAll(rootDir, os.ModePerm)
	filePath := path.Join(rootDir, configs.UserCredentials)
	return filePath
}

func readOrCreateUserCredentialsFile() []byte {
	credPath := GetCredentialsPath()

	_, fileStateErr := os.Stat(credPath)
	if os.IsNotExist(fileStateErr) {
		var file, fileStateErr = os.Create(credPath)
		if fileStateErr != nil {
			exit.Fatal("Failed to create config file: %v", fileStateErr)
		}
		defer file.Close()
	}
	data, err := ioutil.ReadFile(credPath)
	if err != nil {
		exit.Fatal("Failed to read credentials file: %v", err)
	}
	return data
}

func GetUserCredentials(targetProjectName string) ProjectCredential {
	projects := LoadUserCredentials()

	if val, ok := projects[targetProjectName]; ok {
		return val
	} else {
		p := ProjectCredential{
			ProjectName: targetProjectName,
		}
		projects[targetProjectName] = p
		return p
	}
}

func Save(project ProjectCredential) {
	projects := LoadUserCredentials()
	projects[project.ProjectName] = project
	writeCredentialsFile(projects)
}

func writeCredentialsFile(projects ProjectCredentials) {
	credsPath := GetCredentialsPath()
	content, _ := yaml.Marshal(projects)
	err := ioutil.WriteFile(credsPath, content, 0644)
	if err != nil {
		log.Panicf("failed to write config: %v", err)
	}
}
