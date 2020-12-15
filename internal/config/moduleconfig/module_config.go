package moduleconfig

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	yaml "gopkg.in/yaml.v2"

	"github.com/commitdev/zero/pkg/util/flog"
	"github.com/iancoleman/strcase"
)

type ModuleConfig struct {
	Name                string
	Description         string
	Author              string
	DependsOn           []string `yaml:"dependsOn,omitempty"`
	TemplateConfig      `yaml:"template"`
	RequiredCredentials []string `yaml:"requiredCredentials"`
	Parameters          []Parameter
	Conditions          []Condition `yaml:"conditions,omitempty"`
}

type Parameter struct {
	Field           string
	Label           string   `yaml:"label,omitempty"`
	Options         []string `yaml:"options,omitempty"`
	Execute         string   `yaml:"execute,omitempty"`
	Value           string   `yaml:"value,omitempty"`
	Default         string   `yaml:"default,omitempty"`
	Info            string   `yaml:"info,omitempty"`
	FieldValidation Validate `yaml:"fieldValidation,omitempty"`
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

// A "nice" wrapper around findMissing()
func (cfg ModuleConfig) collectMissing() []string {
	var missing []string
	findMissing(reflect.ValueOf(cfg), "", "", &missing)

	return missing
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
