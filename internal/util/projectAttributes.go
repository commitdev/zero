package util

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
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Service struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"`
	GitRepo     string `json:"gitRepo"`
}

type ProjectConfiguration struct {
	ProjectName       string       `json:"projectName"`
	FrontendFramework string       `json:"frontendFramework"`
	Organization      string       `json:"organization"`
	Description       string       `json:"description"`
	Maintainers       []Maintainer `json:"maintainers"`
	Services          []Service    `json:"services"`
}
