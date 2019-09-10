package templator

import (
	"github.com/commitdev/sprout/util"
	"github.com/gobuffalo/packr/v2"
	"text/template"
)

type GoTemplator struct {
	GoMain         *template.Template
	GoMod          *template.Template
	GoModIDL       *template.Template
	GoServer       *template.Template
	GoHealthServer *template.Template
}

type Templator struct {
	Sprout              *template.Template
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
		Sprout:              NewSproutTemplator(box),
	}
}

func NewGoTemplator(box *packr.Box) *GoTemplator {
	goServerTemplateSource, _ := box.FindString("golang/server.tmpl")
	goServerTemplate, _ := template.New("GoServerTemplate").Funcs(util.FuncMap).Parse(goServerTemplateSource)

	goHealthTemplateSource, _ := box.FindString("golang/health_server.tmpl")
	goHealthServerTemplate, _ := template.New("GoHealthServerTemplate").Parse(goHealthTemplateSource)

	goModTemplateSource, _ := box.FindString("golang/go_mod.tmpl")
	goModTemplate, _ := template.New("GoModTemplate").Parse(goModTemplateSource)

	goModIDLTemplateSource, _ := box.FindString("golang/go_mod_idl.tmpl")
	goModIDLTemplate, _ := template.New("GoModTemplate").Parse(goModIDLTemplateSource)

	goMainTemplateSource, _ := box.FindString("golang/main.tmpl")
	goMainTemplate, _ := template.New("GoMainTemplate").Parse(goMainTemplateSource)

	return &GoTemplator{
		GoMain:         goMainTemplate,
		GoMod:          goModTemplate,
		GoModIDL:       goModIDLTemplate,
		GoServer:       goServerTemplate,
		GoHealthServer: goHealthServerTemplate,
	}

}

func NewSproutTemplator(box *packr.Box) *template.Template {
	templateSource, _ := box.FindString("sprout/sprout.tmpl")
	template, _ := template.New("SproutTemplate").Funcs(util.FuncMap).Parse(templateSource)

	return template
}
