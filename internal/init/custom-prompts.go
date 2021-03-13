package init

import (
	"errors"
	"fmt"

	"github.com/commitdev/zero/internal/config/moduleconfig"
	project "github.com/commitdev/zero/pkg/credentials"
)

// CustomPromptHandler handles non-input and enum options prompts
// zero-module's parameters allow prompts to specify types of custom actions
// this allows non-standard enum / string input to be added, such as AWS profile picker
func CustomPromptHandler(promptType string, params map[string]string) error {
	switch promptType {

	case "AWSProfilePicker":
		err := promptAWSProfilePicker(params)
		if err != nil {
			params["useExistingAwsProfile"] = "no"
			return err
		}
	default:
		return errors.New(fmt.Sprintf("Unsupported custom prompt type %s.", promptType))
	}
	return nil
}

func promptAWSProfilePicker(params map[string]string) error {
	profiles, err := project.GetAWSProfiles()
	if err != nil {
		return err
	}

	awsPrompt := PromptHandler{
		Parameter: moduleconfig.Parameter{
			Field:   "aws_profile",
			Label:   "Select AWS Profile",
			Options: listToPromptOptions(profiles),
		},
		Condition: NoCondition,
		Validate:  NoValidation,
	}
	_, value := promptParameter(awsPrompt)
	credErr := project.FillAWSProfile("", value, params)
	if credErr != nil {
		return errors.New("Failed to retrieve profile, falling back to User input")
	}
	return nil
}
