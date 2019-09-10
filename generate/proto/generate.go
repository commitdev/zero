package proto

import (
	"fmt"
	"github.com/commitdev/sprout/util"

	"github.com/commitdev/sprout/config"
	"github.com/commitdev/sprout/templator"
	"log"
	"os"
	"os/exec"
)

func Generate(templator *templator.Templator, config *config.SproutConfig) {
	GenerateProtoToolConfig(templator, config)
	GenerateProtoHealth(templator, config)
	GenerateProtoServices(templator, config)
	GenerateProtoServiceLibs(config)
	GenerateGoModIDL(templator, config)
}

func GenerateGoModIDL(templator *templator.Templator, config *config.SproutConfig) {
	idlPath := fmt.Sprintf("../%s-idl", config.Name)
	idlOutput := fmt.Sprintf("%s/go.mod", idlPath)

	f, err := os.Create(idlOutput)

	err = util.CreateDirIfDoesNotExist(idlPath)
	if err != nil {
		log.Printf("Error: %v", err)
	}

	templator.Go.GoModIDL.Execute(f, config)
}

func GenerateProtoToolConfig(templator *templator.Templator, config *config.SproutConfig) {
	protoPath := fmt.Sprintf("../%s-idl/proto", config.Name)
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

func GenerateProtoHealth(templator *templator.Templator, config *config.SproutConfig) {
	protoHealthPath := fmt.Sprintf("../%s-idl/proto/health", config.Name)
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

func GenerateProtoServices(templator *templator.Templator, config *config.SproutConfig) {
	protoToolConfigPath := fmt.Sprintf("../%s-idl/proto", config.Name)
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
		log.Printf("Generating %v", protoPath)
		cmd.Run()
	}

}

func GenerateProtoServiceLibs(config *config.SproutConfig) {
	protoToolConfigPath := fmt.Sprintf("../%s-idl/proto", config.Name)
	cmd := exec.Command("prototool", "generate")
	cmd.Dir = protoToolConfigPath
	bytes, err := cmd.Output()
	
	log.Print("Generating proto service libs...")
	if err != nil {
		log.Printf("Error executing prototool generate: %v", string(bytes))
	}
}
