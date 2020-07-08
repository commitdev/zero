package flog

import (
	"fmt"
	"log"

	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
)

// Warnf logs a formatted error message
func Infof(format string, a ...interface{}) {
	log.Println(aurora.Cyan(emoji.Sprintf(format, a...)))
}

// Infof prints out a timestamp as prefix, Guidef just prints the message
func Guidef(format string, a ...interface{}) {
	fmt.Println(aurora.Cyan(emoji.Sprintf(format, a...)))
}

// Successf logs a formatted success message
func Successf(format string, a ...interface{}) {
	log.Println(aurora.Green(emoji.Sprintf(":white_check_mark: "+format, a...)))
}

// Warnf logs a formatted warning message
func Warnf(format string, a ...interface{}) {
	log.Println(aurora.Yellow(emoji.Sprintf(":exclamation: "+format, a...)))
}

// Warnf logs a formatted error message
func Errorf(format string, a ...interface{}) {
	log.Println(aurora.Red(emoji.Sprintf(":exclamation: "+format, a...)))
}
