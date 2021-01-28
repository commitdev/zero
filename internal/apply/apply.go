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
	"github.com/hashicorp/terraform/dag"

	"github.com/commitdev/zero/internal/config/globalconfig"
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

	flog.Infof(":tada: Bootstrapping project %s. Please use the zero-project.yml file to modify the project as needed.", projectConfig.Name)

	flog.Infof("Cloud provider: %s", "AWS") // will this come from the config?

	flog.Infof("Runtime platform: %s", "Kubernetes")

	flog.Infof("Infrastructure executor: %s", "Terraform")

	applyAll(rootDir, *projectConfig, environments)

	flog.Infof(":check_mark_button: Done.")

	summarizeAll(rootDir, *projectConfig, environments)
}

func applyAll(dir string, projectConfig projectconfig.ZeroProjectConfig, applyEnvironments []string) {
	environmentArg := fmt.Sprintf("ENVIRONMENT=%s", strings.Join(applyEnvironments, ","))

	graph := projectConfig.GetDAG()

	// Walk the graph of modules and run `make`
	root := []dag.Vertex{projectconfig.GraphRootName}
	graph.DepthFirstWalk(root, func(v dag.Vertex, depth int) error {
		// Don't process the root
		if depth == 0 {
			return nil
		}

		name := v.(string)
		mod := projectConfig.Modules[name]
		// Add env vars for the makefile
		envList := []string{
			environmentArg,
			fmt.Sprintf("PROJECT_NAME=%s", projectConfig.Name),
			fmt.Sprintf("PROJECT_DIR=%s", path.Join(dir, mod.Files.Directory)),
			fmt.Sprintf("REPOSITORY=%s", mod.Files.Repository),
		}

		modulePath := module.GetSourceDir(mod.Files.Source)
		// Passed in `dir` will only be used to find the project path, not the module path,
		// unless the module path is relative
		if module.IsLocal(mod.Files.Source) && !filepath.IsAbs(modulePath) {
			modulePath = filepath.Join(dir, modulePath)
		}
		flog.Debugf("Loaded module: %s from %s", name, modulePath)

		// TODO: in the case user lost the `/tmp` (module source dir), this will fail
		// and we should redownload the module for the user
		modConfig, err := module.ParseModuleConfig(modulePath)
		if err != nil {
			exit.Fatal("Failed to load module config, credentials cannot be injected properly")
		}

		// Get project credentials for the makefile
		credentials := globalconfig.GetProjectCredentials(projectConfig.Name)
		credentialEnvs := credentials.SelectedVendorsCredentialsAsEnv(modConfig.RequiredCredentials)
		envList = util.AppendProjectEnvToCmdEnv(mod.Parameters, envList)
		envList = util.AppendProjectEnvToCmdEnv(credentialEnvs, envList)
		flog.Debugf("Env injected: %#v", envList)
		flog.Infof("Executing apply command for %s...", modConfig.Name)
		util.ExecuteCommand(exec.Command("make"), modulePath, envList)
		return nil
	})
}

// promptEnvironments Prompts the user for the environments to apply against and returns a slice of strings representing the environments
func promptEnvironments() []string {
	items := map[string][]string{
		"Staging":                     {"stage"},
		"Production":                  {"prod"},
		"Both Staging and Production": {"stage", "prod"},
	}

	labels := []string{"Staging", "Production", "Both Staging and Production"}

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

func summarizeAll(dir string, projectConfig projectconfig.ZeroProjectConfig, applyEnvironments []string) {
	flog.Infof("Your projects and infrastructure have been successfully created.  Here are some useful links and commands to get you started:")

	graph := projectConfig.GetDAG()

	// Walk the graph of modules and run `make summary`
	root := []dag.Vertex{projectconfig.GraphRootName}
	graph.DepthFirstWalk(root, func(v dag.Vertex, depth int) error {
		// Don't process the root
		if depth == 0 {
			return nil
		}

		name := v.(string)
		mod := projectConfig.Modules[name]
		// Add env vars for the makefile
		envList := []string{
			fmt.Sprintf("ENVIRONMENT=%s", strings.Join(applyEnvironments, ",")),
			fmt.Sprintf("REPOSITORY=%s", mod.Files.Repository),
			fmt.Sprintf("PROJECT_NAME=%s", projectConfig.Name),
		}

		modulePath := module.GetSourceDir(mod.Files.Source)
		// Passed in `dir` will only be used to find the project path, not the module path,
		// unless the module path is relative
		if module.IsLocal(mod.Files.Source) && !filepath.IsAbs(modulePath) {
			modulePath = filepath.Join(dir, modulePath)
		}
		flog.Debugf("Loaded module: %s from %s", name, modulePath)

		envList = util.AppendProjectEnvToCmdEnv(mod.Parameters, envList)
		flog.Debugf("Env injected: %#v", envList)
		util.ExecuteCommand(exec.Command("make", "summary"), modulePath, envList)
		return nil
	})

	flog.Infof("Happy coding! :smile:")
}
