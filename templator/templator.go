package templator

import (
	"github.com/gobuffalo/packr/v2"
	"text/template"
)

type Templator struct {
	ProtoToolTemplate   *template.Template
	ProtoHealthTemplate *template.Template
}

func NewTemplator(box *packr.Box) *Templator {
	protoToolTemplateSource, _ := box.FindString("prototool.tmpl")
	protoToolTemplate, _ := template.New("ProtoToolTemplate").Parse(protoToolTemplateSource)

	protoHealthTemplateSource, _ := box.FindString("health_proto.tmpl")
	protoHealthTemplate, _ := template.New("ProtoHealthTemplate").Parse(protoHealthTemplateSource)

	return &Templator{
		ProtoToolTemplate:   protoToolTemplate,
		ProtoHealthTemplate: protoHealthTemplate,
	}
}
