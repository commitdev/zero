package main

import (
	"github.com/commitdev/sprout/cmd"
	"github.com/commitdev/sprout/templator"
	"github.com/gobuffalo/packr/v2"
)

func main() {
	templates := packr.New("templates", "./templates")
	templator := templator.NewTemplator(templates)
	cmd.Execute(templator)
}
