package templator

import (
	"path/filepath"
	"strings"
	"sync"
	"text/template"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/util"
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file"
)

// DockerTemplator contains the templates relevent to docker
type DockerTemplator struct {
	ApplicationDocker *template.Template
	HttpGatewayDocker *template.Template
	DockerIgnore      *template.Template
	DockerCompose     *template.Template
}

// GoTemplator contains the templates relevant to a go project
type GoTemplator struct {
	GoMain         *template.Template
	GoMod          *template.Template
	GoModIDL       *template.Template
	GoServer       *template.Template
	GoHealthServer *template.Template
	GoHTTPGW       *template.Template
}

// Templator contains all the templates
type Templator struct {
	Commit0              *template.Template
	GitIgnore            *template.Template
	Readme               *template.Template
	MakefileTemplate     *template.Template
	ProtoHealthTemplate  *template.Template
	ProtoServiceTemplate *template.Template
	Go                   *GoTemplator
	Docker               *DockerTemplator
	React                *DirectoryTemplator
}

func NewTemplator(box *packr.Box) *Templator {
	return &Templator{
		MakefileTemplate:     NewSingleFileTemplator(box, "proto/makefile.tmpl"),
		ProtoHealthTemplate:  NewSingleFileTemplator(box, "proto/health_proto.tmpl"),
		ProtoServiceTemplate: NewSingleFileTemplator(box, "proto/service_proto.tmpl"),
		Go:                   NewGoTemplator(box),
		Commit0:              NewSingleFileTemplator(box, "commit0/commit0.tmpl"),
		GitIgnore:            NewSingleFileTemplator(box, "util/gitignore.tmpl"),
		Readme:               NewSingleFileTemplator(box, "util/README.tmpl"),
		Docker:               NewDockerFileTemplator(box),
		React:                NewDirectoryTemplator(box, "react"),
	}
}

func NewGoTemplator(box *packr.Box) *GoTemplator {
	return &GoTemplator{
		GoMain:         NewSingleFileTemplator(box, "golang/main.tmpl"),
		GoMod:          NewSingleFileTemplator(box, "golang/go_mod.tmpl"),
		GoModIDL:       NewSingleFileTemplator(box, "golang/go_mod_idl.tmpl"),
		GoServer:       NewSingleFileTemplator(box, "golang/server.tmpl"),
		GoHealthServer: NewSingleFileTemplator(box, "golang/health_server.tmpl"),
		GoHTTPGW:       NewSingleFileTemplator(box, "golang/http_gw.tmpl"),
	}

}

func NewDockerFileTemplator(box *packr.Box) *DockerTemplator {
	return &DockerTemplator{
		ApplicationDocker: NewSingleFileTemplator(box, "docker/dockerfile_app.tmpl"),
		HttpGatewayDocker: NewSingleFileTemplator(box, "docker/dockerfile_http.tmpl"),
		DockerIgnore:      NewSingleFileTemplator(box, "docker/dockerignore.tmpl"),
		DockerCompose:     NewSingleFileTemplator(box, "docker/dockercompose.tmpl"),
	}
}

// NewSingleFileTemplator returns a template struct for a given template file
func NewSingleFileTemplator(box *packr.Box, file string) *template.Template {
	source, err := box.FindString(file)
	if err != nil {
		panic(err)
	}

	t, err := template.New(file).Funcs(util.FuncMap).Parse(source)
	if err != nil {
		panic(err)
	}

	return t
}

type DirectoryTemplator struct {
	Templates []*template.Template
}

func (d *DirectoryTemplator) TemplateFiles(config *config.Commit0Config, overwrite bool, wg *sync.WaitGroup) {
	for _, template := range d.Templates {
		d, f := filepath.Split(template.Name())
		if strings.HasSuffix(f, ".tmpl") {
			f = strings.Replace(f, ".tmpl", "", -1)
		}
		if overwrite {
			util.TemplateFileAndOverwrite(d, f, template, wg, config)
		} else {
			util.TemplateFileIfDoesNotExist(d, f, template, wg, config)
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
