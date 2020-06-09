package context

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/commitdev/zero/internal/config/globalconfig"
	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/module"
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/manifoldco/promptui"
)

type Registry map[string][]string

// Create cloud provider context
func Init(outDir string) *projectconfig.ZeroProjectConfig {
	projectConfig := defaultProjConfig()

	projectConfig.Name = getProjectNamePrompt().GetParam(projectConfig.Parameters)

	rootDir := path.Join(outDir, projectConfig.Name)
	flog.Infof(":tada: Initializing project")

	err := os.MkdirAll(rootDir, os.ModePerm)
	if os.IsExist(err) {
		exit.Fatal("Directory %v already exists! Error: %v", projectConfig.Name, err)
	} else if err != nil {
		exit.Fatal("Error creating root: %v ", err)
	}

	moduleSources := chooseStack(getRegistry())
	moduleConfigs := loadAllModules(moduleSources)
	for _ = range moduleConfigs {
		// TODO: initialize module structs inside project
	}

	prompts := getProjectPrompts(projectConfig.Name, moduleConfigs)

	initParams := make(map[string]string)
	projectConfig.ShouldPushRepositories = true
	initParams["ShouldPushRepositories"] = prompts["ShouldPushRepositories"].GetParam(initParams)
	if initParams["ShouldPushRepositories"] == "n" {
		projectConfig.ShouldPushRepositories = false
	}

	// Prompting for push-up stream, then conditionally prompting for github
	initParams["GithubRootOrg"] = prompts["GithubRootOrg"].GetParam(initParams)
	initParams["GithubPersonalToken"] = prompts["GithubPersonalToken"].GetParam(initParams)
	if initParams["GithubRootOrg"] != "" && initParams["GithubPersonalToken"] != globalconfig.GetUserCredentials(projectConfig.Name).AccessToken {
		projectCredential := globalconfig.GetUserCredentials(projectConfig.Name)
		projectCredential.GithubResourceConfig.AccessToken = initParams["GithubPersonalToken"]
		globalconfig.Save(projectCredential)
	}

	projectParameters := promptAllModules(moduleConfigs)
	for k, v := range projectParameters {
		projectConfig.Parameters[k] = v
		// TODO: Add parameters to module structs inside project
	}

	for moduleName, _ := range moduleConfigs {
		// @TODO : Uncomment when this struct is implemented
		repoName := prompts[moduleName].GetParam(initParams)
		repoURL := fmt.Sprintf("%s/%s", initParams["GithubRootOrg"], repoName)
		//projectConfig.Modules[moduleName].Files.Directory = prompts[moduleName].GetParam(initParams)
		//projectConfig.Modules[moduleName].Files.Repository = repoURL
		fmt.Println(repoURL)
	}

	// TODO: load ~/.zero/config.yml (or credentials)
	// TODO: prompt global credentials

	return &projectConfig
}

// loadAllModules takes a list of module sources, downloads those modules, and parses their config
func loadAllModules(moduleSources []string) map[string]moduleconfig.ModuleConfig {
	modules := make(map[string]moduleconfig.ModuleConfig)

	wg := sync.WaitGroup{}
	wg.Add(len(moduleSources))
	for _, moduleSource := range moduleSources {
		go module.FetchModule(moduleSource, &wg)
	}
	wg.Wait()

	for _, moduleSource := range moduleSources {
		mod, err := module.ParseModuleConfig(moduleSource)
		if err != nil {
			exit.Fatal("Unable to load module:  %v\n", err)
		}
		modules[mod.Name] = mod
	}
	return modules
}

// promptAllModules takes a map of all the modules and prompts the user for values for all the parameters
func promptAllModules(modules map[string]moduleconfig.ModuleConfig) map[string]string {
	parameterValues := make(map[string]string)
	for _, config := range modules {
		var err error
		parameterValues, err = PromptModuleParams(config, parameterValues)
		if err != nil {
			exit.Fatal("Exiting prompt:  %v\n", err)
		}
	}
	return parameterValues
}

