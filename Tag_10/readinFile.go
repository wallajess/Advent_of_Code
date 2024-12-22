package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getMap(text string) [][]string {

	file, err := os.Open(text)
	if err != nil {
		fmt.Println("Error opening file '%s': %v\n", text, err)
		return nil
	}

	defer file.Close()

	// Read the file line by line
	var topoMap [][]string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//Split the line by the delimiter ""
		parts := strings.Split(line, "")

		topoMap = append(topoMap, parts)

	}

	// Handle errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return topoMap
}
