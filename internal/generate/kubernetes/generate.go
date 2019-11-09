package kubernetes

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
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

// Generate templates
func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	data := templator.GenericTemplateData{*cfg}
	t.Kubernetes.TemplateFiles(data, false, wg, pathPrefix)
}

func getCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
		panic(err)
	}

	return dir
}

// Execute terrafrom init & plan
func Execute(config *config.Commit0Config, pathPrefix string) {
	if config.Infrastructure.AWS.EKS.Deploy {
		log.Println("Preparing aws environment...")

		dir := getCwd()

		if fileExists(fmt.Sprintf("%s/secrets.yaml", dir)) {
			log.Println("secrets.yaml exists ...")
		} else {
			awsSecrets := promptCredentials()
			writeSecrets(awsSecrets)
		}

		envars := getAwsEnvars(readSecrets())
		log.Println("Planning infrastructure...")
		execute(exec.Command("terraform", "init"), pathPrefix, envars)
		execute(exec.Command("terraform", "plan"), pathPrefix, envars)
	}
}

func execute(cmd *exec.Cmd, pathPrefix string, envars []string) {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
	}
	kubDir := path.Join(pathPrefix, "kubernetes/terraform/environments/staging")
	cmd.Dir = path.Join(dir, kubDir)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	var errStdout, errStderr error

	if envars != nil {
		log.Println("Setting envars to cmd ...")
		cmd.Env = envars
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Starting terraform command failed: %v\n", err)
	}

	go func() {
		_, errStdout = io.Copy(os.Stdout, stdoutPipe)
	}()
	go func() {
		_, errStderr = io.Copy(os.Stderr, stderrPipe)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Executing terraform command failed: %v\n", err)
	}

	if errStdout != nil {
		log.Printf("Failed to capture stdout: %v\n", errStdout)
	}

	if errStderr != nil {
		log.Printf("Failed to capture stderr: %v\n", errStderr)
	}
}

func getAwsEnvars(awsSecrets Secrets) []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsSecrets.Aws.AwsAccessKeyID))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsSecrets.Aws.AwsSecretAccessKey))
	env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", awsSecrets.Aws.Region))

	return env
}

func readSecrets() Secrets {

	dir := getCwd()

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

	fmt.Printf("SECRETS: %v", string(secretsYaml))

	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
		panic(err)
	}

	secretsFile := fmt.Sprintf("%s/secrets.yaml", dir)
	// err = ioutil.WriteFile(secretsFile, []byte(secretsYaml), 0644)
	f, err := os.Create(secretsFile)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	defer f.Close()

	n3, err := f.WriteString(string(secretsYaml))
	f.Sync()
	log.Printf("Wrote %d bytes to %v", n3, secretsFile)

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
		var awsAccessKeyIDPat = regexp.MustCompile(`^[A-Za-z0-9/+=]{40}$`)
		if !awsAccessKeyIDPat.MatchString(input) {
			return errors.New("Invalid aws_secret_access_key")
		}
		return nil
	}

	accessKeyIDPrompt := promptui.Prompt{
		Label:    "Aws Access Key ID: ",
		Validate: validateAKID,
	}

	accessKeyIDResult, err := accessKeyIDPrompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	secretAccessKeyPrompt := promptui.Prompt{
		Label:    "Aws Secret Access Key: ",
		Validate: validateSAK,
		Mask:     '*',
	}

	secretAccessKeyResult, err := secretAccessKeyPrompt.Run()

	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	regionPrompt := promptui.Select{
		Label: "Select AWS Region: ",
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
