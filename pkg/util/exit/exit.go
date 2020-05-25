package exit

import (
	"os"

	"github.com/commitdev/zero/pkg/util/flog"
)

const (
	// CodeOK indicates successful execution.
	CodeOK = 0

	// CodeError indicates erroneous execution.
	CodeError = 1

	// CodeFatal indicates erroneous use by user.
	CodeFatal = 2
)

// Fatal terminates execution using fatal exit code.
func Fatal(format string, a ...interface{}) {
	flog.Errorf(format, a...)
	os.Exit(CodeFatal)
}

// Error terminates execution using unsuccessful execution exit code.
func Error(format string, a ...interface{}) {
	flog.Errorf(format, a...)
	os.Exit(CodeError)
}

// OK terminates execution successfully.
func OK(format string, a ...interface{}) {
	flog.Infof(format, a)
	os.Exit(CodeOK)
}
