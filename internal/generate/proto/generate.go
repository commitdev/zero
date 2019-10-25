package proto

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"sync"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/commitdev/commit0/internal/util"
)

func Generate(templator *templator.Templator, config *config.Commit0Config, wg *sync.WaitGroup) {
	idlPath := fmt.Sprintf("%s-idl", config.Name)
	idlHealthPath := fmt.Sprintf("%s/proto/health", idlPath)

	util.TemplateFileIfDoesNotExist(idlPath, "go.mod", templator.Go.GoModIDL, wg, config)
	util.TemplateFileIfDoesNotExist(idlPath, "Makefile", templator.MakefileTemplate, wg, config)
	GenerateServiceProtobufFiles(templator, config, wg)
	util.TemplateFileIfDoesNotExist(idlHealthPath, "health.proto", templator.ProtoHealthTemplate, wg, config)

	GenerateProtoServiceLibs(config)
}

func GenerateServiceProtobufFiles(templator *templator.Templator, cfg *config.Commit0Config, wg *sync.WaitGroup) {
	protoPath := fmt.Sprintf("%s-idl/proto", cfg.Name)
	for _, s := range cfg.Services {
		serviceProtoDir := fmt.Sprintf("%s/%s", protoPath, s.Name)
		file := fmt.Sprintf("%s.proto", s.Name)

		data := struct {
			*config.Commit0Config
			ServiceName string
		}{
			cfg,
			s.Name,
		}

		util.TemplateFileIfDoesNotExist(serviceProtoDir, file, templator.ProtoServiceTemplate, wg, data)
	}
}

func GenerateProtoServiceLibs(config *config.Commit0Config) {
	idlRoot := fmt.Sprintf("%s-idl", config.Name)
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
