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

type Maintainer struct {
	Name  string
	Email string
}

type Service struct {
	Name        string
	Description string
	Language    string
	GitRepo     string `json:"gitRepo"`
}

type ProjectConfiguration struct {
	ProjectName       string `json:"projectName"`
	FrontendFramework string `json:"frontendFramework"`
	Organization      string
	Description       string
	Maintainers       []Maintainer
	Services          []Service
	Infrastructure    Infrastructure
}

type Infrastructure struct {
	AWS *AWS
}
type AWS struct {
	AccountID string
	Region    string
}
