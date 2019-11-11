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
	"path/filepath"
	"regexp"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
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

// @TODO : These are specific to a k8s version. If we make the version a config option we will need to change this
var amiLookup = map[string]string{
	"us-east-1":    "ami-0392bafc801b7520f",
	"us-east-2":    "ami-082bb518441d3954c",
	"us-west-2":    "ami-05d586e6f773f6abf",
	"eu-west-1":    "ami-059c6874350e63ca9",
	"eu-central-1": "ami-0e21bc066a9dbabfa",
}

// Generate templates
func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	if cfg.Infrastructure.AWS.EKS.WorkerAMI == "" {
		ami, found := amiLookup[cfg.Infrastructure.AWS.Region]
		if !found {
			log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Unable to look up an AMI for the chosen region")))
		}

		cfg.Infrastructure.AWS.EKS.WorkerAMI = ami
	}
	data := templator.GenericTemplateData{Config: *cfg}
	t.Kubernetes.TemplateFiles(data, false, wg, pathPrefix)
}

// Execute terrafrom init & plan
func Execute(config *config.Commit0Config, pathPrefix string) {
	if config.Infrastructure.AWS.EKS.Deploy {
		log.Println("Preparing aws environment...")

		dir := util.GetCwd()

		if fileExists(fmt.Sprintf("%s/secrets.yaml", dir)) {
			log.Println("secrets.yaml exists ...")
		} else {
			awsSecrets := PromptCredentials()
			writeSecrets(awsSecrets)
		}

		envars := GetAwsEnvars(readSecrets())

		pathPrefix = filepath.Join(pathPrefix, "kubernetes/terraform")

		// @TODO : A check here would be nice to see if this stuff exists first, mostly for testing
		log.Println(aurora.Cyan(emoji.Sprintf(":alarm_clock: Initializing remote backend...")))
		ExecuteCmd(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "bootstrap/remote-state"), envars)
		ExecuteCmd(exec.Command("terraform", "apply", "-auto-approve"), filepath.Join(pathPrefix, "bootstrap/remote-state"), envars)

		log.Println(aurora.Cyan(":alarm_clock: Planning infrastructure..."))
		ExecuteCmd(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "environments/staging"), envars)
		ExecuteCmd(exec.Command("terraform", "plan"), filepath.Join(pathPrefix, "environments/staging"), envars)

		log.Println(aurora.Cyan(":alarm_clock: Applying infrastructure configuration..."))
		ExecuteCmd(exec.Command("terraform", "apply"), filepath.Join(pathPrefix, "environments/staging"), envars)

		log.Println(aurora.Cyan(":alarm_clock: Applying kubernetes configuration..."))
		ExecuteCmd(exec.Command("terraform", "init"), filepath.Join(pathPrefix, "environments/staging/kubernetes"), envars)
		ExecuteCmd(exec.Command("terraform", "plan"), filepath.Join(pathPrefix, "environments/staging/kubernetes"), envars)
	}
}

func ExecuteCmd(cmd *exec.Cmd, pathPrefix string, envars []string) {
	dir := util.GetCwd()

	cmd.Dir = path.Join(dir, pathPrefix)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	var errStdout, errStderr error

	if envars != nil {
		cmd.Env = envars
	}

	err := cmd.Start()
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

func GetAwsEnvars(awsSecrets Secrets) []string {
	env := os.Environ()
	env = append(env, fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", awsSecrets.Aws.AwsAccessKeyID))
	env = append(env, fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", awsSecrets.Aws.AwsSecretAccessKey))
	env = append(env, fmt.Sprintf("AWS_DEFAULT_REGION=%s", awsSecrets.Aws.Region))

	return env
}

func readSecrets() Secrets {

	dir := util.GetCwd()

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

	dir := util.GetCwd()

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

func PromptCredentials() Secrets {

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
