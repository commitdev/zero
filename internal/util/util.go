package util

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"text/template"

	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

func CreateDirIfDoesNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

func CleanGoIdentifier(identifier string) string {
	return strings.ReplaceAll(identifier, "-", "")
}

var FuncMap = template.FuncMap{
	"Title":             strings.Title,
	"ToLower":           strings.ToLower,
	"CleanGoIdentifier": CleanGoIdentifier,
}

func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
		panic(err)
	}

	return dir
}

func createTemplatedFile(fullFilePath string, template *template.Template, wg *sync.WaitGroup, data interface{}) {
	f, err := os.Create(fullFilePath)
	if err != nil {
		log.Println(aurora.Red(emoji.Sprintf(":exclamation: Error creating file '%s' : %v", fullFilePath, err)))
	}
	wg.Add(1)
	go func() {
		err = template.Execute(f, data)
		if err != nil {
			log.Println(aurora.Red(emoji.Sprintf(":exclamation: Error templating '%s': %v", fullFilePath, err)))
		}
		log.Println(aurora.Green(emoji.Sprintf(":white_check_mark: Finished templating : %v", fullFilePath)))
		wg.Done()
	}()
}

func TemplateFileAndOverwrite(fileDir string, fileName string, template *template.Template, wg *sync.WaitGroup, data interface{}) {
	fullFilePath := fmt.Sprintf("%v/%v", fileDir, fileName)
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		log.Println(aurora.Red(emoji.Sprintf(":exclamation: Error creating directory %v: %v", fullFilePath, err)))
	}
	createTemplatedFile(fullFilePath, template, wg, data)

}

func TemplateFileIfDoesNotExist(fileDir string, fileName string, template *template.Template, wg *sync.WaitGroup, data interface{}) {
	fullFilePath := path.Join(fileDir, fileName)

	if _, err := os.Stat(fullFilePath); os.IsNotExist(err) {
		if fileDir != "" {
			err := CreateDirIfDoesNotExist(fileDir)
			if err != nil {
				log.Println(aurora.Red(emoji.Sprintf(":exclamation: Error creating directory %v: %v", fullFilePath, err)))
			}
		}
		createTemplatedFile(fullFilePath, template, wg, data)
	} else {
		log.Println(aurora.Yellow(emoji.Sprintf("%v already exists. skipping.", fullFilePath)))
	}
}

func ExecuteCommand(cmd *exec.Cmd, pathPrefix string, envars []string) {
	dir := GetCwd()

	cmd.Dir = path.Join(dir, pathPrefix)

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	var errStdout, errStderr error

	if envars != nil {
		cmd.Env = envars
	}

	err := cmd.Start()
	if err != nil {
		log.Fatalf("Starting command failed: %v\n", err)
	}

	go func() {
		_, errStdout = io.Copy(os.Stdout, stdoutPipe)
	}()
	go func() {
		_, errStderr = io.Copy(os.Stderr, stderrPipe)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("Executing command failed: %v\n", err)
	}

	if errStdout != nil {
		log.Printf("Failed to capture stdout: %v\n", errStdout)
	}

	if errStderr != nil {
		log.Printf("Failed to capture stderr: %v\n", errStderr)
	}
}

func ExecuteCommandOutput(cmd *exec.Cmd, pathPrefix string, envars []string) string {
	dir := GetCwd()

	cmd.Dir = path.Join(dir, pathPrefix)

	if envars != nil {
		cmd.Env = envars
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Executing terraform output failed: %v\n", err)
	}
	return string(out)
}
