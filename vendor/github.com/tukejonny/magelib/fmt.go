package magelib

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	zglob "github.com/mattn/go-zglob"
	"github.com/pkg/errors"
	"golang.org/x/tools/imports"
)

func ignoreWithPatterns(path string, noFormatPatterns []string) error {
	for _, pattern := range noFormatPatterns {
		matched, err := zglob.Match(pattern, path)
		if err != nil {
			return err
		} else if matched {
			return nil
		}
	}
	return nil
}

// Fmt ... Format codes in project.
func Fmt(noFormatPatterns []string) error {
	Info("Format code in project.")

	return filepath.Walk(ProjectDir, func(path string, info os.FileInfo, err error) error {
		// Ignore dirs & Skipping vendor/**
		if info.IsDir() {
			if path == filepath.Join(ProjectDir, "vendor") {
				return filepath.SkipDir
			}
			return nil
		}

		// Test only go file
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		// Ignore with user patterns
		ignoreWithPatterns(path, noFormatPatterns)

		// goimports
		in, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		out, err := imports.Process(path, in, nil)
		if err != nil {
			return errors.Wrapf(err, "while goimports %s, error occured.", path)
		}

		// There are no changes
		if bytes.Equal(in, out) {
			return nil
		}

		return ioutil.WriteFile(path, out, 0644)
	})
}
