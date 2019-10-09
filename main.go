package main

import (
	"github.com/commitdev/commit0/cmd"
	"github.com/commitdev/commit0/templator"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	templates := packr.New("templates", "./templates")
	templator := templator.NewTemplator(templates)
	cmd.Execute(templator)
}
