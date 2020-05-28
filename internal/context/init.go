package context

import (
	"os"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/commitdev/zero/internal/config"
	"github.com/commitdev/zero/internal/module"
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/manifoldco/promptui"
)

type Registry map[string][]string

// Create cloud provider context
func Init(projectName string, outDir string) *config.ZeroProjectConfig {
	rootDir := path.Join(outDir, projectName)
	flog.Infof(":tada: Creating project %s.", projectName)

	err := os.MkdirAll(rootDir, os.ModePerm)
	if os.IsExist(err) {
		exit.Fatal("Directory %v already exists! Error: %v", projectName, err)
	} else if err != nil {
		exit.Fatal("Error creating root: %v ", err)
	}

	projectConfig := defaultProjConfig(projectName)
	projectConfig.Name = promptProjectName(projectName)
	projectConfig.Context["ShouldPushRepoUpstream"] = promptPushRepoUpstream()
	projectConfig.Context["GithubRootOrg"] = promptGithubRootOrg()
	projectConfig.Context["githubPersonalToken"] = promptGithubPersonalToken(projectName)
	projectConfig.Modules = chooseStack(getRegistry())

	// chooseCloudProvider(&projectConfig)
	// fmt.Println(&projectConfig)
	// s := project.GetSecrets(rootDir)
	// fillProviderDetails(&projectConfig, s)
	// fmt.Println(&projectConfig)

	promptAllModules(&projectConfig)
	// TODO: load ~/.zero/config.yml (or credentials)
	// TODO: prompt global credentials

	return &projectConfig
}

func promptAllModules(projectConfig *config.ZeroProjectConfig) {
	// TODO: do we need to run through the modules and extract first
	// or we need to run through twice, potentially still need to pre-process for global auths
	for _, moduleSource := range projectConfig.Modules {
		mod, _ := module.NewTemplateModule(config.ModuleInstance{Source: moduleSource})
		err := mod.PromptParams(projectConfig.Context)
		if err != nil {
			exit.Fatal("Exiting prompt:  %v\n", err)
		}
	}
}

// global configs
func promptPushRepoUpstream() string {
	providerPrompt := promptui.Prompt{
		Label:     "Should the created projects be checked into github automatically? (y/n)",
		Default:   "y",
		AllowEdit: false,
	}
	providerResult, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Exiting prompt:  %v\n", err)
	}
	return providerResult
}

func promptGithubRootOrg() string {
	providerPrompt := promptui.Prompt{
		Label:     "What's the root of the github org to create repositories in?",
		Default:   "github.com/",
		AllowEdit: true,
	}
	result, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Exiting prompt:  %v\n", err)
	}
	return result
}

func promptGithubPersonalToken(projectName string) string {
	defaultToken := ""

	project := config.GetUserCredentials(projectName)
	if project.GithubResourceConfig.AccessToken != "" {
		defaultToken = project.GithubResourceConfig.AccessToken
	}

	providerPrompt := promptui.Prompt{
		Label:   "Github Personal Access Token with access to the above organization",
		Default: defaultToken,
	}
	result, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Prompt failed %v\n", err)
	}

	// If its different from saved token, update it
	if project.GithubResourceConfig.AccessToken != result {
		project.GithubResourceConfig.AccessToken = result
		config.Save(project)
	}
	return result
}

func promptProjectName(projectName string) string {
	providerPrompt := promptui.Prompt{
		Label:     "Project Name",
		Default:   projectName,
		AllowEdit: false,
	}
	result, err := providerPrompt.Run()
	if err != nil {
		exit.Fatal("Prompt failed %v\n", err)
	}
	return result
}

func chooseCloudProvider(projectConfig *config.ZeroProjectConfig) {
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

func fillProviderDetails(projectConfig *config.ZeroProjectConfig, s project.Secrets) {
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

func defaultProjConfig(projectName string) config.ZeroProjectConfig {
	return config.ZeroProjectConfig{
		Name: projectName,
		Infrastructure: config.Infrastructure{
			AWS: nil,
		},
		Context: map[string]string{},
		Modules: []string{},
	}
}
