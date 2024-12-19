package day16

import (
	"advent_of_code/util"
	"math"
)

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

const START = 'S'
const EXIT = 'E'
const MOVE_SCORE = 1
const TURN_SCORE = 1000

type Heading struct {
	position  [2]int
	direction Direction
}

var rotationMap = map[Direction]Direction{
	North: East,
	East:  South,
	South: West,
	West:  North,
}

var directionChanges = map[Direction][2]int{
	North: {0, -1},
	South: {0, 1},
	West:  {-1, 0},
	East:  {1, 0},
}

type DirectionWithScore struct {
	direction Direction
	score     int
}

type state struct {
	heading Heading
	visited [][2]int
	score   int
}

func mapCell(reindeerMap [][]byte, position [2]int) byte {
	return reindeerMap[position[1]][position[0]]
}

func LoadReindeerMap(filename string) [][]byte {
	lines := util.ReadFileLines(filename)

	var reindeerMap [][]byte
	for _, line := range lines {
		reindeerMap = append(reindeerMap, []byte(line))
	}

	return reindeerMap
}

func LowestScoringPath(reindeerMap [][]byte) (bestScore int, bestSeats int) {
	bestScore, bestSeats = breadthFirstSearch(reindeerMap)
	return
}

func findStartingXY(reindeerMap [][]byte) [2]int {
	for y, row := range reindeerMap {
		for x, cell := range row {
			if cell == START {
				return [2]int{x, y}
			}
		}
	}
	return [2]int{}
}

func breadthFirstSearch(reindeerMap [][]byte) (bestScore int, bestSeats int) {
	startPosition := findStartingXY(reindeerMap)
	scoreToBeat := math.MaxInt
	heading := Heading{startPosition, East}

	startState := state{heading, [][2]int{startPosition}, 0}
	queue := []state{startState}

	visited := make(map[Heading]int)
	scoreToVisitedSpaces := make(map[int][][2]int)

	for len(queue) > 0 {
		currentState := queue[0]
		queue = queue[1:]

		if currentState.score > scoreToBeat {
			continue
		}

		position := currentState.heading.position
		if mapCell(reindeerMap, position) == EXIT {
			if currentState.score <= scoreToBeat {
				scoreToBeat = currentState.score
				scoreToVisitedSpaces[scoreToBeat] = append(scoreToVisitedSpaces[scoreToBeat], currentState.visited...)
			}
			continue
		}

		scoresWithDirections := possibleDirections(reindeerMap, position, currentState.heading.direction)
		for _, directionWithScore := range scoresWithDirections {
			newPosition := updatePosition(position, directionWithScore.direction)
			score := currentState.score + directionWithScore.score
			newHeading := Heading{newPosition, directionWithScore.direction}
			if previous, cached := visited[newHeading]; cached {
				if previous < score {
					continue
				}
			}
			visited[newHeading] = score
			nPath := make([][2]int, len(currentState.visited))
			copy(nPath, currentState.visited)

			queue = append(queue, state{newHeading, append(nPath, newPosition), score})
		}

	}

	countMap := make(map[[2]int]bool)
	for _, index := range scoreToVisitedSpaces[scoreToBeat] {
		countMap[index] = true
	}

	return scoreToBeat, len(countMap)
}

func possibleDirections(reindeerMap [][]byte, position [2]int, heading Direction) []DirectionWithScore {
	rightDirection := rotationMap[heading]
	leftDirection := rotationMap[rotationMap[rightDirection]]

	possibleScoredDirections := []DirectionWithScore{}
	if canMove(reindeerMap, position, heading) {
		possibleScoredDirections = append(possibleScoredDirections, DirectionWithScore{heading, MOVE_SCORE})
	}

	for _, otherDirection := range []Direction{rightDirection, leftDirection} {
		if canMove(reindeerMap, position, otherDirection) {
			possibleScoredDirections = append(possibleScoredDirections, DirectionWithScore{otherDirection, MOVE_SCORE + TURN_SCORE})
		}
	}

	return possibleScoredDirections
}

func canMove(reindeerMap [][]byte, position [2]int, direction Direction) bool {
	positionDelta := directionChanges[direction]
	nextPosition := [2]int{position[0] + positionDelta[0], position[1] + positionDelta[1]}
	nextSquare := mapCell(reindeerMap, nextPosition)
	return nextSquare != '#'
}

func updatePosition(position [2]int, direction Direction) [2]int {
	positionDelta := directionChanges[direction]
	return [2]int{position[0] + positionDelta[0], position[1] + positionDelta[1]}
}
