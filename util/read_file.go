package util

import (
	"bufio"
	"io"
	"os"
)

func ReadFileLines(path string) []string {
	fileHandle := openFile(path)
	defer fileHandle.Close()

	var lines []string
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

func ReadFileString(path string) string {
	fileHandle := openFile(path)
	defer fileHandle.Close()

	content, err := io.ReadAll(fileHandle)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func openFile(path string) *os.File {
	fileHandle, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return fileHandle
}
