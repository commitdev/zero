package templator

import (
	"path/filepath"
	"strings"
	"text/template"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/util"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file"
)

type DockerTemplator struct {
	ApplicationDocker *template.Template
	HttpGatewayDocker *template.Template
	DockerIgnore      *template.Template
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
	Commit0              *template.Template
	GitIgnore            *template.Template
	MakefileTemplate     *template.Template
	ProtoHealthTemplate  *template.Template
	ProtoServiceTemplate *template.Template
	Go                   *GoTemplator
	Docker               *DockerTemplator
	React                *DirectoryTemplator
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
		Commit0:              NewCommit0Templator(box),
		GitIgnore:            NewGitIgnoreTemplator(box),
		Docker:               NewDockerFileTemplator(box),
		React:                NewDirectoryTemplator(box, "react"),
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

func NewCommit0Templator(box *packr.Box) *template.Template {
	templateSource, _ := box.FindString("commit0/commit0.tmpl")
	template, _ := template.New("Commit0Template").Funcs(util.FuncMap).Parse(templateSource)

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
		DockerIgnore:      ignoreTemplate,
	}
}

type DirectoryTemplator struct {
	Templates []*template.Template
}

func (d *DirectoryTemplator) TemplateFiles(config *config.Commit0Config, overwrite bool) {
	for _, template := range d.Templates {
		d, f := filepath.Split(template.Name())
		if strings.HasSuffix(f, ".tmpl") {
			f = strings.Replace(f, ".tmpl", "", -1)
		}
		if overwrite {
			util.TemplateFileAndOverwrite(d, f, template, config)
		} else {
			util.TemplateFileIfDoesNotExist(d, f, template, config)
		}
	}
}

func NewDirectoryTemplator(box *packr.Box, dir string) *DirectoryTemplator {
	templates := []*template.Template{}
	for _, file := range getFileNames(box, dir) {
		templateSource, _ := box.FindString(file)
		template, err := template.New(file).Funcs(util.FuncMap).Parse(templateSource)
		if err != nil {
			panic(err)
		}
		templates = append(templates, template)
	}
	return &DirectoryTemplator{
		Templates: templates,
	}
}

func getFileNames(box *packr.Box, dir string) []string {
	keys := []string{}
	box.WalkPrefix(dir, func(path string, info file.File) error {
		if info == nil {
			return nil
		}
		finfo, _ := info.FileInfo()
		if !finfo.IsDir() {
			keys = append(keys, path)
		}
		return nil
	})
	return removeTmplDuplicates(keys)
}

func removeTmplDuplicates(keys []string) []string {
	filteredKeys := []string{}
	for _, key := range keys {
		if !containsStr(keys, key+".tmpl") {
			filteredKeys = append(filteredKeys, key)
		}
	}
	return filteredKeys
}

func containsStr(arr []string, key string) bool {
	for _, val := range arr {
		if val == key {
			return true
		}
	}
	return false
}
