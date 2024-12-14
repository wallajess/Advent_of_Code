package main

import (
	"fmt"
	"slices"
)

func checkUpdateOrder(orderingRules [][]int, updatePages [][]int) []int {
	var middleNums []int

	for _, row := range updatePages {

		correct := true

		for _, rule := range orderingRules {
			//Check if both numbers are in the row
			if slices.Contains(row, rule[0]) && slices.Contains(row, rule[1]) {
				//Ensure correct order of rule numbers
				if slices.Index(row, rule[0]) > slices.Index(row, rule[1]) {
					correct = false
				}

			}

		}
		if correct {
			fmt.Printf("This row is correct for all rules: %v\n", row)
			middleNums = append(middleNums, row[len(row)/2])
		}

	}
	return middleNums
}
