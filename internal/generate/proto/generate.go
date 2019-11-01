package proto

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func Generate(t *templator.Templator, cfg *config.Commit0Config, service config.Service, wg *sync.WaitGroup) {
	idlPath := fmt.Sprintf("%s-idl", service.Name)
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

	GenerateProtoServiceLibs(service)
}

func GenerateProtoServiceLibs(service config.Service) {
	idlRoot := fmt.Sprintf("%s-idl", service.Name)
	cmd := exec.Command("make", "generate")
	cmd.Dir = idlRoot
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()

	log.Print("Generating proto service libs...")
	if err != nil {
		log.Printf("Failed running command in: %v", cmd.Dir)
		log.Printf("Error executing protoc generation: %v %v", err, stderr.String())
	}
}
