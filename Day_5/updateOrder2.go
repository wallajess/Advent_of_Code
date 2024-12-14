package main

import (
	"fmt"
	"slices"
)

func checkUpdateOrder2(orderingRules [][]int, updatePages [][]int) ([]int, []int) {
	var middleNums []int
	var correctedSlices [][]int
	var middleNumsCorrected []int

	for _, row := range updatePages {

		correct := true

		for _, rule := range orderingRules {
			//Check if both numbers are in the row
			if slices.Contains(row, rule[0]) && slices.Contains(row, rule[1]) {
				//Ensure correct order of rule numbers
				if slices.Index(row, rule[0]) > slices.Index(row, rule[1]) {
					correct = false

					//Correct row by swapping numbers
					corrected := make([]int, len(row))
					idx1, idx2 := slices.Index(row, rule[0]), slices.Index(row, rule[1])
					corrected[idx1], corrected[idx2] = corrected[idx2], corrected[idx1]

					//Append to corrected row
					correctedSlices = append(correctedSlices, corrected)
				}
			}

		}
		if correct {
			fmt.Printf("This row is correct for all rules: %v\n", row)
			middleNums = append(middleNums, row[len(row)/2])
		}
	}

	//Remove duplicates
	var noDuplicates [][]int

	for _, correctedSlice := range correctedSlices {
		isDuplicate := false
		for _, slice := range noDuplicates {
			if slices.Equal(slice, correctedSlice) {
				isDuplicate = true
			}

		}
		if !isDuplicate {
			noDuplicates = append(noDuplicates, correctedSlice)
		}
	}

	for _, correctedSlice := range noDuplicates {
		middleNumsCorrected = append(middleNumsCorrected, correctedSlice[len(correctedSlice)/2])
	}
	return middleNums, middleNumsCorrected
}
