package main

func findGuard(patrolMap [][]string) []int {
	var startingPoint []int

	for i, row := range patrolMap {
		for j, position := range row {
			if position == "^" {
				startingPoint[0], startingPoint[1] = i, j
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

func navigateMap(patrolPath [][]string, startingPoint []int, counter int) ([]int, int) {
	var newPosition []int
	//Determine guard's position
	guard := patrolPath[startingPoint[0]][startingPoint[1]]

	switch guard {
	case "^":
		newPosition = moveUp(patrolPath, startingPoint)
	case ">":
		newPosition = moveLeft(patrolPath, startingPoint)
	case "v":
		newPosition = moveDown(patrolPath, startingPoint)
	case "<":
		newPosition = moveRight(patrolPath, startingPoint)
	}
	if patrolPath[newPosition[0]][newPosition[1]] != "#" {
		patrolPath[newPosition[0]][newPosition[1]] = guard
		patrolPath[startingPoint[0]][startingPoint[1]] = "X"
		counter++

	}
	if patrolPath[newPosition[0]][newPosition[1]] == "#" {
		newPosition = startingPoint
		guard = turn90degrees(patrolPath, startingPoint)
	}
	return newPosition, counter
}
