package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isAscending(a []int) bool {
	ascendingSlice := make([]int, len(a))
	copy(ascendingSlice, a)
	sort.Ints(ascendingSlice)

	if Equal(ascendingSlice, a) {
		return true
	}
	return false
}

func isDescending(a []int) bool {
	descendingSlice := make([]int, len(a))
	copy(descendingSlice, a)
	sort.Sort(sort.Reverse(sort.IntSlice(descendingSlice)))

	if Equal(descendingSlice, a) {
		return true
	}
	return false
}
func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return
	}
	defer file.Close()

	var reports [][]int
	var numSafe int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Parse each field into an integer and store in a slice
		var lineNumbers []int
		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error parsing number:", err)
				continue
			}
			lineNumbers = append(lineNumbers, num)
		}

		// Add the line of integers to the main slice
		reports = append(reports, lineNumbers)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	for _, report := range reports {

		levelUnsafe := 0

		//Make sure there are no empty reports
		if len(report) == 0 {
			continue
		}

		//Iterate over the reports and calculate the difference between the numbers
		for i := 1; i < len(report); i++ {
			previousNum := report[i-1]
			diff := math.Abs(float64(report[i] - previousNum))

			//check if the difference between the numbers is between 1 and 3
			if diff == 0 || diff > 3 {
				levelUnsafe++
			}
		}

		//test if a report is neither ascending nor descending
		if !(isAscending(report) || isDescending(report)) {
			levelUnsafe++
		}

		isSafe := levelUnsafe <= 0
		if isSafe {
			numSafe++

		}
		fmt.Printf("This report %v is safe: %v and has %v unsafe levels\n", report, isSafe, levelUnsafe)
	}

	println("Num of safe reports:", numSafe)
}
