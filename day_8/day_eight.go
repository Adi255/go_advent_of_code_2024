package day8

import "advent_of_code/util"

type AntennaGroup map[byte][]Position

type Position [2]int

func LoadAntennaMap(filename string) (antennaMap [][]byte) {
	lines := util.ReadFileLines(filename)
	antennaMap = make([][]byte, len(lines))
	for i, line := range lines {
		antennaMap[i] = []byte(line)
	}
	return antennaMap
}

func LocateAntennas(antennaMap [][]byte) (antennas AntennaGroup) {
	antennas = make(AntennaGroup)
	for rowNum, row := range antennaMap {
		for colNum, col := range row {
			if col != '.' {
				antennas[col] = append(antennas[col], Position{rowNum, colNum})
			}
		}
	}
	return antennas
}

func CountUniqueAntinodes(antennaMap [][]byte) int {
	antennaGroups := LocateAntennas(antennaMap)
	antinodes := LocateAntinodes(antennaGroups, antennaMap)
	return countUniquePositions(antinodes)
}

func countUniquePositions(positions []Position) int {
	positionSet := make(map[Position]bool)
	for _, position := range positions {
		positionSet[position] = true
	}
	return len(positionSet)
}

func CountUniqueHarmonicAntinodes(antennaMap [][]byte) int {
	antennaGroups := LocateAntennas(antennaMap)
	antinodes := LocateHarmonicAntinodes(antennaGroups, antennaMap)
	return countUniquePositions(antinodes)
}

func LocateAntinodes(antennas AntennaGroup, antennaMap [][]byte) (antinodes []Position) {
	antinodes = make([]Position, 0)
	for _, positions := range antennas {
		for _, position := range positions {
			positionAntinodes := antinodesForPosition(position, positions, antennaMap)
			antinodes = append(antinodes, positionAntinodes...)
		}
	}
	return antinodes
}

func antinodesForPosition(position Position, positions []Position, antennaMap [][]byte) []Position {
	positionAntinodes := make([]Position, 0)
	for _, otherPosition := range positions {
		if otherPosition != position {
			antinodeRow := position[0] + (-1 * (otherPosition[0] - position[0]))
			antinodeCol := position[1] + (-1 * (otherPosition[1] - position[1]))
			inBounds := 0 <= antinodeRow && antinodeRow < len(antennaMap) && 0 <= antinodeCol && antinodeCol < len(antennaMap[0])
			if inBounds {
				positionAntinodes = append(positionAntinodes, Position{antinodeRow, antinodeCol})
			}
		}
	}
	return positionAntinodes
}

func LocateHarmonicAntinodes(antennas AntennaGroup, antennaMap [][]byte) (antinodes []Position) {
	antinodes = make([]Position, 0)
	for _, positions := range antennas {
		for _, position := range positions {
			positionAntinodes := harmonicAntinodesForPosition(position, positions, antennaMap)
			antinodes = append(antinodes, positionAntinodes...)
		}
	}
	return antinodes
}

func harmonicAntinodesForPosition(position Position, positions []Position, antennaMap [][]byte) []Position {
	positionAntinodes := make(map[Position]bool)
	for _, otherPosition := range positions {
		if otherPosition != position {
			deltaX := -1 * (otherPosition[0] - position[0])
			deltaY := -1 * (otherPosition[1] - position[1])
			i := 0
			for {
				antinodeRow := position[0] + i*deltaX
				antinodeCol := position[1] + i*deltaY
				inBounds := 0 <= antinodeRow && antinodeRow < len(antennaMap) && 0 <= antinodeCol && antinodeCol < len(antennaMap[0])
				if !inBounds {
					break
				} else {
					positionAntinodes[Position{antinodeRow, antinodeCol}] = true
					i++
				}
			}
		}
	}
	return keySet(positionAntinodes)
}

func keySet(positionMap map[Position]bool) []Position {
	keys := make([]Position, len(positionMap))
	i := 0
	for k := range positionMap {
		keys[i] = k
		i++
	}
	return keys
}
