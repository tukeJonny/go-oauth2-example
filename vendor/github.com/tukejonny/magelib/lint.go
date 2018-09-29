package magelib

import (
	"path/filepath"

	"github.com/magefile/mage/sh"
)

// Lint ... Lint codes in project.
func Lint() error {
	return sh.RunV("gometalinter", "--config=.gometalinter.json", filepath.Join(ProjectDir, "."))
}
