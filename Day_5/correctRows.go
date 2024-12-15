package main

import (
	"fmt"
	"slices"
)

func correctRows(orderingRules [][]int, incorrect [][]int) [][]int {
	var correctedRows [][]int

	for _, row := range incorrect {
		correctedRow := row

		for i, rule := range orderingRules {
			//Check if both numbers are in the row
			fmt.Printf("Rule is %v\n", rule)
			fmt.Printf("CorrectedRow is %v\n", correctedRow)
			if slices.Contains(correctedRow, rule[0]) && slices.Contains(correctedRow, rule[1]) {
				fmt.Printf("The rule is in the row.\n")
				firstIndex := slices.Index(correctedRow, rule[0])
				secondIndex := slices.Index(correctedRow, rule[1])
				if firstIndex > secondIndex {
					correctedRow[firstIndex], correctedRow[secondIndex] = correctedRow[secondIndex], correctedRow[firstIndex]
					fmt.Printf("CorrectedRow now looks like this: %v\n", correctedRow)

				}

			}

			if i == len(orderingRules)-1 {
				correctedRows = append(correctedRows, correctedRow)
			}

		}

	}
	return correctedRows
}
