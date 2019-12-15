package cmd

import (
	"log"
	"os"
	"path"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
	"github.com/commitdev/commit0/internal/util/secrets"
	"github.com/gobuffalo/packr/v2"
	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
}

func Create(projectName string, outDir string, t *templator.Templator) string {
	rootDir := path.Join(outDir, projectName)
	log.Println(aurora.Cyan(emoji.Sprintf(":tada: Creating project %s.", projectName)))
	err := os.MkdirAll(rootDir, os.ModePerm)
	if os.IsExist(err) {
		log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Directory %v already exists! Error: %v", projectName, err)))
	} else if err != nil {
		log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Error creating root: %v ", err)))
	}

	projectConfig := defaultProjConfig(projectName)

	chooseCloudProvider(&projectConfig)

	s := secrets.GetSecrets(rootDir)

	fillProviderDetails(&projectConfig, s)

	var wg sync.WaitGroup
	util.TemplateFileIfDoesNotExist(rootDir, util.CommitYml, t.Commit0, &wg, projectConfig)
	util.TemplateFileIfDoesNotExist(rootDir, ".gitignore", t.GitIgnore, &wg, projectName)

	wg.Wait()
	return rootDir
}

func chooseCloudProvider(projectConfig *util.ProjectConfiguration) {
	providerPrompt := promptui.Select{
		Label: "Select Cloud Provider",
		Items: []string{"Amazon AWS", "Google GCP", "Microsoft Azure"},
	}

	_, providerResult, err := providerPrompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
		panic(err)
	}

	if providerResult == "Amazon AWS" {
		// @TODO : Move this stuff from util into another package
		projectConfig.Infrastructure.AWS = &util.AWS{}
		regionPrompt := promptui.Select{
			Label: "Select AWS Region ",
			Items: []string{"us-west-1", "us-west-2", "us-east-1", "us-east-2", "ca-central-1",
				"eu-central-1", "eu-west-1", "ap-east-1", "ap-south-1"},
		}

		_, regionResult, err := regionPrompt.Run()

		if err != nil {
			log.Fatalf("Prompt failed %v\n", err)
			panic(err)
		}

		projectConfig.Infrastructure.AWS.Region = regionResult
	} else {
		log.Fatalf("Only the AWS provider is available at this time")
	}
}

func fillProviderDetails(projectConfig *util.ProjectConfiguration, s secrets.Secrets) {
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
					log.Fatalf(aerr.Error())
				}
			} else {
				log.Fatalf(err.Error())
			}
		}

		if awsCaller != nil && awsCaller.Account != nil {
			projectConfig.Infrastructure.AWS.AccountID = *awsCaller.Account
		}
	}
}

func defaultProjConfig(projectName string) util.ProjectConfiguration {
	return util.ProjectConfiguration{
		ProjectName:       projectName,
		FrontendFramework: "react",
		Organization:      "mycompany",
		Description:       "",
		Maintainers: []util.Maintainer{{
			Name:  "bob",
			Email: "bob@test.com",
		}},
		Services: []util.Service{{
			Name:        "User",
			Description: "User Service",
			Language:    "go",
			GitRepo:     "github.com/test/repo",
		}},
		Infrastructure: util.Infrastructure{
			AWS: nil,
		},
	}
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new project with provided name.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalln(aurora.Red(emoji.Sprintf(":exclamation: Project name cannot be empty!")))
		}

		templates := packr.New("templates", "../templates")
		t := templator.NewTemplator(templates)

		projectName := args[0]

		Create(projectName, "./", t)
	},
}
