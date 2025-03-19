package ioHelpers

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/mt1976/frantic-core/commonErrors"
	"github.com/mt1976/frantic-core/logHandler"
)

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) (string, bool) {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", true
	}
	return string(data), false
}

func MkdirAll(path string) error {
	logHandler.InfoLogger.Printf("[%v] Creating folder Path=[%v]", strings.ToUpper(name), path)
	return os.MkdirAll(path, os.ModeSticky|os.ModePerm)
}

func GetFolders(path string) ([]string, error) {
	// Get all folders in the backup directory
	files, err := os.ReadDir(path)
	if err != nil {
		return nil, commonErrors.WrapOSError(err)
	}
	var folders []string
	for _, file := range files {
		if file.IsDir() {
			folders = append(folders, file.Name())
		}
	}
	return folders, nil
}
