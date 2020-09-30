package helpers

import (
	"bufio"
	"os"
)

//FileExists function checks if a file exists in the given path and returns the bool result
func FileExists(filePath string) bool {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//ReadCommandsFromFile extracts the lines of the file and returns an array of strings
func ReadCommandsFromFile(filepath string) []string {
	var result []string
	file, err := os.Open(filepath)
	if err != nil {
		return []string{}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return []string{}
	}
	return result
}
