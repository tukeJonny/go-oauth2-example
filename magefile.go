// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/tukejonny/magelib"
)

var (
	binaryName = "goauthex"
	buildKinds = []*magelib.BuildPair{
		magelib.NewBuildPair("darwin", "386"),
	}
	noFormatPatterns = []string{}
	noTestPatterns   = []string{}
)

// Execute all tasks
func All() error {
	magelib.Info("Running all tasks ...")
	mg.SerialDeps(Build)
	mg.SerialDeps(Fmt)
	mg.SerialDeps(Lint)
	mg.SerialDeps(Test)

	return nil
}

// Build binary
func Build() error {
	magelib.Info("Build project.")
	if err := magelib.Build(binaryName, buildKinds); err != nil {
		return err
	}
	return nil
}

// Prettify codes
func Fmt() error {
	magelib.Info("Format code in project.")

	if err := magelib.Fmt(noFormatPatterns); err != nil {
		return err
	}

	return nil
}

// Check codestyle
func Lint() error {
	magelib.Info("Check codingstyle of this project.")
	return magelib.Lint()
}

// Execute tests
func Test() error {
	magelib.Info("Execute tests for this project.")

	if err := magelib.Test(noTestPatterns); err != nil {
		return err
	}

	return nil
}

// Start godoc webserver
func Godoc() error {
	magelib.Info("Run godoc server.")
	return magelib.Godoc()
}
