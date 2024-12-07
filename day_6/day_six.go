package day6

import (
	"advent_of_code/util"
	"fmt"
	"reflect"
	"slices"
)

const guardUp = '^'
const guardRight = '>'
const guardDown = 'v'
const guardLeft = '<'
const visited = 'X'
const obstacle = '#'

var turnMap = map[byte]byte{guardUp: guardRight, guardRight: guardDown, guardDown: guardLeft, guardLeft: guardUp}

var guardPositions = []byte{guardUp, guardRight, guardDown, guardLeft}
var PotentialObstaclePositions = map[[2]int][2]int{}

type Heading struct {
	Position  [2]int
	direction byte
}

func LoadGuardMap(path string) [][]byte {
	mapLines := util.ReadFileLines(path)

	var guardMap [][]byte
	for _, line := range mapLines {
		guardMap = append(guardMap, []byte(line))
	}

	return guardMap
}

func TraverseMap(guardMap [][]byte) (encounteredHeadings []Heading, couldExit bool) {
	_, nextPosition := FindGuard(guardMap)
	heading := Heading{nextPosition, guardUp}
	encounteredHeadings = []Heading{heading}
	var oob bool
	for {
		heading, oob = GuardMove(guardMap, heading)
		if oob {
			return encounteredHeadings, true
		}
		if slices.Contains(encounteredHeadings, heading) {
			return encounteredHeadings, false
		}
		encounteredHeadings = append(encounteredHeadings, heading)
	}
}

func GuardMove(guardMap [][]byte, heading Heading) (Heading, bool) {
	direction := heading.direction
	currentPosition := heading.Position
	nextPosition := determineNextPosition(guardMap, heading)
	if guardOoB(guardMap, nextPosition) {
		return heading, true
	}
	if guardMap[nextPosition[0]][nextPosition[1]] != obstacle {
		return Heading{nextPosition, direction}, false
	} else {
		return Heading{currentPosition, turnMap[direction]}, false
	}
}

func guardOoB(guardMap [][]byte, position [2]int) bool {
	return position[0] < 0 || position[0] >= len(guardMap) || position[1] < 0 || position[1] >= len(guardMap[0])
}

func determineNextPosition(guardMap [][]byte, heading Heading) [2]int {
	currentPosition := heading.Position
	switch heading.direction {
	case guardUp:
		return [2]int{currentPosition[0] - 1, currentPosition[1]}
	case guardRight:
		return [2]int{currentPosition[0], currentPosition[1] + 1}
	case guardDown:
		return [2]int{currentPosition[0] + 1, currentPosition[1]}
	case guardLeft:
		return [2]int{currentPosition[0], currentPosition[1] - 1}
	default:
		panic("Invalid direction")
	}
}

func DeepCopyMap(original [][]byte) [][]byte {
	deepCopy := make([][]byte, len(original))
	for i := range original {
		deepCopy[i] = make([]byte, len(original[i]))
		for j := range original[i] {
			deepCopy[i][j] = original[i][j]
		}
	}
	return deepCopy
}

func FindVisitedPositions(guardMap [][]byte) [][2]int {
	visitedPositions := make([][2]int, 0)
	for rowNum, row := range guardMap {
		for colNum, character := range row {
			if character == visited {
				visitedPositions = append(visitedPositions, [2]int{rowNum, colNum})
			}
		}
	}
	return visitedPositions
}

func GuardStep(guardMap [][]byte, heading Heading) ([][]byte, Heading) {
	if LeavingMap(guardMap, heading) {
		return FinalMap(guardMap, heading), heading
	}
	currentPosition := heading.Position
	direction := heading.direction
	var nextPosition [2]int
	var newHeading Heading
	switch direction {
	case guardUp:
		nextPosition = [2]int{currentPosition[0] - 1, currentPosition[1]}
		canMoveUp := guardMap[nextPosition[0]][nextPosition[1]] != obstacle
		if canMoveUp {
			guardMap = MoveUpOnMap(guardMap, currentPosition)
			newHeading = Heading{nextPosition, guardUp}
		} else {
			guardMap = TurnRight(guardMap, currentPosition, direction)
			newHeading = Heading{currentPosition, guardRight}
		}
	case guardRight:
		nextPosition = [2]int{currentPosition[0], currentPosition[1] + 1}
		canMoveRight := guardMap[nextPosition[0]][nextPosition[1]] != obstacle
		if canMoveRight {
			guardMap = MoveRightOnMap(guardMap, currentPosition)
			newHeading = Heading{nextPosition, guardRight}
		} else {
			guardMap = TurnRight(guardMap, currentPosition, direction)
			newHeading = Heading{currentPosition, guardDown}
		}
	case guardDown:
		nextPosition = [2]int{currentPosition[0] + 1, currentPosition[1]}
		canMoveDown := guardMap[nextPosition[0]][nextPosition[1]] != obstacle
		if canMoveDown {
			guardMap = MoveDownOnMap(guardMap, currentPosition)
			newHeading = Heading{nextPosition, guardDown}
		} else {
			guardMap = TurnRight(guardMap, currentPosition, direction)
			newHeading = Heading{currentPosition, guardLeft}
		}
	case guardLeft:
		nextPosition = [2]int{currentPosition[0], currentPosition[1] - 1}
		canMoveLeft := guardMap[nextPosition[0]][nextPosition[1]] != obstacle
		if canMoveLeft {
			guardMap = MoveLeftOnMap(guardMap, currentPosition)
			newHeading = Heading{nextPosition, guardLeft}
		} else {
			guardMap = TurnRight(guardMap, currentPosition, direction)
			newHeading = Heading{currentPosition, guardUp}
		}
	}

	return guardMap, newHeading
}

