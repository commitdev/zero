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
	"github.com/commitdev/commit0/internal/config"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
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

func MakeAwsEnvars(cfg *config.Commit0Config, awsSecrets Secrets) []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsSecrets.AWS.AccessKeyID))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsSecrets.AWS.SecretAccessKey))
	env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", cfg.Infrastructure.AWS.Region))

	return env
}

func GetSecrets(baseDir string) Secrets {

	secretsFile := filepath.Join(baseDir, "secrets.yaml")
	if fileExists(secretsFile) {
		log.Println("secrets.yaml exists ...")
		return readSecrets(secretsFile)
	} else {
		// Get the user's home dir
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}
		credsFile := filepath.Join(usr.HomeDir, ".aws/credentials")

		var secrets Secrets

		// Load the credentials file to look for profiles
		profiles, err := GetAWSProfiles()
		if err == nil {
			profilePrompt := promptui.Select{
				Label: "Select AWS Profile",
				Items: profiles,
			}

			_, profileResult, _ := profilePrompt.Run()

			creds, err := credentials.NewSharedCredentials(credsFile, profileResult).Get()
			if err == nil {
				secrets = Secrets{
					AWS: AWS{
						AccessKeyID:     creds.AccessKeyID,
						SecretAccessKey: creds.SecretAccessKey,
					},
				}
			}
		}

		// We couldn't load the credentials file, get the user to just paste them
		if secrets.AWS == (AWS{}) {
			promptAWSCredentials(&secrets)
		}

		if secrets.CircleCIKey == "" || secrets.GithubToken == "" {
			ciPrompt := promptui.Select{
				Label: "Which Continuous integration provider do you want to use?",
				Items: []string{"CircleCI", "GitHub Actions"},
			}

			_, ciResult, _ := ciPrompt.Run()

			if ciResult == "CircleCI" {
				promptCircleCICredentials(&secrets)
			} else if ciResult == "GitHub Actions" {
				promptGitHubCredentials(&secrets)
			}
		}

		return secrets
	}
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

func readSecrets(secretsFile string) Secrets {
	data, err := ioutil.ReadFile(secretsFile)
	if err != nil {
		log.Fatalln(err)
	}

	awsSecrets := Secrets{}

	err = yaml.Unmarshal(data, &awsSecrets)
	if err != nil {
		log.Fatalln(err)
	}

	return awsSecrets
}

func writeSecrets(secretsFile string, s Secrets) {
	secretsYaml, err := yaml.Marshal(&s)

	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	err = ioutil.WriteFile(secretsFile, []byte(secretsYaml), 0644)

	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}
}

func promptAWSCredentials(secrets *Secrets) {

	validateAKID := func(input string) error {
		// 20 uppercase alphanumeric characters
		var awsAccessKeyIDPat = regexp.MustCompile(`^[A-Z0-9]{20}$`)
		if !awsAccessKeyIDPat.MatchString(input) {
			return errors.New("Invalid aws_access_key_id")
		}
		return nil
	}

	validateSAK := func(input string) error {
		// 40 base64 characters
		var awsSecretAccessKeyPat = regexp.MustCompile(`^[A-Za-z0-9/+=]{40}$`)
		if !awsSecretAccessKeyPat.MatchString(input) {
			return errors.New("Invalid aws_secret_access_key")
		}
		return nil
	}

	accessKeyIDPrompt := promptui.Prompt{
		Label:    "Aws Access Key ID ",
		Validate: validateAKID,
	}

	accessKeyIDResult, err := accessKeyIDPrompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	secretAccessKeyPrompt := promptui.Prompt{
		Label:    "Aws Secret Access Key ",
		Validate: validateSAK,
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
