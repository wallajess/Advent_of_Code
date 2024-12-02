package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func getInput() ([]int, []int) {
	var listA []int
	var listB []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return nil, nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		//Because Go was finding invalid lines and ignoring them, check and remove BOM and trim whitespaces

		// Check and remove BOM if it exists
		if strings.HasPrefix(line, "\uFEFF") {
			line = strings.TrimPrefix(line, "\uFEFF")
		}

		// Trim leading/trailing whitespace
		line = strings.TrimSpace(line)
		fields := strings.Fields(line)

		// Ensure there are exactly two fields
		if len(fields) != 2 {
			fmt.Println("Skipping invalid line:", line)
			continue
		}

		// Convert fields to integers
		a, err1 := strconv.Atoi(fields[0])
		b, err2 := strconv.Atoi(fields[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Skipping line with non-integer values:", line)
			continue
		}

		// Append integers to respective lists
		listA = append(listA, a)
		listB = append(listB, b)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, nil
	}

	return listA, listB
}

func distances(listA []int, listB []int) {

	var diffs []int
	var result int

	//Sort the numbers in listA from smallest to largest
	for i := 0; i < len(listA); i++ {
		for j := 0; j < len(listA); j++ {
			if listA[i] < listA[j] {
				listA[i], listA[j] = listA[j], listA[i]
			}
		}
	}

	//Sort the numbers in listA from smallest to largest
	for i := 0; i < len(listB); i++ {
		for j := 0; j < len(listB); j++ {
			if listB[i] < listB[j] {
				listB[i], listB[j] = listB[j], listB[i]
			}
		}
	}

	for i := 0; i < len(listA); i++ {
		diffs = append(diffs, int(math.Abs(float64(listA[i])-float64(listB[i]))))
	}

	for i := 0; i < len(diffs); i++ {
		result += diffs[i]
	}

	fmt.Println(result)
}

func similarity(listA []int, listB []int) {
	var similarityScores []int
	var result int

	for i := 0; i < len(listA); i++ {

		for j := 0; j < len(listA); j++ {
			occurrences := 0
			if listA[i] == listB[j] {
				occurrences++
			}
			similarityScores = append(similarityScores, listA[i]*occurrences)

		}
	}
	for i := 0; i < len(similarityScores); i++ {
		result += similarityScores[i]
	}

	fmt.Println(result)
}

func main() {
	var listA, listB []int
	listA, listB = getInput()
	distances(listA, listB)
	similarity(listA, listB)
}
