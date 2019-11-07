package util

const (
	Go         = "go"
	React      = "react"
	Kubernetes = "kubernetes"
)

const ApiGeneratedYamlName = "generated-project.yml"

var supportedLanguages = [...]string{Go, React, Kubernetes}

func ValidateLanguage(language string) bool {
	for _, l := range supportedLanguages {
		if l == language {
			return true
		}
	}

	return false
}
