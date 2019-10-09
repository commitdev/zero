package proto

import (
	"bytes"
	"fmt"

	"github.com/commitdev/commit0/config"
	"github.com/commitdev/commit0/templator"
	"github.com/commitdev/commit0/util"
	"log"
	"os"
	"os/exec"
)

func Generate(templator *templator.Templator, config *config.Commit0Config) {
	GenerateIDLMakefile(templator, config)
	GenerateProtoHealth(templator, config)
	GenerateServiceProtobufFiles(templator, config)
	GenerateProtoServiceLibs(config)
	GenerateGoModIDL(templator, config)
}

func GenerateGoModIDL(templator *templator.Templator, config *config.Commit0Config) {
	idlPath := fmt.Sprintf("%s-idl", config.Name)
	idlOutput := fmt.Sprintf("%s/go.mod", idlPath)
	err := util.CreateDirIfDoesNotExist(idlPath)
	f, err := os.Create(idlOutput)

	if err != nil {
		log.Printf("Error: %v", err)
	}

	templator.Go.GoModIDL.Execute(f, config)
}

func GenerateIDLMakefile(templator *templator.Templator, config *config.Commit0Config) {
	makeFilePath := fmt.Sprintf("%s-idl", config.Name)
	makeFileOutput := fmt.Sprintf("%s/Makefile", makeFilePath)

	err := util.CreateDirIfDoesNotExist(makeFilePath)
	if err != nil {
		log.Printf("Error generating prototool config: %v", err)
	}

	f, err := os.Create(makeFileOutput)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	templator.MakefileTemplate.Execute(f, config)
}

func GenerateProtoHealth(templator *templator.Templator, config *config.Commit0Config) {
	protoHealthPath := fmt.Sprintf("%s-idl/proto/health", config.Name)
	protoHealthOutput := fmt.Sprintf("%s/health.proto", protoHealthPath)

	err := util.CreateDirIfDoesNotExist(protoHealthPath)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	f, err := os.Create(protoHealthOutput)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	templator.ProtoHealthTemplate.Execute(f, config)
}

func GenerateServiceProtobufFiles(templator *templator.Templator, cfg *config.Commit0Config) {
	protoPath := fmt.Sprintf("%s-idl/proto", cfg.Name)
	for _, s := range cfg.Services {
		serviceProtoDir := fmt.Sprintf("%s/%s", protoPath, s.Name)
		err := os.Mkdir(serviceProtoDir, os.ModePerm)
		if os.IsExist(err) {
			log.Printf("%s service proto exists skipping.", s.Name)
			continue
		}

		serviceProtoFilePath := fmt.Sprintf("%s/%s.proto", serviceProtoDir, s.Name)

		f, err := os.Create(serviceProtoFilePath)

		data := struct {
			*config.Commit0Config
			ServiceName string
		}{
			cfg,
			s.Name,
		}

		templator.ProtoServiceTemplate.Execute(f, data)
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
