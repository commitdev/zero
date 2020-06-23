package apply

import (
	"fmt"
	"path/filepath"

	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/commitdev/zero/internal/module"
	"github.com/commitdev/zero/internal/util"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/manifoldco/promptui"
)

func Apply(rootDir string, configPath string, environments []string) {
	if strings.Trim(configPath, " ") == "" {
		exit.Fatal("config path cannot be empty!")
	}
	configFilePath := path.Join(rootDir, configPath)
	projectConfig := projectconfig.LoadConfig(configFilePath)

	if len(environments) == 0 {
		fmt.Println(`Choose the environments to apply. This will create infrastructure, CI pipelines, etc.
At this point, real things will be generated that may cost money!
Only a single environment may be suitable for an initial test, but for a real system we suggest setting up both staging and production environments.`)
		environments = promptEnvironments()
	}

	validateEnvironments(environments)

	flog.Infof(":tada: Bootstrapping project %s. Please use the zero.yml file to modify the project as needed.", projectConfig.Name)

	flog.Infof("Cloud provider: %s", "AWS") // will this come from the config?

	flog.Infof("Runtime platform: %s", "Kubernetes")

	flog.Infof("Infrastructure executor: %s", "Terraform")

	applyAll(rootDir, *projectConfig, environments)

	// TODO Summary
	flog.Infof(":check_mark_button: Done - Summary goes here.")
}

func applyAll(dir string, projectConfig projectconfig.ZeroProjectConfig, applyEnvironments []string) {
	environmentArg := fmt.Sprintf("ENVIRONMENT=%s", strings.Join(applyEnvironments, ","))

	for _, mod := range projectConfig.Modules {
		dirArg := fmt.Sprintf("PROJECT_DIR=%s", path.Join(dir, mod.Files.Directory))
		envList := []string{environmentArg, dirArg}

		modulePath := module.GetSourceDir(mod.Files.Source)
		// Passed in `dir` will only be used to find the project path, not the module path,
		// unless the module path is relative
		if module.IsLocal(mod.Files.Source) && !filepath.IsAbs(modulePath) {
			modulePath = filepath.Join(dir, modulePath)
		}

		envList = util.AppendProjectEnvToCmdEnv(mod.Parameters, envList)
		util.ExecuteCommand(exec.Command("make"), modulePath, envList)
	}
}

// promptEnvironments Prompts the user for the environments to apply against and returns a slice of strings representing the environments
func promptEnvironments() []string {
	items := map[string][]string{
		"Staging ":                    {"staging"},
		"Production":                  {"production"},
		"Both Staging and Production": {"staging", "production"},
	}

	var labels []string
	for label := range items {
		labels = append(labels, label)
	}

	providerPrompt := promptui.Select{
		Label: "Environments",
		Items: labels,
	}
	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}
	return items[providerResult]
}

func validateEnvironments(applyEnvironments []string) {
	// Strict for now, we can brainstorm how much we want to support custom environments later
	for _, env := range applyEnvironments {
		if env != "staging" && env != "production" {
			exit.Fatal("The currently supported environments are \"staging\" and \"production\"")
		}
	}
}
