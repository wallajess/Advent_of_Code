package main

import (
	"fmt"
	"slices"
)

func correctRows(orderingRules [][]int, incorrect [][]int) [][]int {
	var correctedRows [][]int

	for _, row := range incorrect {
		correctedRow := row

		slices.SortFunc(correctedRow, func(a, b, int) int {
			for i, rule := range orderingRules {
			//Check if both numbers are in the row
				fmt.Printf("Rule is %v\n", rule)
				fmt.Printf("CorrectedRow is %v\n", correctedRow)
				if slices.Contains(correctedRow, rule[0]) && slices.Contains(correctedRow, rule[1]) {
					####
			}

			if i == len(orderingRules)-1 {
				correctedRows = append(correctedRows, correctedRow)
			}

		}

	}
	return correctedRows
}
