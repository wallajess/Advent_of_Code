package main

import "fmt"

func main() {
	orderingRules := getOrderingRules("test_ordering.txt")
	fmt.Printf("Ordering rules: %v\n", orderingRules)
	updatePages := getUpdatePages("test_update.txt")
	fmt.Printf("Update pages: %v\n", updatePages)
	correctOrder := checkUpdateOrder(orderingRules, updatePages)

	var result int
	for _, midNum := range correctOrder {

		result += midNum
		fmt.Println("Middle number:", midNum)
	}
	fmt.Println(result)

}
