package main

func findGuard(patrolMap [][]string) []int {
	var startingPoint = make([]int, 2)

	for i, row := range patrolMap {
		for j, position := range row {
			if position == "^" {
				startingPoint[0], startingPoint[1] = i, j
			}
		}
	}
	return startingPoint
}

func moveUp(startingPoint []int) []int {

	newPosition := []int{startingPoint[0] - 1, startingPoint[1]}
	return newPosition
}

func moveDown(startingPoint []int) []int {

	newPosition := []int{startingPoint[0] + 1, startingPoint[1]}
	return newPosition
}

func moveRight(startingPoint []int) []int {

	newPosition := []int{startingPoint[0], startingPoint[1] - 1}
	return newPosition
}

func moveLeft(startingPoint []int) []int {

	newPosition := []int{startingPoint[0], startingPoint[1] + 1}
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

func navigateMap(patrolPath [][]string, startingPoint []int) ([]int, bool) {
	var newPosition = make([]int, 2)
	//Determine guard's position
	guard := patrolPath[startingPoint[0]][startingPoint[1]]

	if patrolPath[newPosition[0]][newPosition[1]] != "#" {
		switch guard {
		case "^":
			newPosition = moveUp(startingPoint)
		case ">":
			newPosition = moveLeft(startingPoint)
		case "v":
			newPosition = moveDown(startingPoint)
		case "<":
			newPosition = moveRight(startingPoint)
		}

		// Check for out-of-bounds
		if newPosition[0] < 0 || newPosition[0] >= len(patrolPath) ||
			newPosition[1] < 0 || newPosition[1] >= len(patrolPath[0]) {
			patrolPath[startingPoint[0]][startingPoint[1]] = "X"

			// End program when out of bounds

			return newPosition, true
		}

		//Check for obstacles
		if patrolPath[newPosition[0]][newPosition[1]] == "#" {
			guard = turn90degrees(patrolPath, startingPoint)
			patrolPath[startingPoint[0]][startingPoint[1]] = guard
			return startingPoint, false
		}

		if patrolPath[startingPoint[0]][startingPoint[1]] != "X" {

			patrolPath[startingPoint[0]][startingPoint[1]] = "X" // Mark the previous position as visited

		}
		patrolPath[newPosition[0]][newPosition[1]] = guard
	}
	return newPosition, false
}
