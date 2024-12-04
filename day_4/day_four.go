package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
	NorthEast
	NorthWest
	SouthEast
	SouthWest
)

const targetWord = "XMAS"
const targetWordLength = len(targetWord)

func LoadGrid(file string) [][]rune {
	fileHandle, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(fileHandle)
	for scanner.Scan() {
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return grid
}

func CountXmasOccurrences(wordGrid [][]rune) int {
	count := 0
	for rowNum, row := range wordGrid {
		for colNum, character := range row {
			if character == 'X' {
				possibleSearchDirections := DetermineSearchDirections(rowNum, colNum, wordGrid)
				for _, direction := range possibleSearchDirections {
					if SearchForXmas(wordGrid, rowNum, colNum, direction) {
						count++
					}
				}
			}
		}
	}

	return count
}

func DetermineSearchDirections(rowNum, colNum int, grid [][]rune) []Direction {
	possibleSearchDirections := []Direction{}
	targetIndexOffset := targetWordLength - 1
	bottomIdx := len(grid) - 1
	rightIdx := len(grid[0]) - 1
	if rowNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, North)
	}
	if colNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, West)
	}
	if bottomIdx-rowNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, South)
	}
	if rightIdx-colNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, East)
	}
	if rowNum >= targetIndexOffset && colNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, NorthWest)
	}
	if rowNum >= targetIndexOffset && rightIdx-colNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, NorthEast)
	}
	if bottomIdx-rowNum >= targetIndexOffset && colNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, SouthWest)
	}
	if bottomIdx-rowNum >= targetIndexOffset && rightIdx-colNum >= targetIndexOffset {
		possibleSearchDirections = append(possibleSearchDirections, SouthEast)
	}
	return possibleSearchDirections
}

func SearchForXmas(wordGrid [][]rune, rowNum, colNum int, direction Direction) bool {
	switch direction {
	case North:
		return wordGrid[rowNum-1][colNum] == 'M' && wordGrid[rowNum-2][colNum] == 'A' && wordGrid[rowNum-3][colNum] == 'S'
	case East:
		return wordGrid[rowNum][colNum+1] == 'M' && wordGrid[rowNum][colNum+2] == 'A' && wordGrid[rowNum][colNum+3] == 'S'
	case South:
		return wordGrid[rowNum+1][colNum] == 'M' && wordGrid[rowNum+2][colNum] == 'A' && wordGrid[rowNum+3][colNum] == 'S'
	case West:
		return wordGrid[rowNum][colNum-1] == 'M' && wordGrid[rowNum][colNum-2] == 'A' && wordGrid[rowNum][colNum-3] == 'S'
	case NorthWest:
		return wordGrid[rowNum-1][colNum-1] == 'M' && wordGrid[rowNum-2][colNum-2] == 'A' && wordGrid[rowNum-3][colNum-3] == 'S'
	case NorthEast:
		return wordGrid[rowNum-1][colNum+1] == 'M' && wordGrid[rowNum-2][colNum+2] == 'A' && wordGrid[rowNum-3][colNum+3] == 'S'
	case SouthWest:
		return wordGrid[rowNum+1][colNum-1] == 'M' && wordGrid[rowNum+2][colNum-2] == 'A' && wordGrid[rowNum+3][colNum-3] == 'S'
	case SouthEast:
		return wordGrid[rowNum+1][colNum+1] == 'M' && wordGrid[rowNum+2][colNum+2] == 'A' && wordGrid[rowNum+3][colNum+3] == 'S'
	default:
		panic("unknown direction")
	}
}

func CountCrossMasOccurrences(wordGrid [][]rune) int {
	count := 0
	for rowNum, row := range wordGrid {
		for colNum, character := range row {
			if character == 'A' && crossMasFound(rowNum, colNum, wordGrid) {
				count++
			}
		}
	}
	return count
}

func crossMasFound(rowNum, colNum int, wordGrid [][]rune) bool {
	if impossibleCrossMasPosition(rowNum, colNum, wordGrid) {
		return false
	}

	leftToRightCrossMas := (wordGrid[rowNum-1][colNum-1] == 'M' && wordGrid[rowNum+1][colNum+1] == 'S') || (wordGrid[rowNum-1][colNum-1] == 'S' && wordGrid[rowNum+1][colNum+1] == 'M')
	rightToLeftCrossMas := (wordGrid[rowNum-1][colNum+1] == 'M' && wordGrid[rowNum+1][colNum-1] == 'S') || (wordGrid[rowNum-1][colNum+1] == 'S' && wordGrid[rowNum+1][colNum-1] == 'M')

	return leftToRightCrossMas && rightToLeftCrossMas
}

func impossibleCrossMasPosition(rowNum, colNum int, wordGrid [][]rune) bool {
	rightIdx := len(wordGrid[0]) - 1
	bottomIdx := len(wordGrid) - 1
	return rowNum == 0 || colNum == 0 || rowNum == bottomIdx || colNum == rightIdx
}

func main() {
	wordGrid := LoadGrid("day4_input.txt")
	fmt.Printf("XMAS count %d\n", CountXmasOccurrences(wordGrid))
	fmt.Printf("X-MAS count %d\n", CountCrossMasOccurrences(wordGrid))
}
