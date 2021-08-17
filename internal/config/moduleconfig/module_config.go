package moduleconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	goVersion "github.com/hashicorp/go-version"
	yaml "gopkg.in/yaml.v2"

	"github.com/commitdev/zero/internal/config/projectconfig"
	"github.com/commitdev/zero/internal/constants"
	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/commitdev/zero/version"
	"github.com/iancoleman/strcase"
)

type ModuleConfig struct {
	Name                string
	Description         string
	Author              string
	Commands            ModuleCommands `yaml:"commands,omitempty"`
	DependsOn           []string       `yaml:"dependsOn,omitempty"`
	TemplateConfig      `yaml:"template"`
	RequiredCredentials []string           `yaml:"requiredCredentials"`
	ZeroVersion         VersionConstraints `yaml:"zeroVersion,omitempty"`
	Parameters          []Parameter
	Conditions          []Condition `yaml:"conditions,omitempty"`
}

type ModuleCommands struct {
	Apply   string `yaml:"apply,omitempty"`
	Check   string `yaml:"check,omitempty"`
	Summary string `yaml:"summary,omitempty"`
}

func checkVersionAgainstConstrains(vc VersionConstraints, versionString string) bool {
	v, err := goVersion.NewVersion(versionString)
	if err != nil {
		return false
	}

	return vc.Check(v)
}

// ValidateZeroVersion receives a module config, and returns whether the running zero's binary
// is compatible with the module
func ValidateZeroVersion(mc ModuleConfig) bool {
	if mc.ZeroVersion.String() == "" {
		return true
	}

	zeroVersion := version.AppVersion
	flog.Debugf("Checking Zero version (%s) against %s", zeroVersion, mc.ZeroVersion)

	// Unreleased versions or test runs, defaults to SNAPSHOT when not declared
	if zeroVersion == "SNAPSHOT" {
		return true
	}

	return checkVersionAgainstConstrains(mc.ZeroVersion, zeroVersion)
}

type Parameter struct {
	Field               string
	Label               string        `yaml:"label,omitempty"`
	Options             yaml.MapSlice `yaml:"options,omitempty"`
	Execute             string        `yaml:"execute,omitempty"`
	Value               string        `yaml:"value,omitempty"`
	Default             string        `yaml:"default,omitempty"`
	Info                string        `yaml:"info,omitempty"`
	FieldValidation     Validate      `yaml:"fieldValidation,omitempty"`
	Type                string        `yaml:"type,omitempty"`
	OmitFromProjectFile bool          `yaml:"omitFromProjectFile,omitempty"`
	Conditions          []Condition   `yaml:"conditions,omitempty"`
	EnvVarName          string        `yaml:"envVarName,omitempty"`
}

type Condition struct {
	Action     string   `yaml:"action"`
	MatchField string   `yaml:"matchField"`
	WhenValue  string   `yaml:"whenValue"`
	Data       []string `yaml:"data,omitempty"`
}

type Validate struct {
	Type         string `yaml:"type,omitempty"`
	Value        string `yaml:"value,omitempty"`
	ErrorMessage string `yaml:"errorMessage,omitempty"`
}

type TemplateConfig struct {
	StrictMode bool
	Delimiters []string
	InputDir   string `yaml:"inputDir"`
	OutputDir  string `yaml:"outputDir"`
}

type VersionConstraints struct {
	goVersion.Constraints
}

// A "nice" wrapper around findMissing()
func (cfg ModuleConfig) collectMissing() []string {
	var missing []string
	findMissing(reflect.ValueOf(cfg), "", "", &missing)

	return missing
}

// GetParamEnvVarTranslationMap returns a map for translating parameter's `Field` into env-var keys
// It loops through each parameter then adds to translation map if applicable
// for zero apply / zero init's prompt execute,
// this is useful for translating params like AWS credentials for running the AWS cli
func (cfg ModuleConfig) GetParamEnvVarTranslationMap() map[string]string {
	translationMap := make(map[string]string)
	for i := 0; i < len(cfg.Parameters); i++ {
		param := cfg.Parameters[i]
		if param.EnvVarName != "" {
			translationMap[param.Field] = param.EnvVarName
		}
	}
	return translationMap
}

func LoadModuleConfig(filePath string) (ModuleConfig, error) {
	config := ModuleConfig{}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	missing := config.collectMissing()
	if len(missing) > 0 {
		flog.Errorf("%v is missing information", filePath)

		for _, m := range missing {
			flog.Errorf("\t %v", m)
		}

		log.Fatal("")
	}

	if !ValidateZeroVersion(config) {
		constraint := config.ZeroVersion.Constraints.String()
		errTpl := `Module(%s) requires Zero to be version %s. Your current Zero version is: %s
Please update your Zero version to %s.
Please check %s for available releases.`
		return config, errors.New(fmt.Sprintf(errTpl, config.Name, constraint, version.AppVersion, constraint, constants.ZeroReleaseURL))
	}
	return config, nil
}

