package flog

import (
	"fmt"
	"os"

	"github.com/kyokomi/emoji"
	"github.com/logrusorgru/aurora"
	"github.com/sirupsen/logrus"
)

const LogEnvVariable = "LOG_LEVEL"
const defaultLogLevel = "info"

var logger = getLogger()
var infoFormatter = new(InfoFormatter)
var debugFormatter = &logrus.TextFormatter{
	DisableLevelTruncation:    true,
	FullTimestamp:             true,
	EnvironmentOverrideColors: true,
}

func getLogger() *logrus.Logger {
	logger := logrus.New()

	lvl, ok := os.LookupEnv(LogEnvVariable)
	if !ok {
		lvl = defaultLogLevel
	}
	logLevel, _ := logrus.ParseLevel(lvl)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logLevel)
	return logger
}

// Warnf logs a formatted error message
func Infof(format string, a ...interface{}) {
	logger.SetFormatter(infoFormatter)
	logger.Info(aurora.Cyan(emoji.Sprintf(format, a...)))
}

func Debugf(format string, a ...interface{}) {
	logger.SetFormatter(debugFormatter)
	logger.Debug(aurora.Green(emoji.Sprintf(format, a...)))
}

// Infof prints out a timestamp as prefix, Guidef just prints the message
func Guidef(format string, a ...interface{}) {
	fmt.Println(aurora.Cyan(emoji.Sprintf(format, a...)))
}

// Successf logs a formatted success message
func Successf(format string, a ...interface{}) {
	logger.Info(aurora.Green(emoji.Sprintf(":white_check_mark: "+format, a...)))
}

// Warnf logs a formatted warning message
func Warnf(format string, a ...interface{}) {
	logger.Warn(aurora.Yellow(emoji.Sprintf(":exclamation: "+format, a...)))
}

// Warnf logs a formatted error message
func Errorf(format string, a ...interface{}) {
	logger.Error(aurora.Red(emoji.Sprintf(":exclamation: "+format, a...)))
}

// Info formatter is to not display the LOG_LEVEL in front of the command eg. INFO[2020-070-01T15:22:22] Hello World
type InfoFormatter struct {
}

func (f *InfoFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// extra line break stops the prompts from overtaking Existing line
	return []byte(entry.Message + "\n"), nil
}
