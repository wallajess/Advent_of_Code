package main

import "fmt"

func main() {
	patrolPath := getMap("map.txt")

	startingPoint := findGuard(patrolPath)

	outOfBounds := false

	for !outOfBounds {
		var newPosition []int
		newPosition, outOfBounds = navigateMap(patrolPath, startingPoint)

		startingPoint = newPosition
	}
	var distinctPos int
	for i, row := range patrolPath {
		for j, _ := range row {
			if patrolPath[i][j] == "X" {
				distinctPos++
			}
		}
	}

	fmt.Println(distinctPos)
	fmt.Println("Final Patrol Path:")
	for _, row := range patrolPath {
		fmt.Println(row)
	}
}
