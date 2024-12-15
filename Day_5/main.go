package main

import "fmt"

func main() {
	orderingRules := getOrderingRules("ordering.txt")
	fmt.Printf("Ordering rules: %v\n", orderingRules)
	updatePages := getUpdatePages("update.txt")
	fmt.Printf("Update pages: %v\n", updatePages)
	correctOrder, incorrect := checkUpdateOrder(orderingRules, updatePages)

	var result int
	for _, midNum := range correctOrder {

		result += midNum
		fmt.Println("Middle number:", midNum)
	}
	fmt.Println(result)

	var middleCorrected []int
	correctedRows := correctRows(orderingRules, incorrect)
	for _, row := range correctedRows {
		middleCorrected = append(middleCorrected, row[len(row)/2])
	}

	var sumCorrected int
	for _, midNum := range middleCorrected {
		sumCorrected += midNum
		fmt.Println("Middle number:", midNum)
	}

	fmt.Println(sumCorrected)

}
