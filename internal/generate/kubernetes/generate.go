package kubernetes

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

type Secrets struct {
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	data := templator.GenericTemplateData{*cfg}
	t.Kubernetes.TemplateFiles(data, false, wg, pathPrefix)
}

func Execute(config *config.Commit0Config, pathPrefix string) {
	if config.Infrastructure.AWS.EKS.Deploy {
		log.Println("Preparing aws environment...")

		dir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Getting working directory failed: %v\n", err)
		}

		if fileExists(fmt.Sprintf("%s/secrets.yaml", dir)) {
			log.Println("secrets.yaml exists ...")
		} else {
			awsSecrets := promptCredentials()
			writeSecrets(awsSecrets)
		}

		log.Println("Planning infrastructure...")
		execute(exec.Command("terraform", "init"), pathPrefix)
		execute(exec.Command("terraform", "plan"), pathPrefix)
	}
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
	log.Printf("wrote %d bytes to %v", n3, secretsFile)

}

func promptCredentials() Secrets {

	// validate := func(input string) error {
	// 	var awsAccessKeyIDPat = regexp.MustCompile(`[A-Z0-9]{20}(?![A-Z0-9])`)
	// 	if !awsAccessKeyIDPat.MatchString(input) {
	// 		return errors.New("Invalid number")
	// 	}
	// 	return nil
	// }

	accessKeyIDPrompt := promptui.Prompt{
		Label: "Aws Access Key ID",
	}

	accessKeyIDResult, err1 := accessKeyIDPrompt.Run()

	if err1 != nil {
		log.Fatalf("Prompt failed %v\n", err1)
		panic(err1)
	}

	// awsSecrets.aws.awsAccessKeyID = accessKeyIDResult

	secretAccessKeyPrompt := promptui.Prompt{
		Label: "Aws Access Key ID",
	}

	secretAccessKeyResult, err2 := secretAccessKeyPrompt.Run()

	if err2 != nil {
		log.Fatalf("Prompt failed %v\n", err2)
		panic(err2)
	}

	// awsSecrets.aws.awsSecretAccessKey = secretAccessKeyResult
	awsSecrets := Secrets{
		accessKeyIDResult,
		secretAccessKeyResult,
	}

	return awsSecrets

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func execute(cmd *exec.Cmd) {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
	}
	kubDir := path.Join(pathPrefix, "kubernetes/terraform/environments/staging")
	cmd.Dir = path.Join(dir, kubDir)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	var errStdout, errStderr error
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
