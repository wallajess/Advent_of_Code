package main

func main() {
	patrolMap := getMap("test_map.txt")

	//Determine where the guard is
	position := findGuard(patrolMap)

	//Determine guard's position
	guard = patrolMap[position[0]][position[1]]

	switch guard {
	case "^":
		newPosition := moveUp(patrolMap, position)
		if patrolMap[newPosition[0]][newPosition[1]] == "." {
			
		}

		}
		newPosition = moveUp
	case ">":
		guard = "v"
	case "v":
		guard = "<"
	case "<":
		guard = "^"

	}
}
