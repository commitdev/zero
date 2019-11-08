package config

type reactApp struct {
	Name string
}

type reactHeader struct {
	Enabled bool
}

type reactSidenavItem struct {
	Path  string
	Label string
	Icon  string
}
type reactSidenav struct {
	Enabled bool
	Items   []reactSidenavItem
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
	App       reactApp
	Account   reactAccount
	Header    reactHeader
	Sidenav   reactSidenav
	Views     []reactView
	CI        CI
}
