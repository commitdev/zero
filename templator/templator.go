package templator

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/commitdev/sprout/util"
	"text/template"
)

type GoTemplator struct {
	GoMod *template.Template
	GoServer *template.Template
	GoHealthServer *template.Template

}

type Templator struct {
	Sprout *template.Template
	ProtoToolTemplate   *template.Template
	ProtoHealthTemplate *template.Template
	Go                  *GoTemplator
}

func NewTemplator(box *packr.Box) *Templator {
	protoToolTemplateSource, _ := box.FindString("proto/prototool.tmpl")
	protoToolTemplate, _ := template.New("ProtoToolTemplate").Parse(protoToolTemplateSource)

	protoHealthTemplateSource, _ := box.FindString("proto/health_proto.tmpl")
	protoHealthTemplate, _ := template.New("ProtoHealthTemplate").Parse(protoHealthTemplateSource)

	return &Templator{
		ProtoToolTemplate:   protoToolTemplate,
		ProtoHealthTemplate: protoHealthTemplate,
		Go:                  NewGoTemplator(box),
		Sprout: NewSproutTemplator(box),
	}
}

func NewGoTemplator(box *packr.Box) *GoTemplator {
	goServerTemplateSource, _ := box.FindString("golang/server.tmpl")
	goServerTemplate, _ := template.New("GoServerTemplate").Funcs(util.FuncMap).Parse(goServerTemplateSource)

	goHealthTemplateSource, _ := box.FindString("golang/health_server.tmpl")
	goHealthServerTemplate, _ := template.New("GoHealthServerTemplate").Parse(goHealthTemplateSource)

	goModTemplateSource, _ := box.FindString("golang/go_mod.tmpl")
	goModTemplate, _ := template.New("GoModTemplate").Parse(goModTemplateSource)

	return &GoTemplator{
		GoMod: goModTemplate,
		GoServer: goServerTemplate,
		GoHealthServer: goHealthServerTemplate,
	}

}

func NewSproutTemplator(box *packr.Box) *template.Template {
	templateSource, _ := box.FindString("sprout/sprout.tmpl")
	template, _ := template.New("SproutTemplate").Funcs(util.FuncMap).Parse(templateSource)

	return template
}

