package main

func findGuard(patrolMap [][]string) []int {
	var startingPoint []int

	for i, row := range patrolMap {
		for j, position := range row {
			if position == "^" {
				startingPoint = make([]int, i, j)
			}
		}
	}
	return startingPoint
}

func moveUp(patrolPath [][]string, startingPoint []int) []int {
	patrolPath[startingPoint[0]][startingPoint[1]] = "X"
	newPosition := make([]int, startingPoint[0]+1, startingPoint[1])
	return newPosition
}

func moveDown(patrolPath [][]string, startingPoint []int) []int {
	patrolPath[startingPoint[0]][startingPoint[1]] = "X"
	newPosition := make([]int, startingPoint[0]-1, startingPoint[1])
	return newPosition
}

func moveRight(patrolPath [][]string, startingPoint []int) []int {
	patrolPath[startingPoint[0]][startingPoint[1]] = "X"
	newPosition := make([]int, startingPoint[0], startingPoint[1]+1)
	return newPosition
}

func moveLeft(patrolPath [][]string, startingPoint []int) []int {
	patrolPath[startingPoint[0]][startingPoint[1]] = "X"
	newPosition := make([]int, startingPoint[0], startingPoint[1]-1)
	return newPosition
}

func turn90degrees(patrolPath [][]string, startingPoint []int) string {
	guard := patrolPath[startingPoint[0]][startingPoint[1]]
	switch guard {
	case "^":
		guard = ">"
	case ">":
		guard = "v"
	case "v":
		guard = "<"
	case "<":
		guard = "^"

	}
	return guard
}
