package init

import (
	"github.com/commitdev/zero/internal/config/moduleconfig"
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/flog"
)

func CustomPromptHandler(promptType string, params map[string]string) {
	switch promptType {

	case "AWSProfilePicker":
		AWSProfilePicker(params)
	}
}

func AWSProfilePicker(params map[string]string) {
	profiles, err := project.GetAWSProfiles()
	if err != nil {
		profiles = []string{}
	}

	awsPrompt := PromptHandler{
		Parameter: moduleconfig.Parameter{
			Field:   "aws_profile",
			Label:   "Select AWS Profile",
			Options: profiles,
		},
		Condition: NoCondition,
		Validate:  NoValidation,
	}
	_, value := promptParameter(awsPrompt)
	credErr := project.FillAWSProfile(value, params)
	if credErr != nil {
		flog.Errorf("Failed to retrieve profile, falling back to User input")
		params["useExistingAwsProfile"] = "no"
	}
}
