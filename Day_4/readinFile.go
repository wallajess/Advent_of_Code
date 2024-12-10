package main

import (
	"bufio"
	"fmt"
	"os"
)

func getInput(text string) [][]string {

	file, err := os.Open(text)
	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return [][]string{}
	}

	defer file.Close()

	// Read the file line by line
	var wordSearch [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//Convert the line into a slice of single character strings
		characters := []string{}
		for _, char := range line {
			characters = append(characters, string(char))
		}
		//Append the slice of characters to wordSearch
		wordSearch = append(wordSearch, characters)
	}

	// Handle errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return [][]string{}
	}

	return wordSearch
}
