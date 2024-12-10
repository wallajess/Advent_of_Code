package main

func main() {
	lines := getInput("input.txt")
	println(searchWord(lines, "XMAS"))
	println(crossSearch(lines, "MAS"))
}
