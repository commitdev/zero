package config

type reactApp struct {
	Name string
}

type environment struct {
	CognitoPoolID   string
	CognitoClientID string
}

type frontend struct {
	Framework string
	Hostname  string
	App       reactApp
	CI        CI
	Env       environment
}
