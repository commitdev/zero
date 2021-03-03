package credentials

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/commitdev/zero/internal/util"
)

type AWSResourceConfig struct {
	AccessKeyID     string `yaml:"accessKeyId,omitempty" env:"AWS_ACCESS_KEY_ID,omitempty"`
	SecretAccessKey string `yaml:"secretAccessKey,omitempty" env:"AWS_SECRET_ACCESS_KEY,omitempty"`
}

var GetAWSCredsPath = awsCredsPath

func awsCredsPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".aws/credentials")
}

func fetchAWSConfig(awsPath string, profileName string) (error, AWSResourceConfig) {

	awsCreds, err := credentials.NewSharedCredentials(awsPath, profileName).Get()
	if err != nil {
		return err, AWSResourceConfig{}
	}
	return nil, AWSResourceConfig{
		AccessKeyID:     awsCreds.AccessKeyID,
		SecretAccessKey: awsCreds.SecretAccessKey,
	}
}

func FillAWSProfile(profileName string, paramsToFill map[string]string) error {
	awsPath := GetAWSCredsPath()
	err, awsCreds := fetchAWSConfig(awsPath, profileName)
	if err != nil {
		return err
	}
	util.ReflectStructValueIntoMap(awsCreds, "yaml", paramsToFill)
	return nil
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
