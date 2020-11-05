package init

import (
	"fmt"
	"os"
	"path"
	"sync"

	"github.com/commitdev/zero/internal/config/globalconfig"
	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/module"
	"github.com/commitdev/zero/internal/registry"
	"github.com/commitdev/zero/internal/util"
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/manifoldco/promptui"
)

// Create cloud provider context
func Init(outDir string, localModulePath string) *projectconfig.ZeroProjectConfig {
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

	moduleSources := chooseStack(registry.GetRegistry(localModulePath))
	moduleConfigs, mappedSources := loadAllModules(moduleSources)

	prompts := getProjectPrompts(projectConfig.Name, moduleConfigs)

	initParams := make(map[string]string)
	projectConfig.ShouldPushRepositories = true
	initParams["ShouldPushRepositories"] = prompts["ShouldPushRepositories"].GetParam(initParams)
	if initParams["ShouldPushRepositories"] == "n" {
		projectConfig.ShouldPushRepositories = false
	}

	// Prompting for push-up stream, then conditionally prompting for github
	initParams["GithubRootOrg"] = prompts["GithubRootOrg"].GetParam(initParams)
	projectCredentials := globalconfig.GetProjectCredentials(projectConfig.Name)
	credentialPrompts := getCredentialPrompts(projectCredentials, moduleConfigs)
	projectCredentials = promptCredentialsAndFillProjectCreds(credentialPrompts, projectCredentials)
	globalconfig.Save(projectCredentials)
	projectParameters := promptAllModules(moduleConfigs, projectCredentials)

	// Map parameter values back to specific modules
	for moduleName, module := range moduleConfigs {
		repoName := prompts[moduleName].GetParam(initParams)
		repoURL := fmt.Sprintf("%s/%s", initParams["GithubRootOrg"], repoName)
		projectModuleParams := make(projectconfig.Parameters)
		projectModuleConditions := []projectconfig.Condition{}

		// Loop through all the prompted values and find the ones relevant to this module
		for parameterKey, parameterValue := range projectParameters {
			for _, moduleParameter := range module.Parameters {
				if moduleParameter.Field == parameterKey {
					projectModuleParams[parameterKey] = parameterValue
				}
			}
		}

		for _, condition := range module.Conditions {
			newCond := projectconfig.Condition{
				Action:     condition.Action,
				MatchField: condition.MatchField,
				WhenValue:  condition.WhenValue,
				Data:       condition.Data,
			}
			projectModuleConditions = append(projectModuleConditions, newCond)
		}

		projectConfig.Modules[moduleName] = projectconfig.NewModule(
			projectModuleParams,
			repoName,
			repoURL,
			mappedSources[moduleName],
			module.DependsOn,
			projectModuleConditions,
		)
	}

	// TODO: load ~/.zero/config.yml (or credentials)
	// TODO: prompt global credentials

	return &projectConfig
}

// loadAllModules takes a list of module sources, downloads those modules, and parses their config
func loadAllModules(moduleSources []string) (map[string]moduleconfig.ModuleConfig, map[string]string) {
	modules := make(map[string]moduleconfig.ModuleConfig)
	mappedSources := make(map[string]string)

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
		mappedSources[mod.Name] = moduleSource
	}
	return modules, mappedSources
}

// promptAllModules takes a map of all the modules and prompts the user for values for all the parameters
func promptAllModules(modules map[string]moduleconfig.ModuleConfig, projectCredentials globalconfig.ProjectCredential) map[string]string {
	parameterValues := map[string]string{"projectName": projectCredentials.ProjectName}
	for _, config := range modules {
		var err error

		parameterValues, err = PromptModuleParams(config, parameterValues, projectCredentials)
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
		Parameter: moduleconfig.Parameter{
			Field:   "projectName",
			Label:   "Project Name",
			Default: "",
		},
		Condition: NoCondition,
		Validate:  ValidateProjectName,
	}
}

func getProjectPrompts(projectName string, modules map[string]moduleconfig.ModuleConfig) map[string]PromptHandler {
	handlers := map[string]PromptHandler{
		"ShouldPushRepositories": {
			Parameter: moduleconfig.Parameter{
				Field:   "ShouldPushRepositories",
				Label:   "Should the created projects be checked into github automatically? (y/n)",
				Default: "y",
			},
			Condition: NoCondition,
			Validate:  SpecificValueValidation("y", "n"),
		},
		"GithubRootOrg": {
			Parameter: moduleconfig.Parameter{
				Field:   "GithubRootOrg",
				Label:   "What's the root of the github org to create repositories in?",
				Default: "github.com/",
			},
			Condition: KeyMatchCondition("ShouldPushRepositories", "y"),
			Validate:  NoValidation,
		},
	}

	for moduleName, module := range modules {
		label := fmt.Sprintf("What do you want to call the %s project?", moduleName)

		handlers[moduleName] = PromptHandler{
			Parameter: moduleconfig.Parameter{
				Field:   moduleName,
				Label:   label,
				Default: module.OutputDir,
			},
			Condition: NoCondition,
			Validate:  NoValidation,
		}
	}

	return handlers
}

