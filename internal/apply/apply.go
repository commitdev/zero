package apply

import (
	"errors"
	"fmt"
	"path/filepath"

	"log"
	"os/exec"
	"path"
	"strings"

	"github.com/commitdev/zero/internal/module"
	"github.com/commitdev/zero/internal/util"
	"github.com/hashicorp/terraform/dag"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/manifoldco/promptui"
)

func Apply(rootDir string, configPath string, environments []string) error {
	var err error
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

	flog.Infof(":mag: checking project %s's module requirements.", projectConfig.Name)

	err = modulesWalkCmd("check", rootDir, projectConfig, []string{"make", "check"}, environments)
	if err != nil {
		return errors.New(fmt.Sprintf("Module checks failed: %s", err.Error()))
	}

	flog.Infof(":tada: Bootstrapping project %s. Please use the zero-project.yml file to modify the project as needed.", projectConfig.Name)

	flog.Infof("Cloud provider: %s", "AWS") // will this come from the config?

	flog.Infof("Runtime platform: %s", "Kubernetes")

	flog.Infof("Infrastructure executor: %s", "Terraform")

	err = modulesWalkCmd("apply", rootDir, projectConfig, []string{"make"}, environments)
	if err != nil {
		return errors.New(fmt.Sprintf("Module Apply failed: %s", err.Error()))
	}

	flog.Infof(":check_mark_button: Done.")

	flog.Infof("Your projects and infrastructure have been successfully created.  Here are some useful links and commands to get you started:")
	err = modulesWalkCmd("summary", rootDir, projectConfig, []string{"make", "summary"}, environments)
	if err != nil {
		return errors.New(fmt.Sprintf("Module summary failed: %s", err.Error()))
	}
	return err
}

func modulesWalkCmd(lifecycleName string, dir string, projectConfig *projectconfig.ZeroProjectConfig, cmdArgs []string, environments []string) error {
	graph := projectConfig.GetDAG()
	root := []dag.Vertex{projectconfig.GraphRootName}
	environmentArg := fmt.Sprintf("ENVIRONMENT=%s", strings.Join(environments, ","))
	err := graph.DepthFirstWalk(root, func(v dag.Vertex, depth int) error {
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

		envVarTranslationMap := modConfig.GetParamEnvVarTranslationMap()
		envList = util.AppendProjectEnvToCmdEnv(mod.Parameters, envList, envVarTranslationMap)
		flog.Debugf("Env injected: %#v", envList)

		// only print msg for apply, or else it gets a little spammy
		if lifecycleName == "apply" {
			flog.Infof("Executing %s command for %s...", lifecycleName, modConfig.Name)
		}

		execErr := util.ExecuteCommand(exec.Command(cmdArgs[0], cmdArgs[1:]...), modulePath, envList)
		if execErr != nil {
			return errors.New(fmt.Sprintf("Module (%s) %s", modConfig.Name, execErr.Error()))
		}
		return nil
	})

	return err
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
