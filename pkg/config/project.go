package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func LocateProjectRootDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	var accessed = make(map[string]any)

	// Traverse upwards until either we find the project root dir or we access a directory again
	// That's a safeguard in case we reach the filesystem root directory.

	for {
		if _, ok := accessed[dir]; ok {
			return "", fmt.Errorf("no go.mod file found")
		}

		accessed[dir] = struct{}{}

		if isProjectRootDir(dir) {
			return dir, nil
		}

		dir = filepath.Dir(dir)
	}
}

func isProjectRootDir(dir string) bool {
	goModPath := filepath.Join(dir, "go.mod")
	_, err := os.Stat(goModPath)
	return err == nil
}
