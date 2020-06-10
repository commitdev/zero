package apply

import (
	"fmt"

	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/commitdev/zero/internal/util"
	"github.com/commitdev/zero/pkg/util/flog"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/manifoldco/promptui"
)

// Apply will bootstrap the runtime environment for the project
func Apply(dir string, applyConfigPath string, applyEnvironments []string) []string {
	context := loadContext(dir, applyConfigPath, applyEnvironments)

	flog.Infof(":tada: Bootstrapping project %s. Please use the zero.[hcl, yaml] file to modify the project as needed. %s.", context.Name)

	flog.Infof("Cloud provider: %s", "AWS") // will this come from the config?

	flog.Infof("Runtime platform: %s", "Kubernetes")

	flog.Infof("Infrastructure executor: %s", "Terraform")

	// other details...

	return makeAll(dir, context, applyEnvironments)
}

// loadContext will load the context/configuration to be used by the apply command
func loadContext(dir string, applyConfigPath string, applyEnvironments []string) *projectconfig.ZeroProjectConfig {
	if len(applyEnvironments) == 0 {
		fmt.Println(`Choose the environments to apply. This will create infrastructure, CI pipelines, etc.
At this point, real things will be generated that may cost money!
Only a single environment may be suitable for an initial test, but for a real system we suggest setting up both staging and production environments.`)
		applyEnvironments = promptEnvironments()
	}

	validateEnvironments(applyEnvironments)

	if applyConfigPath == "" {
		exit.Fatal("config path cannot be empty!")
	}
	configPath := path.Join(dir, applyConfigPath)
	projectConfig := projectconfig.LoadConfig(configPath)
	return projectConfig
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

func makeAll(dir string, projectContext *projectconfig.ZeroProjectConfig, applyEnvironments []string) []string {
	environmentArg := fmt.Sprintf("ENVIRONMENT=%s", strings.Join(applyEnvironments, ","))
	envList := []string{environmentArg}
	outputs := []string{}

	for _, mod := range projectContext.Modules {
		modulePath := path.Join(dir, mod.Files.Directory)
		envList = util.AppendProjectEnvToCmdEnv(mod.Parameters, envList)

		output := util.ExecuteCommandOutput(exec.Command("make"), modulePath, envList)
		outputs = append(outputs, output)
	}
	return outputs
}
