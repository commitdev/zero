package proto

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"

	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, service config.Service, wg *sync.WaitGroup, pathPrefix string) {
	idlName := fmt.Sprintf("%s-idl", cfg.Name)
	idlPath := path.Join(pathPrefix, idlName)
	idlHealthPath := filepath.Join(idlPath, "proto", "health")

	data := templator.GolangTemplateData{
		*cfg,
		service,
	}

	util.TemplateFileIfDoesNotExist(idlPath, "go.mod", t.Go.GoModIDL, wg, data)
	util.TemplateFileIfDoesNotExist(idlPath, "Makefile", t.MakefileTemplate, wg, data)
	util.TemplateFileIfDoesNotExist(idlHealthPath, "health.proto", t.ProtoHealthTemplate, wg, data)

	serviceProtoDir := filepath.Join(idlPath, "proto", service.Name)
	file := fmt.Sprintf("%s.proto", service.Name)
	util.TemplateFileIfDoesNotExist(serviceProtoDir, file, t.ProtoServiceTemplate, wg, data)

	GenerateProtoServiceLibs(idlPath)
}

func GenerateProtoServiceLibs(idlPath string) {
	cmd := exec.Command("make", "generate")
	cmd.Dir = idlPath
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	log.Print("Generating proto service libs...")
	if err != nil {
		log.Println(aurora.Red(emoji.Sprintf(":exclamation: Failed running command in: %v", cmd.Dir)))
		log.Println(aurora.Red(emoji.Sprintf(":exclamation: Error executing protoc generation: %v %v", err, stderr.String())))
	}
}
