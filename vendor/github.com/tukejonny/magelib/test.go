package magelib

import (
	"bufio"
	"bytes"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
	pipeline "github.com/mattn/go-pipeline"
	zglob "github.com/mattn/go-zglob"
)

func listTestPackages(noTestPatterns []string) (testPackages []string, err error) {
	// Buffering `go list` results.
	out, err := pipeline.Output(
		[]string{"go", "list", filepath.Join(ProjectPkg, "...")},
		[]string{"grep", "-v", "vendor"},
	)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(bytes.NewReader(out))
	for scanner.Scan() {
		// Readline
		pkg := scanner.Text()

		// Ignore vendor directory by defaults
		if strings.Contains(pkg, "vendor") {
			continue
		}

		// Check pkg in noTestPackages
		testFlag := true
		for _, pattern := range noTestPatterns {
			matched, err := zglob.Match(pattern, pkg)
			if err != nil {
				return nil, err
			} else if matched {
				testFlag = false
				break
			}
		}

		// Do test
		if testFlag {
			testPackages = append(testPackages, pkg)
		}
	}

	return
}

// Test ... Execute tests
func Test(noTestPackages []string) error {
	pkgs, err := listTestPackages(noTestPackages)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		if err := sh.RunV("go", "test", "-v", pkg); err != nil {
			return err
		}
	}

	return nil
}
