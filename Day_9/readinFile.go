package main

import (
	"bufio"
	"fmt"
	"os"
)

func getDiscmap(text string) string {

	file, err := os.Open(text)
	if err != nil {
		fmt.Println("Error opening file '%s': %v\n", text, err)
		return "Terminated due to error"
	}

	defer file.Close()

	// Read the file line by line
	var discmap string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		discmap = scanner.Text()
	}

	// Handle errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return "Terminated due to error"
	}

	return discmap
}
