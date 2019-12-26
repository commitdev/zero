package util

// @TODO : Move this stuff from util into another package

const (
	Go         = "go"
	React      = "react"
	Kubernetes = "kubernetes"
)

const CommitYml = "commit0.yml"

var supportedLanguages = [...]string{Go, React, Kubernetes}

func ValidateLanguage(language string) bool {
	for _, l := range supportedLanguages {
		if l == language {
			return true
		}
	}

	return false
}

type ProjectConfiguration struct {
	ProjectName    string `json:"projectName"`
	Infrastructure Infrastructure
}

type Infrastructure struct {
	AWS *AWS
}
type AWS struct {
	AccountID string
	Region    string
}
