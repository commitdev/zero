package constants

const (
	TmpRegistryYml    = "tmp/registry.yaml"
	TemplatesDir      = "tmp/templates"
	ZeroProjectYml    = "zero-project.yml"
	ZeroModuleYml     = "zero-module.yml"
	ZeroHomeDirectory = ".zero"
	IgnoredPaths      = "(?i)zero.module.yml|.git/"
	TemplateExtn      = ".tmpl"

	// prompt constants

	MaxPnameLength     = 16
	MaxOnameLength     = 39
	RegexValidation    = "regex"
	FunctionValidation = "function"
	ZeroReleaseURL     = "https://github.com/commitdev/zero/releases"
)
