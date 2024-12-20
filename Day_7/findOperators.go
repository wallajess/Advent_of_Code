package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func calculateCombinations(list []int, index int, currentResult int, results []int) []int {
	if index == len(list) {
		return append(results, currentResult)
	}
	currentNumber := list[index]

	results = calculateCombinations(list, index+1, currentResult, results)
	results = calculateCombinations(list, index+1, currentNumber+currentResult, results)

	if index == 0 {
		results = calculateCombinations(list, index+1, currentNumber, results)
	} else {
		results = calculateCombinations(list, index+1, currentResult*currentNumber, results)

	}

	return results
}

func findOperators(input [][]string) []int {
	var correctTotals []int

	for _, row := range input {

		var nums []int
		var results []int

		if len(row) < 2 {
			fmt.Printf("Skipping invalid row: %v\n", row)
			continue
		}

		lineTotal, err := strconv.Atoi(row[0])
		if err != nil {
			fmt.Printf("Skipping row due to conversion error: %v\n", row)
			continue
		}
		numStrings := strings.Split(row[1], " ")

		for _, num := range numStrings {
			numInt, e := strconv.Atoi(num)
			if e == nil {
				nums = append(nums, numInt)
			}
		}

		results = calculateCombinations(nums, 0, 0, results)

		if slices.Contains(results, lineTotal) {
			correctTotals = append(correctTotals, lineTotal)
		}

	}
	return correctTotals
}
