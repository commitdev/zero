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
	AccessKeyID     string `key:"accessKeyId"`
	SecretAccessKey string `key:"secretAccessKey"`
}

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

// FillAWSProfile receives the AWS profile name, then parses
// the accessKeyId / secretAccessKey values into a map
func FillAWSProfile(pathToCredentialsFile string, profileName string, paramsToFill map[string]string) error {
	if pathToCredentialsFile == "" {
		pathToCredentialsFile = awsCredsPath()
	}

	err, awsCreds := fetchAWSConfig(pathToCredentialsFile, profileName)
	if err != nil {
		return err
	}
	util.ReflectStructValueIntoMap(awsCreds, "key", paramsToFill)
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
