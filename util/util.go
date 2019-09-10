package util

import (
	"os"
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
