package util

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"text/template"
)

func CreateDirIfDoesNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		return err
	}
	return nil
}

var FuncMap = template.FuncMap{
	"Title": strings.Title,
}

func createTemplatedFile(fullFilePath string, template *template.Template, data interface{}) {
	f, err := os.Create(fullFilePath)
	if err != nil {
		log.Printf("Error creating file: %v", err)
	}
	err = template.Execute(f, data)
	if err != nil {
		log.Printf("Error templating: %v", err)
	}
}

func TemplateFileAndOverwrite(fileDir string, fileName string, template *template.Template, data interface{}) {
	fullFilePath := fmt.Sprintf("%v/%v", fileDir, fileName)
	err := os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		log.Printf("Error creating directory %v: %v", fullFilePath, err)
	}
	createTemplatedFile(fullFilePath, template, data)

}

func TemplateFileIfDoesNotExist(fileDir string, fileName string, template *template.Template, data interface{}) {
	fullFilePath := path.Join(fileDir, fileName)

	if _, err := os.Stat(fullFilePath); os.IsNotExist(err) {
		if fileDir != "" {
			err := CreateDirIfDoesNotExist(fileDir)
			if err != nil {
				log.Printf("Error creating directory %v: %v", fullFilePath, err)
			}
		}
		createTemplatedFile(fullFilePath, template, data)
	} else {
		log.Printf("%v already exists. skipping.", fullFilePath)
	}
}
