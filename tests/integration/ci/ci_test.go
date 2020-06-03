package ci_test

// @TODO refactor into new set of integration tests
// import (
// 	"bytes"
// 	"io/ioutil"
// 	"os"
// 	"sync"
// 	"testing"

// 	"github.com/commitdev/zero/internal/config"
// 	"github.com/commitdev/zero/internal/generate/ci"
// 	"github.com/commitdev/zero/internal/templator"
// 	"github.com/gobuffalo/packr/v2"
// )

// var testData = "../../test_data/ci/"

// // setupTeardown removes all the generated test files before and after
// // the test runs to ensure clean data.
// func setupTeardown(t *testing.T) func(t *testing.T) {
// 	os.RemoveAll("../../test_data/ci/actual")
// 	return func(t *testing.T) {
// 		os.RemoveAll("../../test_data/ci/actual")
// 	}
// }

// func TestGenerateJenkins(t *testing.T) {
// 	teardown := setupTeardown(t)
// 	defer teardown(t)

// 	templates := packr.New("templates", "../../../templates")
// 	testTemplator := templator.NewTemplator(templates)

// 	var waitgroup sync.WaitGroup

// 	testConf := &projectconfig.ZeroProjectConfig{}
// 	testCI := config.CI{
// 		System:       "jenkins",
// 		BuildImage:   "golang/golang",
// 		BuildTag:     "1.12",
// 		BuildCommand: "make build",
// 		TestCommand:  "make test",
// 	}

// 	err := ci.Generate(testTemplator.CI, testConf, testCI, testData+"/actual", &waitgroup)
// 	if err != nil {
// 		t.Errorf("Error when executing test. %s", err)
// 	}
// 	waitgroup.Wait()

// 	actual, err := ioutil.ReadFile(testData + "actual/Jenkinsfile")
// 	if err != nil {
// 		t.Errorf("Error reading created file: %s", err.Error())
// 	}
// 	expected, err := ioutil.ReadFile(testData + "/expected/Jenkinsfile")
// 	if err != nil {
// 		t.Errorf("Error reading created file: %s", err.Error())
// 	}

// 	if !bytes.Equal(expected, actual) {
// 		t.Errorf("want:\n%s\n\n, got:\n%s\n\n", string(expected), string(actual))
// 	}
// }

// func TestGenerateCircleCI(t *testing.T) {
// 	teardown := setupTeardown(t)
// 	defer teardown(t)

// 	templates := packr.New("templates", "../../../templates")
// 	testTemplator := templator.NewTemplator(templates)

// 	var waitgroup sync.WaitGroup

// 	testConf := &projectconfig.ZeroProjectConfig{}
// 	testCI := config.CI{
// 		System:       "circleci",
// 		BuildImage:   "golang/golang",
// 		BuildTag:     "1.12",
// 		BuildCommand: "make build",
// 		TestCommand:  "make test",
// 	}

// 	err := ci.Generate(testTemplator.CI, testConf, testCI, testData+"/actual", &waitgroup)
// 	if err != nil {
// 		t.Errorf("Error when executing test. %s", err)
// 	}
// 	waitgroup.Wait()

// 	actual, err := ioutil.ReadFile(testData + "actual/.circleci/config.yml")
// 	if err != nil {
// 		t.Errorf("Error reading created file: %s", err.Error())
// 	}
// 	expected, err := ioutil.ReadFile(testData + "/expected/.circleci/config.yml")
// 	if err != nil {
// 		t.Errorf("Error reading created file: %s", err.Error())
// 	}

// 	if !bytes.Equal(expected, actual) {
// 		t.Errorf("want:\n%s\n\ngot:\n%s\n\n", string(expected), string(actual))
// 	}
// }

// func TestGenerateTravisCI(t *testing.T) {
// 	teardown := setupTeardown(t)
// 	defer teardown(t)

// 	templates := packr.New("templates", "../../../templates")
// 	testTemplator := templator.NewTemplator(templates)

// 	var waitgroup sync.WaitGroup

// 	testConf := &projectconfig.ZeroProjectConfig{}
// 	testCI := config.CI{
// 		System:       "travisci",
// 		Language:     "go",
// 		BuildImage:   "golang/golang",
// 		BuildTag:     "1.12",
// 		BuildCommand: "make build",
// 		TestCommand:  "make test",
// 	}
// 	err := ci.Generate(testTemplator.CI, testConf, testCI, testData+"/actual", &waitgroup)
// 	if err != nil {
// 		t.Errorf("Error when executing test. %s", err)
// 	}
// 	waitgroup.Wait()

// 	actual, err := ioutil.ReadFile(testData + "actual/.travis.yml")
// 	if err != nil {
// 		t.Errorf("Error reading created file: %s", err.Error())
// 	}
// 	expected, err := ioutil.ReadFile(testData + "/expected/.travis.yml")
// 	if err != nil {
// 		t.Errorf("Error reading created file: %s", err.Error())
// 	}

// 	if !bytes.Equal(expected, actual) {
// 		t.Errorf("want:\n%s\n\n, got:\n%s\n\n", string(expected), string(actual))
// 	}
// }