// Recurses through a datastructure to find any missing data.
// This assumes several things:
// 1. The structure matches that defined by ModuleConfig and its child datastructures.
// 2. YAML struct field metadata is sufficient to define whether an attribute is missing or not.
//    That is, "yaml:foo,omitempty" tells us this is not a required field because we can omit it.
// 3. Slices and arrays are assumed to be optional.
//
// As this function recurses through the datastructure, it builds up a string
// path representing each node's path within the datastructure.
// If the value of the current node is equal to the zero value for its datatype
// and its struct field does *not* have a "omitempty" value, then we assume it
// is missing and add it to the resultset.
func findMissing(obj reflect.Value, path, metadata string, missing *[]string) {
	t := obj.Type()
	switch t.Kind() {
	case reflect.String:
		if obj.String() == "" && !strings.Contains(metadata, "omitempty") {
			*missing = append(*missing, path)
		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < obj.Len(); i++ {
			prefix := fmt.Sprintf("%v[%v]", path, i)
			findMissing(obj.Index(i), prefix, metadata, missing)
		}

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			fieldType := t.Field(i)
			fieldTags, _ := fieldType.Tag.Lookup("yaml")
			fieldVal := obj.Field(i)

			tags := strings.Split(fieldTags, ",")

			hasOmitEmpty := false
			// We have all metadata yaml tags, now let's remove the "omitempty" tag if
			// it is present.
			// Then if we have only one tag remaining, this must be the expected yaml
			// identifer.
			// Otherwise the name of the yaml identifier should match the struct
			// attribute name.
			for i := len(tags) - 1; i >= 0; i-- {
				tag := tags[i]
				if tag == "omitempty" {
					hasOmitEmpty = true
					tags = append(tags[:i], tags[i+1:]...)
				}
			}

			yamlName := strcase.ToLowerCamel(fieldType.Name)
			if len(tags) == 1 && tags[0] != "" { // For some reason, empty tag lists are giving a count of 1.
				yamlName = tags[0]
			}

			prefix := yamlName
			if path != "" {
				prefix = fmt.Sprintf("%v.%v", path, yamlName)
			}

			zeroVal := reflect.Zero(fieldType.Type)
			if fieldVal == zeroVal && !hasOmitEmpty {
				*missing = append(*missing, prefix)
			}

			findMissing(fieldVal, prefix, fieldTags, missing)
		}
	}
}

// SummarizeParameters receives all parameters gathered from prompts during `Zero init`
// and based on module definition to construct the parameters for each module for zero-project.yml
// filters out parameters defined as OmitFromProjectFile: true
func SummarizeParameters(module ModuleConfig, allParams map[string]string) map[string]string {
	moduleParams := make(projectconfig.Parameters)
	// Loop through all the prompted values and find the ones relevant to this module
	for parameterKey, parameterValue := range allParams {
		for _, moduleParameter := range module.Parameters {
			if moduleParameter.Field == parameterKey {
				if moduleParameter.OmitFromProjectFile {
					flog.Debugf("Omitted %s from %s", parameterKey, module.Name)
				} else {
					moduleParams[parameterKey] = parameterValue
				}
			}
		}
	}
	return moduleParams
}

// SummarizeConditions based on conditions from zero-module.yml
// creates and returns slice of conditions for project config
func SummarizeConditions(module ModuleConfig) []projectconfig.Condition {
	moduleConditions := make([]projectconfig.Condition, len(module.Conditions))

	for i, condition := range module.Conditions {
		moduleConditions[i] = projectconfig.Condition{
			Action:     condition.Action,
			MatchField: condition.MatchField,
			WhenValue:  condition.WhenValue,
			Data:       condition.Data,
		}
	}
	return moduleConditions
}

// UnmarshalYAML Parses a version constraint string into go-version constraint during yaml parsing
func (semVer *VersionConstraints) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var versionString string
	err := unmarshal(&versionString)
	if err != nil {
		return err
	}
	if versionString != "" {
		constraints, constErr := goVersion.NewConstraint(versionString)
		// If an invalid constraint is declared in a module
		// instead of erroring out we just print a warning message
		if constErr != nil {
			flog.Warnf("Zero version constraint invalid format: %s", constErr)
		}

		*semVer = VersionConstraints{constraints}
	}
	return nil
}
