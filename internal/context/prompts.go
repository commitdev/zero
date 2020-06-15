package context

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/commitdev/zero/internal/config/globalconfig"
	"github.com/commitdev/zero/internal/config/moduleconfig"
	"github.com/commitdev/zero/pkg/util/exit"
	"github.com/manifoldco/promptui"
	"gopkg.in/yaml.v2"
)

type PromptHandler struct {
	moduleconfig.Parameter
	Condition func(map[string]string) bool
	Validate  func(string) error
}

func NoCondition(map[string]string) bool {
	return true
}

func KeyMatchCondition(key string, value string) func(map[string]string) bool {
	return func(param map[string]string) bool {
		return param[key] == value
	}
}

func NoValidation(string) error {
	return nil
}

func SpecificValueValidation(values ...string) func(string) error {
	return func(checkValue string) error {
		for _, allowedValue := range values {
			if checkValue == allowedValue {
				return nil
			}
		}
		return fmt.Errorf("Please choose one of %s", strings.Join(values, "/"))
	}
}

// TODO: validation / allow prompt retry ...etc
func (p PromptHandler) GetParam(projectParams map[string]string) string {
	var err error
	var result string
	if p.Condition(projectParams) {
		// TODO: figure out scope of projectParams per project
		// potentially dangerous to have cross module env leaking
		// so if community module has an `execute: twitter tweet $ENV`
		// it wouldnt leak things the module shouldnt have access to
		if p.Parameter.Execute != "" {
			result = executeCmd(p.Parameter.Execute, projectParams)
		} else if p.Parameter.Value != "" {
			result = p.Parameter.Value
		} else {
			err, result = promptParameter(p)
		}
		if err != nil {
			exit.Fatal("Exiting prompt:  %v\n", err)
		}

		return sanitizeParameterValue(result)
	}
	return ""
}

func promptParameter(prompt PromptHandler) (error, string) {
	param := prompt.Parameter
	label := param.Label
	if param.Label == "" {
		label = param.Field
	}
	defaultValue := param.Default

	var err error
	var result string
	if len(param.Options) > 0 {
		prompt := promptui.Select{
			Label: label,
			Items: param.Options,
		}
		_, result, err = prompt.Run()

	} else {
		prompt := promptui.Prompt{
			Label:     label,
			Default:   defaultValue,
			AllowEdit: true,
			Validate:  prompt.Validate,
		}
		result, err = prompt.Run()
	}
	if err != nil {
		exit.Fatal("Exiting prompt:  %v\n", err)
	}

	return nil, result
}

func executeCmd(command string, envVars map[string]string) string {
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = appendProjectEnvToCmdEnv(envVars, os.Environ())
	out, err := cmd.Output()

	if err != nil {
		log.Fatalf("Failed to execute  %v\n", err)
	}
	return string(out)
}

// aws cli prints output with linebreak in them
func sanitizeParameterValue(str string) string {
	re := regexp.MustCompile("\\n")
	return re.ReplaceAllString(str, "")
}

func appendProjectEnvToCmdEnv(envMap map[string]string, envList []string) []string {
	for key, val := range envMap {
		if val != "" {
			envList = append(envList, fmt.Sprintf("%s=%s", key, val))
		}
	}
	return envList
}

// PromptParams renders series of prompt UI based on the config
func PromptModuleParams(moduleConfig moduleconfig.ModuleConfig, parameters map[string]string) (map[string]string, error) {

	for _, promptConfig := range moduleConfig.Parameters {
		// deduplicate fields already prompted and received
		if _, isAlreadySet := parameters[promptConfig.Field]; isAlreadySet {
			continue
		}
		promptHandler := PromptHandler{
			promptConfig,
			NoCondition,
			NoValidation,
		}
		result := promptHandler.GetParam(parameters)

		parameters[promptConfig.Field] = result
	}
	return parameters, nil
}

func promptCredentialsAndFillProjectCreds(credentialPrompts map[string][]PromptHandler, credentials globalconfig.ProjectCredential) globalconfig.ProjectCredential {
	promptsValues := map[string]map[string]string{}

	for vendor, prompts := range credentialPrompts {
		vendorPromptValues := map[string]string{}

		// vendors like AWS have multiple prompts (accessKeyId and secretAccessKey)
		for _, prompt := range prompts {
			vendorPromptValues[prompt.Field] = prompt.GetParam(map[string]string{})
		}
		promptsValues[vendor] = vendorPromptValues
	}

	// FIXME: what is a good way to dynamically modify partial data of a struct
	// current just marashing to yaml, then unmarshaling into the base struct
	yamlContent, _ := yaml.Marshal(promptsValues)
	yaml.Unmarshal(yamlContent, &credentials)
	return credentials
}

func appendToSet(set []string, toAppend []string) []string {
	for _, appendee := range toAppend {
		if !itemInSlice(set, appendee) {
			set = append(set, appendee)
		}
	}
	return set
}

func itemInSlice(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
