package magelib

import (
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
)

// Project settings
var (
	ProjectPkg = func() string {
		result, err := exec.Command("go", "list", ".").Output()
		if err != nil {
			panic(err)
		}
		return strings.TrimSpace(string(result))
	}()
	ProjectDir = func() string {
		result, err := os.Getwd()
		if err != nil {
			panic(errors.Wrap(err, "We can't specify cwd!"))
		}
		return result
	}()
)
