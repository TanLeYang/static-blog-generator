package directoryhelper

import (
	"log"
	"os"
	"path"
	"path/filepath"
)

// RootDir returns the root directory of the project
func RootDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	return path.Dir(ex)
}

func DataDir() string {
	root := RootDir()
	return filepath.Join(root, "data")
}

func OutputDir() string {
	root := RootDir()
	return filepath.Join(root, "output")
}
