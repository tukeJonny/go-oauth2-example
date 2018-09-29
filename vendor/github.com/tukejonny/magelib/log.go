package magelib

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

// Color settings
var (
	infoColor = chalk.Cyan
	warnColor = chalk.Red
)

// Info ... Print informations to stderr
func Info(msg ...interface{}) {
	fmt.Fprint(os.Stderr, infoColor)
	fmt.Fprint(os.Stderr, msg...)
	fmt.Fprint(os.Stderr, chalk.Reset)
	fmt.Fprintln(os.Stderr, "")
}

// Warn ... Print warnings to stderr
func Warn(msg ...interface{}) {
	fmt.Fprint(os.Stderr, warnColor)
	fmt.Fprint(os.Stderr, msg...)
	fmt.Fprint(os.Stderr, chalk.Reset)
	fmt.Fprintln(os.Stderr, "")
}
