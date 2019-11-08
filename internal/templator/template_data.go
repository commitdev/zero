package templator

import "github.com/commitdev/commit0/internal/config"

// GenericTemplateData holds data for use in any template, it just contains the config struct
type GenericTemplateData struct {
	Config config.Commit0Config
}

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
