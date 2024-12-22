package main

import "fmt"

func main() {
	discmap := getDiscmap("discmap.txt")

	result := fileCompactor(discmap)
	fmt.Println(result)
}
