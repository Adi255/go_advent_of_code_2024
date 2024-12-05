package util

import (
	"bufio"
	"os"
)

func ReadFileLines(path string) []string {
	fileHandle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	var lines []string
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
