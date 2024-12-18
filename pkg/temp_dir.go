package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetTestTempDir() (string, error) {
	baseDir, err := findGoModDir()
	if err != nil {
		return "", err
	}
	tempDir := filepath.Join(baseDir, "tmp")
	if _, err := os.Stat(tempDir); os.IsNotExist(err) {
		err := os.Mkdir(tempDir, 0755)
		if err != nil {
			return "", err
		}
	}
	return tempDir, nil
}

// findGoModDir returns the directory containing the go.mod file from the current working directory
// I did not find a better way to do this
func findGoModDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", fmt.Errorf("go.mod file not found")
		}
		dir = parentDir
	}
}
