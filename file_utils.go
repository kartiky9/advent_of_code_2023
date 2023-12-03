package fileutils

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filePath string) []string {
	var lines []string

	// Open the file.
	// Replace "yourfile.txt" with the path to the file you want to read.
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// Loop over all lines in the file.
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Check for any errors encountered while reading the file.
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	return lines
}
