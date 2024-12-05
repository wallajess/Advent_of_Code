package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func findMul(text string) [][]string {
	var numsToMultiply [][]string

	//First find mul(int, int)
	pattern := "mul\\(\\d{1,3}\\,\\d{1,3}\\)"

	//compile the regex
	re, err := regexp.Compile(pattern)

	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return [][]string{}
	}

	//find all matches
	allMatches := re.FindAllString(text, -1)

	patternNums := "\\d{1,3}"

	for _, value := range allMatches {
		//compile the regex
		re, err := regexp.Compile(patternNums)

		if err != nil {
			fmt.Println("Error compiling regex:", err)
			return [][]string{}
		}

		//check if pattern matches
		numMatches := re.FindAllString(value, -1)
		numsToMultiply = append(numsToMultiply, numMatches)

	}
	return numsToMultiply
}

func multiply(list [][]string) int {
	result := 0

	for _, num := range list {
		var num1, num2 int
		num1, _ = (strconv.Atoi(num[0]))
		num2, _ = (strconv.Atoi(num[1]))
		result += num1 * num2
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println("Error opening input.txt", err)
		return
	}
	defer file.Close()

	// Read the file line by line
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Handle errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Combine all lines into a single string
	textToSearch := strings.Join(lines, " ")

	muls := findMul(textToSearch)
	result := multiply(muls)

	fmt.Println(result)

}
