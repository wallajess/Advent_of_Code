package main

import "fmt"

func main() {
	patrolPath := getMap("test_map.txt")

	startingPoint := findGuard(patrolPath)
	row, col := startingPoint[0], startingPoint[1]
	counter := 0

	for {
		navigateMap(patrolPath, startingPoint, counter)
		counter++
		if row < 0 || row >= len(patrolPath) || col < 0 || col >= len(patrolPath[0]) {
			break
		}
	}

	fmt.Println()
}
