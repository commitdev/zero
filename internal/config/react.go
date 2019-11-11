package config

type reactApp struct {
	Name string
}

type reactAccount struct {
	Enabled  bool
	Required bool
}

type reactView struct {
	Path      string
	Component string
}

type frontend struct {
	Framework string
	Hostname  string
	App       reactApp
	CI        CI
}
