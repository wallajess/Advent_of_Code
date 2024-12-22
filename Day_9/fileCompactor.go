package main

import (
	"fmt"
	"strconv"
)

func fileCompactor(discmap string) int {

	var discmapSlice []string

	//Iterate over each item in the discmap
	for i, char := range discmap {
		//Convert num to int
		number, err := strconv.Atoi(string(char))
		if err != nil {
			fmt.Printf("Invalid character in discmap: %c", char)
			continue
		}
		//If the index is even, it is a file. The number says how large the block is
		if i%2 == 0 {
			//Files have an ID number corresponding to the order the files appear, starting with 0
			//id := 0
			fmt.Println(number)
			for j := 0; j < number; j++ {
				discmapSlice = append(discmapSlice, strconv.Itoa(i/2))
			}
		} else {
			for k := 0; k < number; k++ {
				discmapSlice = append(discmapSlice, ".")
			}

		}

	}
	fmt.Printf("Discmap before it has been changed around; %v\n", discmapSlice)

	for i := 0; i < len(discmapSlice); i++ {
		lastIndex := len(discmapSlice) - 1
		if discmapSlice[i] == "." {
			if discmapSlice[lastIndex] == "." {
				lastIndex = lastIndex - 1
			}

			discmapSlice[i] = discmapSlice[lastIndex]
			discmapSlice = discmapSlice[:lastIndex]
		}
	}
	fmt.Println(discmapSlice)

	var checksum int

	for i, num := range discmapSlice {
		var intNum int
		intNum, _ = strconv.Atoi(string(num))
		checksum = checksum + (i * intNum)
	}

	return checksum
}