func getCredentialPrompts(projectCredentials globalconfig.ProjectCredential, moduleConfigs map[string]moduleconfig.ModuleConfig) []CredentialPrompts {
	var uniqueVendors []string
	for _, module := range moduleConfigs {
		uniqueVendors = appendToSet(uniqueVendors, module.RequiredCredentials)
	}

	// map is to keep track of which vendor they belong to, to fill them back into the projectConfig
	prompts := []CredentialPrompts{}
	for _, vendor := range AvailableVendorOrders {
		if util.ItemInSlice(uniqueVendors, vendor) {
			vendorPrompts := CredentialPrompts{vendor, mapVendorToPrompts(projectCredentials, vendor)}
			prompts = append(prompts, vendorPrompts)
		}
	}
	return prompts
}

func mapVendorToPrompts(projectCred globalconfig.ProjectCredential, vendor string) []PromptHandler {
	var prompts []PromptHandler
	profiles, err := project.GetAWSProfiles()
	if err != nil {
		profiles = []string{}
	}

	// if no profiles available, dont prompt use to pick profile
	customAwsPickProfileCondition := func(param map[string]string) bool {
		if len(profiles) == 0 {
			flog.Infof(":warning: No AWS profiles found, please manually input AWS credentials")
			return false
		} else {
			return true
		}
	}

	// condition for prompting manual AWS credentials input
	customAwsMustInputCondition := func(param map[string]string) bool {
		toPickProfile := awsPickProfile
		if val, ok := param["use_aws_profile"]; ok && val != toPickProfile {
			return true
		}
		return false
	}

	switch vendor {
	case "aws":
		awsPrompts := []PromptHandler{
			{
				Parameter: moduleconfig.Parameter{
					Field:   "use_aws_profile",
					Label:   "Use credentials from existing AWS profiles?",
					Options: []string{awsPickProfile, awsManualInputCredentials},
				},
				Condition: customAwsPickProfileCondition,
				Validate:  NoValidation,
			},
			{
				Parameter: moduleconfig.Parameter{
					Field:   "aws_profile",
					Label:   "Select AWS Profile",
					Options: profiles,
				},
				Condition: KeyMatchCondition("use_aws_profile", awsPickProfile),
				Validate:  NoValidation,
			},
			{
				Parameter: moduleconfig.Parameter{
					Field:   "accessKeyId",
					Label:   "AWS Access Key ID",
					Default: projectCred.AWSResourceConfig.AccessKeyID,
					Info: `AWS Access Key ID/Secret: used for provisioning infrastructure in AWS
The token can be generated at https://console.aws.amazon.com/iam/home?#/security_credentials`,
				},
				Condition: CustomCondition(customAwsMustInputCondition),
				Validate:  ValidateAKID,
			},
			{
				Parameter: moduleconfig.Parameter{
					Field:   "secretAccessKey",
					Label:   "AWS Secret access key",
					Default: projectCred.AWSResourceConfig.SecretAccessKey,
				},
				Condition: CustomCondition(customAwsMustInputCondition),
				Validate:  ValidateSAK,
			},
		}
		prompts = append(prompts, awsPrompts...)
	case "github":
		githubPrompt := PromptHandler{
			Parameter: moduleconfig.Parameter{
				Field:   "accessToken",
				Label:   "Github Personal Access Token with access to the above organization",
				Default: projectCred.GithubResourceConfig.AccessToken,
				Info: `Github personal access token: used for creating repositories for your project
Requires the following permissions: [repo::public_repo, admin::orgread:org]
The token can be created at https://github.com/settings/tokens`,
			},
			Condition: NoCondition,
			Validate:  NoValidation,
		}
		prompts = append(prompts, githubPrompt)
	case "circleci":
		circleCiPrompt := PromptHandler{
			Parameter: moduleconfig.Parameter{
				Field:   "apiKey",
				Label:   "Circleci api key for CI/CD",
				Default: projectCred.CircleCiResourceConfig.ApiKey,
				Info: `CircleCI api token: used for setting up CI/CD for your project
The token can be created at https://app.circleci.com/settings/user/tokens`,
			},
			Condition: NoCondition,
			Validate:  NoValidation,
		}
		prompts = append(prompts, circleCiPrompt)
	}
	return prompts
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

func chooseStack(reg registry.Registry) []string {
	providerPrompt := promptui.Select{
		Label: "Pick a stack you'd like to use",
		Items: registry.AvailableLabels(reg),
	}
	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Prompt failed %v\n", err)
	}

	return registry.GetModulesByName(reg, providerResult)
}

func defaultProjConfig() projectconfig.ZeroProjectConfig {
	return projectconfig.ZeroProjectConfig{
		Name:       "",
		Parameters: map[string]string{},
		Modules:    projectconfig.Modules{},
	}
}
