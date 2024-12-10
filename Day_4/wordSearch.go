package main

// Helper function to make sure we don't go out of bounds
func isValid(text [][]string, i, j int) bool {
	return i >= 0 && i < len(text) && j >= 0 && j < len(text[i])
}

func searchWord(text [][]string, word string) int {
	counter := 0

	for i := range text {
		for j, letter := range text[i] {
			if letter == string(word[0]) {
				//test horizontal forwards
				if isValid(text, i, j+1) && isValid(text, i, j+2) && isValid(text, i, j+3) &&
					text[i][j+1] == string(word[1]) &&
					text[i][j+2] == string(word[2]) &&
					text[i][j+3] == string(word[3]) {
					counter++
				}
				//test horizontal backwards
				if isValid(text, i, j-1) && isValid(text, i, j-2) && isValid(text, i, j-3) &&
					text[i][j-1] == string(word[1]) &&
					text[i][j-2] == string(word[2]) &&
					text[i][j-3] == string(word[3]) {
					counter++
				}
				//test diagonal up backwards
				if isValid(text, i-1, j-1) && isValid(text, i-2, j-2) && isValid(text, i-3, j-3) &&
					text[i-1][j-1] == string(word[1]) &&
					text[i-2][j-2] == string(word[2]) &&
					text[i-3][j-3] == string(word[3]) {
					counter++
				}
				//test diagonal up forwards
				if isValid(text, i-1, j+1) && isValid(text, i-2, j+2) && isValid(text, i-3, j+3) &&
					text[i-1][j+1] == string(word[1]) &&
					text[i-2][j+2] == string(word[2]) &&
					text[i-3][j+3] == string(word[3]) {
					counter++
				}
				//test diagonal down backwards
				if isValid(text, i+1, j-1) && isValid(text, i+2, j-2) && isValid(text, i+3, j-3) &&
					text[i+1][j-1] == string(word[1]) &&
					text[i+2][j-2] == string(word[2]) &&
					text[i+3][j-3] == string(word[3]) {
					counter++
				}
				//test diagonal down forwards
				if isValid(text, i+1, j+1) && isValid(text, i+2, j+2) && isValid(text, i+3, j+3) &&
					text[i+1][j+1] == string(word[1]) &&
					text[i+2][j+2] == string(word[2]) &&
					text[i+3][j+3] == string(word[3]) {
					counter++
				}
				//test vertical down
				if isValid(text, i+1, j) && isValid(text, i+2, j) && isValid(text, i+3, j) &&
					text[i+1][j] == string(word[1]) &&
					text[i+2][j] == string(word[2]) &&
					text[i+3][j] == string(word[3]) {
					counter++
				}
				//test vertical up
				if isValid(text, i-1, j) && isValid(text, i-2, j) && isValid(text, i-3, j) &&
					text[i-1][j] == string(word[1]) &&
					text[i-2][j] == string(word[2]) &&
					text[i-3][j] == string(word[3]) {
					counter++

				}
			}
		}
	}
	return counter
}
