package main

func crossSearch(text [][]string, word string) int {
	counter := 0

	for i := range text {
		for j, letter := range text[i] {
			if letter == string(word[1]) {
				//test Ms on top
				if isValid(text, i-1, j-1) && isValid(text, i-1, j+1) &&
					isValid(text, i+1, j-1) && isValid(text, i+1, j+1) &&
					text[i-1][j-1] == string(word[0]) &&
					text[i-1][j+1] == string(word[0]) &&
					text[i+1][j-1] == string(word[2]) &&
					text[i+1][j+1] == string(word[2]) {
					counter++
				}
				//test Ms on bottom
				if isValid(text, i-1, j-1) && isValid(text, i-1, j+1) &&
					isValid(text, i+1, j-1) && isValid(text, i+1, j+1) &&
					text[i+1][j-1] == string(word[0]) &&
					text[i+1][j+1] == string(word[0]) &&
					text[i-1][j-1] == string(word[2]) &&
					text[i-1][j+1] == string(word[2]) {
					counter++
				}
				//test Ms left
				if isValid(text, i-1, j-1) && isValid(text, i-1, j+1) &&
					isValid(text, i+1, j-1) && isValid(text, i+1, j+1) &&
					text[i-1][j-1] == string(word[0]) &&
					text[i-1][j+1] == string(word[2]) &&
					text[i+1][j-1] == string(word[0]) &&
					text[i+1][j+1] == string(word[2]) {
					counter++
				}
				//test Ms right
				if isValid(text, i-1, j-1) && isValid(text, i-1, j+1) &&
					isValid(text, i+1, j-1) && isValid(text, i+1, j+1) &&
					text[i-1][j-1] == string(word[2]) &&
					text[i-1][j+1] == string(word[0]) &&
					text[i+1][j-1] == string(word[2]) &&
					text[i+1][j+1] == string(word[0]) {
					counter++
				}
			}
		}
	}
	return counter
}
