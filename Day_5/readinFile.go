package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getOrderingRules(text string) [][]int {

	file, err := os.Open(text)
	if err != nil {
		fmt.Println("Error opening file '%s': %v\n", text, err)
		return nil
	}

	defer file.Close()

	// Read the file line by line
	var orderingRules [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//Split the line by the delimiter "|"
		parts := strings.Split(line, "|")
		if len(parts) != 2 {
			fmt.Printf("Skipping malformed line: %s\n", line)
			continue
		}

		// Convert each part to an integer
		row := make([]int, 2)
		for i, part := range parts {
			num, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				fmt.Printf("Skipping line with invalid number: %s\n", line)
				continue
			}
			row[i] = num
		}

		orderingRules = append(orderingRules, row)

	}

	// Handle errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	return orderingRules
}

func getUpdatePages(text string) [][]int {
	file, err := os.Open(text)
	if err != nil {
		fmt.Printf("Error opening file '%s': %v\n", file, err)
		return nil
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		//Split the line by the delimiter ","
		parts := strings.Split(line, ",")

		var row []int
		for _, part := range parts {
			// Convert each part to an integer
			num, err := strconv.Atoi(part)
			if err != nil {
				fmt.Printf("Skipping invalid number '%s' in line: %s\n", part, line)
				continue
			}
			row = append(row, num)
		}

		// Append the row of integers to the result
		result = append(result, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file '%s': %v\n", file, err)
		return nil
	}

	return result
}
