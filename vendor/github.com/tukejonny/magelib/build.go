package magelib

import (
	"path/filepath"

	"github.com/magefile/mage/sh"
)

type BuildPair struct {
	Os   string
	Arch string
}

func NewBuildPair(os string, arch string) *BuildPair {
	return &BuildPair{
		Os:   os,
		Arch: arch,
	}
}

// Build ... Build binary for project
func Build(binaryName string, kinds []*BuildPair) error {
	buildBin := filepath.Join(ProjectDir, binaryName)
	for _, pair := range kinds {
		env := map[string]string{
			"GOOS":   pair.Os,
			"GOARCH": pair.Arch,
		}

		if err := sh.RunWith(env, "go", "build", "-o", buildBin, "."); err != nil {
			return err
		}
	}
	return nil
}
