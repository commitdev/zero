package templator

import "github.com/commitdev/commit0/internal/config"

// GolangTemplateData holds data for use in golang related templates
type GolangTemplateData struct {
	Config  config.Commit0Config
	Service config.Service
}

// CITemplateData holds data for use in CI related templates
type CITemplateData struct {
	Config config.Commit0Config
	CI     config.CI
}
