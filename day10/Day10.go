package day10

const startPoint = "S"

type Point struct {
	y     int
	x     int
	value string
}

func PartOne(input []string) (result int) {
	start := findStart(input)
	loop := findLoop(start, input)

	return len(loop) / 2
}

func PartTwo(input []string) (result int) {
	return 1
}

func findStart(input []string) (start Point) {
	for i := 0; i < len(input); i++ {
		for z := 0; z < len(input[i]); z++ {
			if string(input[i][z]) == startPoint {
				start.y = i
				start.x = z
				start.value = "S"
				return
			}
		}
	}
	return
}

func findLoop(start Point, input []string) (loop []Point) {

	currentStep := start
	var nextStep = Point{}
	loop = append(loop, start)
	for nextStep.value != "NULL" {
		nextStep = findNextStep(currentStep, input, loop)
		if nextStep.value != "NULL" {
			currentStep = nextStep
			loop = append(loop, nextStep)
		}
	}

	return
}

func findNextStep(currentStep Point, input []string, loop []Point) Point {

	possibleDirections := getPossibleDirections(currentStep)

	var nextStep = Point{}
	for _, direction := range possibleDirections {
		if direction == "left" {
			nextStep = getLeft(currentStep, input)
			if nextStep.value != "NULL" && canWeGoTo(nextStep, direction) && !loopContainsPoint(loop, nextStep) {
				return nextStep
			}
		}
		if direction == "right" {
			nextStep = getRight(currentStep, input)
			if nextStep.value != "NULL" && canWeGoTo(nextStep, direction) && !loopContainsPoint(loop, nextStep) {
				return nextStep
			}
		}
		if direction == "up" {
			nextStep = getUp(currentStep, input)
			if nextStep.value != "NULL" && canWeGoTo(nextStep, direction) && !loopContainsPoint(loop, nextStep) {
				return nextStep
			}
		}
		if direction == "down" {
			nextStep = getDown(currentStep, input)
			if nextStep.value != "NULL" && canWeGoTo(nextStep, direction) && !loopContainsPoint(loop, nextStep) {
				return nextStep
			}
		}
	}

	return Point{-1, -1, "NULL"}
}

func getPossibleDirections(currentPoint Point) (directions []string) {
	switch currentPoint.value {
	case "-":
		return []string{"left", "right"}

	case "|":
		return []string{"up", "down"}

	case "L":
		return []string{"up", "right"}

	case "J":
		return []string{"up", "left"}

	case "7":
		return []string{"left", "down"}

	case "F":
		return []string{"right", "down"}

	case "S":
		return []string{"up", "left", "right", "down"}
	}
	return []string{}
}

func canWeGoTo(to Point, direction string) bool {
	if direction == "right" {
		return to.value == "-" || to.value == "J" || to.value == "7"
	}
	if direction == "left" {
		return to.value == "-" || to.value == "L" || to.value == "F"
	}
	if direction == "up" {
		return to.value == "|" || to.value == "7" || to.value == "F"
	}
	if direction == "down" {
		return to.value == "|" || to.value == "J" || to.value == "L"
	}
	return false
}

func loopContainsPoint(loop []Point, point Point) bool {
	for _, v := range loop {
		if v.x == point.x && v.y == point.y {
			return true
		}
	}
	return false
}

func getLeft(currentPoint Point, input []string) (leftPoint Point) {
	if currentPoint.x-1 >= 0 {
		leftPoint.y = currentPoint.y
		leftPoint.x = currentPoint.x - 1
		leftPoint.value = string(input[leftPoint.y][leftPoint.x])
	} else {
		leftPoint.value = "NULL"
		leftPoint.y = -1
		leftPoint.x = -1
	}
	return
}

func getRight(currentPoint Point, input []string) (leftPoint Point) {
	if currentPoint.x+1 < len(input[currentPoint.y]) {
		leftPoint.y = currentPoint.y
		leftPoint.x = currentPoint.x + 1
		leftPoint.value = string(input[leftPoint.y][leftPoint.x])
	} else {
		leftPoint.value = "NULL"
		leftPoint.y = -1
		leftPoint.x = -1
	}
	return
}

func getUp(currentPoint Point, input []string) (leftPoint Point) {
	if currentPoint.y-1 >= 0 {
		leftPoint.y = currentPoint.y - 1
		leftPoint.x = currentPoint.x
		leftPoint.value = string(input[leftPoint.y][leftPoint.x])
	} else {
		leftPoint.value = "NULL"
		leftPoint.y = -1
		leftPoint.x = -1
	}
	return
}

func getDown(currentPoint Point, input []string) (leftPoint Point) {
	if currentPoint.y+1 < len(input) {
		leftPoint.y = currentPoint.y + 1
		leftPoint.x = currentPoint.x
		leftPoint.value = string(input[leftPoint.y][leftPoint.x])
	} else {
		leftPoint.value = "NULL"
		leftPoint.y = -1
		leftPoint.x = -1
	}
	return
}
