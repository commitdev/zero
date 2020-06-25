package util

// @TODO split up and move into /pkg directory

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/google/uuid"
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

// @TODO how can we make these type of helpers extensible?
var FuncMap = template.FuncMap{
	"Title":             strings.Title,
	"ToLower":           strings.ToLower,
	"CleanGoIdentifier": CleanGoIdentifier,
	"GenerateUUID":      uuid.New,
}

func GetCwd() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Getting working directory failed: %v\n", err)
		panic(err)
	}

	return dir
}

func ExecuteCommand(cmd *exec.Cmd, pathPrefix string, envars []string) {

	cmd.Dir = pathPrefix
	if !filepath.IsAbs(pathPrefix) {
		dir := GetCwd()
		cmd.Dir = path.Join(dir, pathPrefix)
	}

	stdoutPipe, _ := cmd.StdoutPipe()
	stderrPipe, _ := cmd.StderrPipe()

	var errStdout, errStderr error

	cmd.Env = os.Environ()
	if envars != nil {
		cmd.Env = append(os.Environ(), envars...)
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

// ExecuteCommandOutput runs the command and returns its
// combined standard output and standard error.
func ExecuteCommandOutput(cmd *exec.Cmd, pathPrefix string, envars []string) string {

	cmd.Dir = pathPrefix
	if !filepath.IsAbs(pathPrefix) {
		dir := GetCwd()
		cmd.Dir = path.Join(dir, pathPrefix)
	}

	cmd.Env = os.Environ()
	if envars != nil {
		cmd.Env = append(os.Environ(), envars...)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Executing command with output failed: (%v) %s\n", err, out)
	}
	return string(out)
}

// AppendProjectEnvToCmdEnv will add all the keys and values from envMap
// into envList as key-value pair strings (e.g.: "key=value")
func AppendProjectEnvToCmdEnv(envMap map[string]string, envList []string) []string {
	for key, val := range envMap {
		if val != "" {
			envList = append(envList, fmt.Sprintf("%s=%s", key, val))
		}
	}
	return envList
}

// IndentString will Add x space char padding at the beginging of each line.
func IndentString(content string, spaces int) string {
	var result string
	subStr := strings.Split(content, "\n")
	for _, s := range subStr {
		result += fmt.Sprintf("%"+strconv.Itoa(spaces)+"s%s\n", "", s)
	}
	return result
}

func ItemInSlice(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
