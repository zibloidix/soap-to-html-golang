package main

import (
	"bufio"
	"os"
	"strings"
)

func isWSDL(fileName string) bool {
	return strings.Contains(fileName, ".wsdl") || strings.Contains(fileName, ".xsd")
}

func isXSD(fileName string) bool {
	return strings.Contains(fileName, ".wsdl") || strings.Contains(fileName, ".xsd")
}

func getIgnorePathes(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pathes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pathes = append(pathes, scanner.Text())
	}
	return pathes, scanner.Err()
}

func isNotIgnoredPath(path string, ignoredPathes []string) bool {
	for _, ignoredPath := range ignoredPathes {
		if strings.Contains(path, ignoredPath) {
			return false
		}
	}
	return true
}

func isValidFile(path string, info os.FileInfo) bool {
	ignoredPathes, _ := getIgnorePathes(".ignore")
	return isNotIgnoredPath(path, ignoredPathes) && isNotDir(info) && isValidFileType(info)
}

func isNotDir(info os.FileInfo) bool {
	return !info.IsDir()
}

func isValidFileType(info os.FileInfo) bool {
	return isWSDL(info.Name()) || isXSD(info.Name())
}
