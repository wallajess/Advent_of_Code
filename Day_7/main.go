package main

import "fmt"

func main() {
	callibrations := getCallibrations("callibrations.txt")

	for _, callibration := range callibrations {
		println(len(callibration))
	}

	results := findOperators(callibrations)
	var solution int

	for _, total := range results {
		solution += total
	}
	fmt.Println(solution)
}
