package kubernetes

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup, pathPrefix string) {
	data := templator.GenericTemplateData{*cfg}
	t.Kubernetes.TemplateFiles(data, false, wg, pathPrefix)
}

func Execute(config *config.Commit0Config, pathPrefix string) {
	if config.Infrastructure.AWS.EKS.Deploy {
		log.Println("Planning infrastructure...")
		execute(exec.Command("terraform", "init"), pathPrefix)
		execute(exec.Command("terraform", "plan"), pathPrefix)
	}
}

func execute(cmd *exec.Cmd, pathPrefix string) {
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
