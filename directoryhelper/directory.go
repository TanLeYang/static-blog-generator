package directoryhelper

import (
	"path"
	"path/filepath"
	"runtime"
)

// RootDir returns the root directory of the project
func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func DataDir() string {
	root := RootDir()
	return filepath.Join(root, "data")
}

func OutputDir() string {
	root := RootDir()
	return filepath.Join(root, "output")
}
