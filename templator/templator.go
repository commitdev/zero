package templator

import (
	"github.com/commitdev/sprout/util"
	"github.com/gobuffalo/packr/v2"
	"text/template"
)

type DockerTemplator struct {
	ApplicationDocker *template.Template
	HttpGatewayDocker *template.Template
	DockerIgnore *template.Template
}

type GoTemplator struct {
	GoMain         *template.Template
	GoMod          *template.Template
	GoModIDL       *template.Template
	GoServer       *template.Template
	GoHealthServer *template.Template
	GoHttpGW       *template.Template
}

type Templator struct {
	Sprout               *template.Template
	GitIgnore            *template.Template
	MakefileTemplate     *template.Template
	ProtoHealthTemplate  *template.Template
	ProtoServiceTemplate *template.Template
	Go                   *GoTemplator
	Docker               *DockerTemplator
}

func NewTemplator(box *packr.Box) *Templator {
	makeFileTemplateSource, _ := box.FindString("proto/makefile.tmpl")
	makeFileTemplate, _ := template.New("ProtoToolTemplate").Parse(makeFileTemplateSource)

	protoHealthTemplateSource, _ := box.FindString("proto/health_proto.tmpl")
	protoHealthTemplate, _ := template.New("ProtoHealthTemplate").Parse(protoHealthTemplateSource)

	protoServiceTemplateSource, _ := box.FindString("proto/service_proto.tmpl")
	protoServiceTemplate, _ := template.New("ProtoServiceTemplate").Funcs(util.FuncMap).Parse(protoServiceTemplateSource)

	return &Templator{
		MakefileTemplate:     makeFileTemplate,
		ProtoHealthTemplate:  protoHealthTemplate,
		ProtoServiceTemplate: protoServiceTemplate,
		Go:                   NewGoTemplator(box),
		Sprout:               NewSproutTemplator(box),
		GitIgnore:            NewGitIgnoreTemplator(box),
		Docker:               NewDockerFileTemplator(box),
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
	goMainTemplate, _ := template.New("GoMainTemplate").Funcs(util.FuncMap).Parse(goMainTemplateSource)

	goHttpTemplateSource, _ := box.FindString("golang/http_gw.tmpl")
	goHttpTemplate, _ := template.New("GoHttpGWTemplate").Funcs(util.FuncMap).Parse(goHttpTemplateSource)

	return &GoTemplator{
		GoMain:         goMainTemplate,
		GoMod:          goModTemplate,
		GoModIDL:       goModIDLTemplate,
		GoServer:       goServerTemplate,
		GoHealthServer: goHealthServerTemplate,
		GoHttpGW:       goHttpTemplate,
	}

}

func NewSproutTemplator(box *packr.Box) *template.Template {
	templateSource, _ := box.FindString("sprout/sprout.tmpl")
	template, _ := template.New("SproutTemplate").Funcs(util.FuncMap).Parse(templateSource)

	return template
}

func NewGitIgnoreTemplator(box *packr.Box) *template.Template {
	templateSource, _ := box.FindString("util/gitignore.tmpl")
	template, _ := template.New("GitIgnore").Parse(templateSource)
	return template
}

func NewDockerFileTemplator(box *packr.Box) *DockerTemplator {
	appTemplateSource, _ := box.FindString("docker/dockerfile_app.tmpl")
	appTemplate, _ := template.New("AppDockerfile").Parse(appTemplateSource)

	httpTemplateSource, _ := box.FindString("docker/dockerfile_http.tmpl")
	httpTemplate, _ := template.New("HttpDockerfile").Parse(httpTemplateSource)

	ignoreTemplateSource, _ := box.FindString("docker/dockerignore.tmpl")
	ignoreTemplate, _ := template.New("Dockerignore").Parse(ignoreTemplateSource)

	return &DockerTemplator{
		ApplicationDocker: appTemplate,
		HttpGatewayDocker: httpTemplate,
		DockerIgnore: ignoreTemplate,
	}
}
