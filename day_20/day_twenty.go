package day20

import (
	"advent_of_code/util"
	"math"
	"slices"
)

const START = 'S'
const END = 'E'
const TRACK = '.'

func NavigateMaze(mazePath string) [][2]int {
	raceMap, startPosition := loadRaceMapAndStart(mazePath)
	return raceRoute(raceMap, startPosition)
}

func raceRoute(raceMap [][]byte, startPosition [2]int) [][2]int {
	trackPositions := [][2]int{startPosition}
	position := startPosition
	cell := cellAt(raceMap, position)
	for cell != END {
		position = nextPosition(raceMap, position, trackPositions)
		trackPositions = append(trackPositions, position)
		cell = cellAt(raceMap, position)
	}
	return trackPositions
}

func cellAt(raceMap [][]byte, position [2]int) byte {
	return raceMap[position[1]][position[0]]
}

func loadRaceMapAndStart(mazePath string) ([][]byte, [2]int) {
	lines := util.ReadFileLines(mazePath)
	var raceMap [][]byte
	var startPosition [2]int
	for row, line := range lines {
		cols := []byte(line)
		for col, cell := range cols {
			if cell == START {
				startPosition = [2]int{col, row}
			}
		}
		raceMap = append(raceMap, cols)
	}

	return raceMap, startPosition
}

func nextPosition(raceMap [][]byte, position [2]int, visited [][2]int) [2]int {
	for _, dy := range []int{-1, 1} {
		pos := [2]int{position[0], position[1] + dy}
		if !slices.Contains(visited, pos) && (cellAt(raceMap, pos) == TRACK || cellAt(raceMap, pos) == END) {
			return pos
		}
	}

	for _, dx := range []int{-1, 1} {
		pos := [2]int{position[0] + dx, position[1]}
		if !slices.Contains(visited, pos) && (cellAt(raceMap, pos) == TRACK || cellAt(raceMap, pos) == END) {
			return pos
		}
	}
	panic("ran out of road")
}

func CountTimeSavingCheats(raceTrack [][2]int, thresholdPs int, cheatRange int) int {
	cheatCount := 0
	trackLength := len(raceTrack)
	for idx, pos := range raceTrack[:trackLength-thresholdPs] {
		// cheat must save at least thresholdPs (i.e reachable in 3 moves and
		// saving (otherIdx - currentIdx - 3) must be greater than thresholdPs)
		potentialDestinations := raceTrack[idx+3:]
		// normal time between points = otherIdx - currentIdx
		// time saved = normal time - distance
		cheatsForPos := timeSavingCheatsFromPosition(pos, potentialDestinations, cheatRange)
		for _, cheat := range cheatsForPos {
			if cheat.saving >= thresholdPs {
				cheatCount++
			}
		}
	}
	return cheatCount
}

type Cheat struct {
	startPos [2]int
	endPos   [2]int
	saving   int
}

func timeSavingCheatsFromPosition(position [2]int, potentialDestinations [][2]int, cheatRange int) []Cheat {
	cheats := []Cheat{}
	for otherIdx, otherPos := range potentialDestinations {
		distance := distance(position, otherPos)
		if distance >= 2 && distance <= cheatRange {
			cheats = append(cheats, Cheat{position, otherPos, otherIdx + 3 - distance})
		}
	}
	return cheats
}

func distance(a, b [2]int) int {
	dx := float64(a[0] - b[0])
	dy := float64(a[1] - b[1])
	return int(math.Abs(dx) + math.Abs(dy))
}