// Project name is prompt individually because the rest of the prompts
// requires the projectName to populate defaults
func getProjectNamePrompt() PromptHandler {
	return PromptHandler{
		moduleconfig.Parameter{
			Field:   "projectName",
			Label:   "Project Name",
			Default: "",
		},
		NoCondition,
		NoValidation,
	}
}

func getProjectPrompts(projectName string, modules map[string]moduleconfig.ModuleConfig) map[string]PromptHandler {
	handlers := map[string]PromptHandler{
		"ShouldPushRepositories": {
			moduleconfig.Parameter{
				Field:   "ShouldPushRepositories",
				Label:   "Should the created projects be checked into github automatically? (y/n)",
				Default: "y",
			},
			NoCondition,
			SpecificValueValidation("y", "n"),
		},
		"GithubRootOrg": {
			moduleconfig.Parameter{
				Field:   "GithubRootOrg",
				Label:   "What's the root of the github org to create repositories in?",
				Default: "github.com/",
			},
			KeyMatchCondition("ShouldPushRepositories", "y"),
			NoValidation,
		},
		"GithubPersonalToken": {
			moduleconfig.Parameter{
				Field:   "GithubPersonalToken",
				Label:   "Github Personal Access Token with access to the above organization",
				Default: globalconfig.GetUserCredentials(projectName).AccessToken,
			},
			KeyMatchCondition("ShouldPushRepositories", "y"),
			NoValidation,
		},
	}

	for moduleName, module := range modules {
		label := fmt.Sprintf("What do you want to call the %s project?", moduleName)

		handlers[moduleName] = PromptHandler{
			moduleconfig.Parameter{
				Field:   moduleName,
				Label:   label,
				Default: module.OutputDir,
			},
			NoCondition,
			NoValidation,
		}
	}

	return handlers
}

func chooseCloudProvider(projectConfig *projectconfig.ZeroProjectConfig) {
	// @TODO move options into configs
	providerPrompt := promptui.Select{
		Label: "Select Cloud Provider",
		Items: []string{"Amazon AWS", "Google GCP", "Microsoft Azure"},
	}

	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Prompt failed %v\n", err)
	}

	if providerResult != "Amazon AWS" {
		exit.Fatal("Only the AWS provider is available at this time")
	}
}

func getRegistry() Registry {
	return Registry{
		// TODO: better place to store these options as configuration file or any source
		"EKS + Go + React": []string{
			"github.com/commitdev/zero-aws-eks-stack",
			"github.com/commitdev/zero-deployable-backend",
			"github.com/commitdev/zero-deployable-react-frontend",
		},
		"Custom": []string{},
	}
}

func (registry Registry) availableLabels() []string {
	labels := make([]string, len(registry))
	i := 0
	for label := range registry {
		labels[i] = label
		i++
	}
	return labels
}

func chooseStack(registry Registry) []string {
	providerPrompt := promptui.Select{
		Label: "Pick a stack you'd like to use",
		Items: registry.availableLabels(),
	}
	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Prompt failed %v\n", err)
	}
	return registry[providerResult]

}

func fillProviderDetails(projectConfig *projectconfig.ZeroProjectConfig, s project.Secrets) {
	if projectConfig.Infrastructure.AWS != nil {
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(projectConfig.Infrastructure.AWS.Region),
			Credentials: credentials.NewStaticCredentials(s.AWS.AccessKeyID, s.AWS.SecretAccessKey, ""),
		})

		svc := sts.New(sess)
		input := &sts.GetCallerIdentityInput{}

		awsCaller, err := svc.GetCallerIdentity(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					exit.Error(aerr.Error())
				}
			} else {
				exit.Error(err.Error())
			}
		}

		if awsCaller != nil && awsCaller.Account != nil {
			projectConfig.Infrastructure.AWS.AccountID = *awsCaller.Account
		}
	}
}

func defaultProjConfig() projectconfig.ZeroProjectConfig {
	return projectconfig.ZeroProjectConfig{
		Name: "",
		Infrastructure: projectconfig.Infrastructure{
			AWS: nil,
		},
		Parameters: map[string]string{},
		Modules:    []string{},
	}
}
