package util

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

// Secrets - AWS prompted credentials
type Secrets struct {
	Aws struct {
		AwsAccessKeyID     string
		AwsSecretAccessKey string
		Region             string
	}
}

func MakeAwsEnvars(awsSecrets Secrets) []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsSecrets.Aws.AwsAccessKeyID))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsSecrets.Aws.AwsSecretAccessKey))
	env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", awsSecrets.Aws.Region))

	return env
}

func GetSecrets() Secrets {
	dir := GetCwd()

	if fileExists(fmt.Sprintf("%s/secrets.yaml", dir)) {
		log.Println("secrets.yaml exists ...")
		return readSecrets()
	} else {
		awsSecrets := promptCredentials()
		writeSecrets(awsSecrets)
		return awsSecrets
	}
}

func readSecrets() Secrets {

	dir := GetCwd()

	secretsFile := fmt.Sprintf("%s/secrets.yaml", dir)

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

func writeSecrets(s Secrets) {
	secretsYaml, err := yaml.Marshal(&s)

	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	dir := GetCwd()

	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
		panic(err)
	}

	secretsFile := filepath.Join(dir, "secrets.yaml")
	err = ioutil.WriteFile(secretsFile, []byte(secretsYaml), 0644)

	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}
}

func promptCredentials() Secrets {

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

	regionPrompt := promptui.Select{
		Label: "Select AWS Region ",
		Items: []string{"us-west-1", "us-west-2", "us-east-1", "us-east-2", "ca-central-1",
			"eu-central-1", "eu-west-1", "ap-east-1", "ap-south-1"},
	}

	_, regionResult, err := regionPrompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	awsSecrets := Secrets{}
	awsSecrets.Aws.AwsAccessKeyID = accessKeyIDResult
	awsSecrets.Aws.AwsSecretAccessKey = secretAccessKeyResult
	awsSecrets.Aws.Region = regionResult

	return awsSecrets

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
