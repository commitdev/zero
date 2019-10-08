package util

import (
	"os"
	"strings"
	"fmt"
	"log"
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

func TemplateFileIfDoesNotExist(fileDir string, fileName string, template *template.Template, data interface{}) {
	fullFilePath := fmt.Sprintf("%v/%v", fileDir, fileName)

	if _, err := os.Stat(fullFilePath); os.IsNotExist(err) {
		err := CreateDirIfDoesNotExist(fileDir)
		f, err := os.Create(fullFilePath)
		if err != nil {
			log.Printf("Error creating file: %v", err)
		}
		err = template.Execute(f, data)
		if err != nil {
			log.Printf("Error templating: %v", err)
		}
	} else {
		log.Printf("%v already exists. skipping.", fullFilePath)
	}
}
