package utils

import (
	"os"
	"path/filepath"
)

func ReadFile(filename string) string {
	filePath := filepath.Join("data", filename)
	fileContents, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	return string(fileContents)
}

func ReadFilesInDirectory(directoryPath string) (string, error) {
	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return "", err
	}

	var result string

	for _, file := range files {
		if !file.IsDir() {
			filePath := filepath.Join(directoryPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				return "", err
			}
			result += string(content)
		}
	}

	return result, nil
}
