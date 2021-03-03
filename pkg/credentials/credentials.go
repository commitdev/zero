package credentials

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/commitdev/zero/internal/config/globalconfig"
)

type AWSResourceConfig struct {
	AccessKeyID     string `yaml:"accessKeyId,omitempty" env:"AWS_ACCESS_KEY_ID,omitempty"`
	SecretAccessKey string `yaml:"secretAccessKey,omitempty" env:"AWS_SECRET_ACCESS_KEY,omitempty"`
}

func AwsCredsPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".aws/credentials")
}

func FillAWSProfile(profileName string, paramsToFill map[string]string) error {
	awsPath := AwsCredsPath()
	awsCreds, err := credentials.NewSharedCredentials(awsPath, profileName).Get()
	if err != nil {
		return err
	}
	paramsToFill["accessKeyId"] = awsCreds.AccessKeyID
	paramsToFill["secretAccessKey"] = awsCreds.SecretAccessKey
	return nil
}

func GetAWSProfileProjectCredentials(profileName string, creds globalconfig.ProjectCredential) globalconfig.ProjectCredential {
	awsPath := AwsCredsPath()
	return GetAWSProfileCredentials(awsPath, profileName, creds)
}

func GetAWSProfileCredentials(credsPath string, profileName string, creds globalconfig.ProjectCredential) globalconfig.ProjectCredential {
	awsCreds, err := credentials.NewSharedCredentials(credsPath, profileName).Get()
	if err != nil {
		log.Fatal(err)
	}
	creds.AWSResourceConfig.AccessKeyID = awsCreds.AccessKeyID
	creds.AWSResourceConfig.SecretAccessKey = awsCreds.SecretAccessKey
	return creds
}

// GetAWSProfiles returns a list of AWS forprofiles set up on the user's sytem
func GetAWSProfiles() ([]string, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	// Load the credentials file to look for profiles
	credsFile := filepath.Join(usr.HomeDir, ".aws/credentials")
	creds, err := ioutil.ReadFile(credsFile)
	if err != nil {
		return nil, err
	}
	// Get all profiles
	re := regexp.MustCompile(`\[(.*)\]`)
	profileMatches := re.FindAllStringSubmatch(string(creds), -1)
	profiles := make([]string, len(profileMatches))
	for i, p := range profileMatches {
		profiles[i] = p[1]
	}
	return profiles, nil
}
