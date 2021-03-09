package init

import (
	"errors"
	"fmt"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	project "github.com/commitdev/zero/pkg/credentials"
	"github.com/commitdev/zero/pkg/util/flog"
)

// CustomPromptHandler handles non-input and enum options prompts
// zero-module's parameters allow prompts to specify types of custom actions
// this allows non-standard enum / string input to be added, such as AWS profile picker
func CustomPromptHandler(promptType string, params map[string]string) error {
	switch promptType {

	case "AWSProfilePicker":
		promptAWSProfilePicker(params)
	default:
		return errors.New(fmt.Sprintf("Unsupported custom prompt type %s.", promptType))
	}
	return nil
}

func promptAWSProfilePicker(params map[string]string) {
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
