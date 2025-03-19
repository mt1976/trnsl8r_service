package ioHelpers

import (
	"os"
)

// FileExists returns true if the specified file existing on the filesystem
func fileExists(filename string) bool {
	return touch(filename)
}

// Touch returns true if the specified file existing on the filesystem
func touch(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
