package context

import (
	"log"
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
	promptProjectName(projectName, &projectConfig)
	chooseStack(&projectConfig)

	// TODO: load ~/.zero/config.yml (or credentials)
	// TODO: prompt global credentials

	// chooseCloudProvider(&projectConfig)
	// fmt.Println(&projectConfig)
	// s := project.GetSecrets(rootDir)
	// fillProviderDetails(&projectConfig, s)
	// fmt.Println(&projectConfig)

	promptAllModules(&projectConfig)

	return &projectConfig
}

func promptAllModules(projectConfig *config.ZeroProjectConfig) {
	// TODO: do we need to run through the modules and extract first
	// or we need to run through twice, potentially still need to pre-process for global auths
	for _, moduleSource := range projectConfig.Modules {
		mod, _ := module.NewTemplateModule(config.ModuleInstance{Source: moduleSource})
		err := mod.PromptParams(projectConfig.Context)
		if err != nil {
			log.Fatalf("Exiting prompt:  %v\n", err)
			panic(err)
		}
	}
}

func promptProjectName(projectName string, projectConfig *config.ZeroProjectConfig) {
	providerPrompt := promptui.Prompt{
		Label:     "Project Name",
		Default:   projectName,
		AllowEdit: false,
	}
	providerPrompt.Run()
}

func chooseCloudProvider(projectConfig *config.ZeroProjectConfig) {
	// @TODO move options into configs
	providerPrompt := promptui.Select{
		Label: "Select Cloud Provider",
		Items: []string{"Amazon AWS", "Google GCP", "Microsoft Azure"},
	}

	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	if providerResult != "Amazon AWS" {
		exit.Fatal("Only the AWS provider is available at this time")
	}
}

func chooseStack(projectConfig *config.ZeroProjectConfig) {
	items := map[string][]string{
		// TODO: better place to store these options as configuration file or any source
		"EKS + Go + React": []string{
			"github.com/commitdev/zero-aws-eks-stack",
			"github.com/commitdev/zero-deployable-backend",
			"github.com/commitdev/zero-deployable-react-frontend",
		},
		"Custom": []string{},
	}

	labels := make([]string, len(items))
	i := 0
	for label := range items {
		labels[i] = label
		i++
	}

	providerPrompt := promptui.Select{
		Label: "Pick a stack you'd like to use",
		Items: labels,
	}
	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}
	projectConfig.Modules = items[providerResult]
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