func FindGuard(guardMap [][]byte) (byte, [2]int) {
	for rowNum, row := range guardMap {
		for colNum, character := range row {
			if slices.Contains(guardPositions, character) {
				return character, [2]int{rowNum, colNum}
			}
		}
	}
	PrintGuardMap(guardMap)
	panic("guard not found")
}

func LeavingMap(guardMap [][]byte, heading Heading) bool {
	switch heading.direction {
	case guardUp:
		return heading.Position[0] == 0
	case guardRight:
		return heading.Position[1] == len(guardMap[0])-1
	case guardDown:
		return heading.Position[0] == len(guardMap)-1
	case guardLeft:
		return heading.Position[1] == 0
	}
	return false
}

func FinalMap(guardMap [][]byte, heading Heading) [][]byte {
	guardMap[heading.Position[0]][heading.Position[1]] = visited
	return guardMap
}

func MoveUpOnMap(guardMap [][]byte, currentPosition [2]int) [][]byte {
	guardMap[currentPosition[0]][currentPosition[1]] = visited
	guardMap[currentPosition[0]-1][currentPosition[1]] = guardUp
	return guardMap
}

func MoveRightOnMap(guardMap [][]byte, currentPosition [2]int) [][]byte {
	guardMap[currentPosition[0]][currentPosition[1]] = visited
	guardMap[currentPosition[0]][currentPosition[1]+1] = guardRight
	return guardMap
}

func MoveDownOnMap(guardMap [][]byte, currentPosition [2]int) [][]byte {
	guardMap[currentPosition[0]][currentPosition[1]] = visited
	guardMap[currentPosition[0]+1][currentPosition[1]] = guardDown
	return guardMap
}

func MoveLeftOnMap(guardMap [][]byte, currentPosition [2]int) [][]byte {
	guardMap[currentPosition[0]][currentPosition[1]] = visited
	guardMap[currentPosition[0]][currentPosition[1]-1] = guardLeft
	return guardMap
}

func TurnRight(guardMap [][]byte, currentPosition [2]int, direction byte) [][]byte {
	row, col := currentPosition[0], currentPosition[1]
	switch direction {
	case guardUp:
		guardMap[row][col] = guardRight
	case guardRight:
		guardMap[row][col] = guardDown
	case guardDown:
		guardMap[row][col] = guardLeft
	case guardLeft:
		guardMap[row][col] = guardUp
	}
	return guardMap
}

func ClosedLoop(guardMap [][]byte, heading Heading, obstaclePosition [2]int) bool {
	// _, guardStartingPosition := FindGuard(guardMap)
	// if guardStartingPosition[0] == obstaclePosition[0] && guardStartingPosition[1] == obstaclePosition[1] {
	// 	return false
	// }
	updatedMap := AddObstacleToMap(guardMap, obstaclePosition)
	encounteredHeadings := []Heading{}
	prevHeading := heading
	for {
		// PrintGuardMap(updatedMap)
		_, newHeading := GuardStep(updatedMap, prevHeading)
		if reflect.DeepEqual(prevHeading, newHeading) {
			break
		}
		if slices.ContainsFunc(encounteredHeadings, func(encounteredHeading Heading) bool {
			return reflect.DeepEqual(encounteredHeading, newHeading)
		}) {
			return true
		}
		encounteredHeadings = append(encounteredHeadings, newHeading)
		prevHeading = newHeading
	}
	return false
}

func CountGoodObstaclePositions(guardMap [][]byte, startHeading Heading, visitedPositions [][2]int) int {
	goodObstacleCount := 0
	// for rowNum := range guardMap {
	// 	for colNum := range guardMap[rowNum] {
	for _, position := range visitedPositions {
		if ClosedLoop(guardMap, startHeading, position) {
			fmt.Printf("Position (%d, %d) is a good obstacle\n", position[0], position[1])
			goodObstacleCount++
		}
		// }
	}
	return goodObstacleCount
}

func PrintGuardMap(guardMap [][]byte) {
	for _, row := range guardMap {
		for _, character := range row {
			print(string(character))
		}
		println()
		println()
	}
}

func AddObstacleToMap(guardMap [][]byte, obstaclePosition [2]int) [][]byte {
	updatedMap := DeepCopyMap(guardMap)
	updatedMap[obstaclePosition[0]][obstaclePosition[1]] = obstacle
	return updatedMap
}

func CurrentGuardHeading(updatedMap [][]byte) Heading {
	direction, currentPosition := FindGuard(updatedMap)
	return Heading{currentPosition, direction}
}
