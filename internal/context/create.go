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
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/manifoldco/promptui"
)

// Create cloud provider context
func Create(projectName string, outDir string) string {
	rootDir := path.Join(outDir, projectName)
	flog.Infof(":tada: Creating project %s.", projectName)

	err := os.MkdirAll(rootDir, os.ModePerm)
	if os.IsExist(err) {
		exit.Fatal("Directory %v already exists! Error: %v", projectName, err)
	} else if err != nil {
		exit.Fatal("Error creating root: %v ", err)
	}

	projectConfig := defaultProjConfig(projectName)
	chooseCloudProvider(&projectConfig)

	s := project.GetSecrets(rootDir)
	fillProviderDetails(&projectConfig, s)

	return rootDir
}

func chooseCloudProvider(projectConfig *config.Commit0Config) {
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

func fillProviderDetails(projectConfig *config.Commit0Config, s project.Secrets) {
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

func defaultProjConfig(projectName string) config.Commit0Config {
	return config.Commit0Config{
		Name: projectName,
		Infrastructure: config.Infrastructure{
			AWS: nil,
		},
	}
}
