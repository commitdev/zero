package credentials

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"regexp"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/commitdev/zero/internal/config/globalconfig"
	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/manifoldco/promptui"
)

// Secrets - AWS prompted credentials
type Secrets struct {
	AWS         AWS
	CircleCIKey string
	GithubToken string
}

type AWS struct {
	AccessKeyID     string
	SecretAccessKey string
}

func MakeAwsEnvars(cfg *projectconfig.ZeroProjectConfig, awsSecrets Secrets) []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsSecrets.AWS.AccessKeyID))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsSecrets.AWS.SecretAccessKey))
	env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", cfg.Infrastructure.AWS.Region))

	return env
}

func AwsCredsPath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return filepath.Join(usr.HomeDir, ".aws/credentials")
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

func ValidateAKID(input string) error {
	// 20 uppercase alphanumeric characters
	var awsAccessKeyIDPat = regexp.MustCompile(`^[A-Z0-9]{20}$`)
	if !awsAccessKeyIDPat.MatchString(input) {
		return errors.New("Invalid aws_access_key_id")
	}
	return nil
}

func ValidateSAK(input string) error {
	// 40 base64 characters
	var awsSecretAccessKeyPat = regexp.MustCompile(`^[A-Za-z0-9/+=]{40}$`)
	if !awsSecretAccessKeyPat.MatchString(input) {
		return errors.New("Invalid aws_secret_access_key")
	}
	return nil
}

func promptAWSCredentials(secrets *Secrets) {
	accessKeyIDPrompt := promptui.Prompt{
		Label:    "Aws Access Key ID ",
		Validate: ValidateAKID,
	}

	accessKeyIDResult, err := accessKeyIDPrompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	secretAccessKeyPrompt := promptui.Prompt{
		Label:    "Aws Secret Access Key ",
		Validate: ValidateSAK,
		Mask:     '*',
	}

	secretAccessKeyResult, err := secretAccessKeyPrompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	secrets.AWS.AccessKeyID = accessKeyIDResult
	secrets.AWS.SecretAccessKey = secretAccessKeyResult
}

func promptGitHubCredentials(secrets *Secrets) {
}

func promptCircleCICredentials(secrets *Secrets) {
	validateKey := func(input string) error {
		// 40 base64 characters
		var awsSecretAccessKeyPat = regexp.MustCompile(`^[A-Za-z0-9]{40}$`)
		if !awsSecretAccessKeyPat.MatchString(input) {
			return errors.New("Invalid CircleCI API Key")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Please enter your CircleCI API key (you can create one at https://circleci.com/account/api) ",
		Validate: validateKey,
	}

	key, err := prompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}
	secrets.CircleCIKey = key
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
