package kubernetes

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
)

func Generate(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	templator.Kubernetes.TemplateFiles(config, false, wg)

	if config.Kubernetes.Deploy {
		_tf_init := tf_init()
		_tf_plan := tf_plan()
		execute(_tf_init)
		execute(_tf_plan)
	}

}

// Terraform init cmd
func tf_init() *exec.Cmd {

	return exec.Command("terraform", "init")
}

// Terraform plan cmd
func tf_plan() *exec.Cmd {

	return exec.Command("terraform", "plan")
}

func execute(cmd *exec.Cmd) {
	dir, err1 := filepath.Abs(filepath.Dir(os.Args[0]))
	if err1 != nil {
		log.Fatal(err1)
	}

	cmd.Dir = dir + "/kubernetes/terraform/environments/staging"

	var stdoutBuf, stderrBuf bytes.Buffer
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()

	var errStdout, errStderr error
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
		wg.Done()
	}()

	_, errStderr = io.Copy(stderr, stderrIn)
	wg.Wait()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := string(stdoutBuf.Bytes()), string(stderrBuf.Bytes())
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}
