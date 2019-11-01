package kubernetes

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	templator.Kubernetes.TemplateFiles(config, false, wg)

}

func Execute(config *config.Commit0Config) {
	if config.Infrastructure.AWS.EKS.Deploy {
		log.Println("Planning infrastructure...")
		execute(exec.Command("terraform", "init"))
		execute(exec.Command("terraform", "plan"))
	}
}

func execute(cmd *exec.Cmd) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
	}

	cmd.Dir = fmt.Sprintf("%s/kubernetes/terraform/environments/staging", dir)

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
