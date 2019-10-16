package config

type reactApp struct {
	Name string
}

type reactHeader struct {
	Enabled bool
}

type reactSidenav struct {
	Enabled bool
}

type reactAccount struct {
	Enabled  bool
	Required bool
}

type reactView struct {
	Path      string
	Component string
}

type React struct {
	App     reactApp
	Account reactAccount
	Header  reactHeader
	Sidenav reactSidenav
	Views   []reactView
}
