package golang

import (
	"fmt"
	"github.com/commitdev/sprout/util"

	"github.com/commitdev/sprout/config"
	"github.com/commitdev/sprout/templator"
	"log"
	"os"
	"os/exec"
)

func Generate(templator *templator.Templator, config *config.SproutConfig, outPath string) {
	GenerateProtoToolConfig(templator, config, outPath)
	GenerateProtoHealth(templator, config, outPath)
	GenerateProtoServices(templator, config, outPath)
	GenerateProtoServiceLibs(config, outPath)
}

func GenerateProtoToolConfig(templator *templator.Templator, config *config.SproutConfig, outPath string) {
	protoPath := fmt.Sprintf("%s/%s/idl/proto", outPath, config.Name)
	protoToolOutput := fmt.Sprintf("%s/prototool.yaml", protoPath)

	err := util.CreateDirIfDoesNotExist(protoPath)
	if err != nil {
		log.Printf("Error generating prototool config: %v", err)
	}

	f, err := os.Create(protoToolOutput)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	templator.ProtoToolTemplate.Execute(f, config)
}

func GenerateProtoHealth(templator *templator.Templator, config *config.SproutConfig, outPath string) {
	protoHealthPath := fmt.Sprintf("%s/%s/idl/proto/health", outPath, config.Name)
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

func GenerateProtoServices(templator *templator.Templator, config *config.SproutConfig, outPath string) {
	protoToolConfigPath := fmt.Sprintf("%s/%s/idl/proto", outPath, config.Name)
	for _, s := range config.Services {
		idlPath := fmt.Sprintf("%s/%s", protoToolConfigPath, s.Name)
		err := util.CreateDirIfDoesNotExist(idlPath)
		if err != nil {
			log.Printf("Error generating service proto: %v", err)
		}

		//local paths
		protoPath := fmt.Sprintf("%s/%s.proto", s.Name, s.Name)
		cmd := exec.Command("prototool", "create", protoPath)
		cmd.Dir = protoToolConfigPath
		cmd.Run()
	}

}

func GenerateProtoServiceLibs(config *config.SproutConfig, outPath string) {
	protoToolConfigPath := fmt.Sprintf("%s/%s/idl/proto", outPath, config.Name)
	cmd := exec.Command("prototool", "generate")
	cmd.Dir = protoToolConfigPath
	err := cmd.Run()
	if err != nil {
		log.Printf("Error executing prototool generate: %v", err)
	}
}
